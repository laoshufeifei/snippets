package geohash

// https://halfrost.com/go_spatial_search/
// http://www.movable-type.co.uk/scripts/geohash.html
// https://mp.weixin.qq.com/s/2B-VJ2xgwxrmsSkE6zuoPA
// https://rawgit.com/rzanato/geohashgrid/master/geohashgrid.html

/*
故意多加了一行一列
偶数:
zprxz
ynqwy
vjmtv
uhksu
g57eg
f46df
c139c
b028b
zprxz

奇数:
zbcfguvyz
x89destwx
r2367kmqr
p0145hjnp
zbcfguvyz
*/

import (
	"bytes"
	"fmt"
	"strings"
)

const (
	maxLatitude  float64 = 90.
	minLatitude  float64 = -90.
	maxLongitude float64 = 180.
	minLongitude float64 = -180.
)

type directionType int

const (
	north directionType = iota
	south
	east
	west
)

var (
	bits        = []int{16, 8, 4, 2, 1} // 组成 base32 字符的每一位的权重 (二进制)
	base32Bytes = []byte("0123456789bcdefghjkmnpqrstuvwxyz")

	// 偶数用第一个，奇数用第二个
	// 以 north 的偶数举例，从原本的 0 往下偏移一个，即从 p 开始重新编码
	// 以 2 举例，2在新的编码里是第四个，所以 2 的 north 就是原编码中的第四个(3)
	baseNeighbours = [4][2][]byte{
		{
			[]byte("p0r21436x8zb9dcf5h7kjnmqesgutwvy"),
			[]byte("bc01fg45238967deuvhjyznpkmstqrwx"),
		}, // north
		{[]byte("14365h7k9dcfesgujnmqp0r2twvyx8zb"), []byte("238967debc01fg45kmstqrwxuvhjyznp")}, // south
		{[]byte("bc01fg45238967deuvhjyznpkmstqrwx"), []byte("p0r21436x8zb9dcf5h7kjnmqesgutwvy")}, // east
		{[]byte("238967debc01fg45kmstqrwxuvhjyznp"), []byte("14365h7k9dcfesgujnmqp0r2twvyx8zb")}, // west
	}
	borders = [4][2]string{
		{"prxz", "bcfguvyz"}, // north
		{"028b", "0145hjnp"}, // south
		{"bcfguvyz", "prxz"}, // east
		{"0145hjnp", "028b"}, // west
	}
)

type cell struct {
	minLat, maxLat float64 // 纬度
	minLng, maxLng float64 // 经度
	code           string
}

func encode(longitude, latitude float64, precision int) *cell {
	var code bytes.Buffer
	minLat, maxLat := minLatitude, maxLatitude
	minLng, maxLng := minLongitude, maxLongitude
	var mid float64

	bit, ch, length, isEven := 0, 0, 0, true
	for length < precision {
		if isEven {
			// 偶数，取经度
			mid = (minLng + maxLng) / 2.
			if mid < longitude {
				ch |= bits[bit]
				minLng = mid
			} else {
				maxLng = mid
			}
		} else {
			// 奇数，取纬度
			mid = (minLat + maxLat) / 2.
			if mid < latitude {
				ch |= bits[bit]
				minLat = mid
			} else {
				maxLat = mid
			}
		}

		isEven = !isEven
		bit++
		if bit == 5 {
			code.WriteByte(base32Bytes[ch])
			length++
			bit, ch = 0, 0
		}
	}

	return &cell{
		minLat: minLat,
		maxLat: maxLat,
		minLng: minLng,
		maxLng: maxLng,
		code:   code.String(),
	}
}

func decode(code string) *cell {
	var mid float64
	minLat, maxLat := minLatitude, maxLatitude
	minLng, maxLng := minLongitude, maxLongitude

	isEven := true
	codes := []byte(code)
	for i := 0; i < len(codes); i++ {
		char := codes[i]
		idx := bytes.IndexByte(base32Bytes, char)
		if idx == -1 {
			break
		}

		for j := 4; j >= 0; j-- {
			bit := (idx >> j) & 0x1
			if isEven {
				// 偶数，经度
				mid = (minLng + maxLng) / 2.
				if bit == 1 {
					minLng = mid
				} else {
					maxLng = mid
				}
			} else {
				// 奇数，纬度
				mid = (minLat + maxLat) / 2.
				if bit == 1 {
					minLat = mid
				} else {
					maxLat = mid
				}
			}
			isEven = !isEven
		}
	}

	return &cell{
		minLat: minLat,
		maxLat: maxLat,
		minLng: minLng,
		maxLng: maxLng,
		code:   code,
	}
}

func adjacent(code string, direction directionType) string {
	length := len(code)
	lastChar, parent := code[length-1], code[:length-1]
	flag := length & 0x1 // 偶数 0、奇数 1

	border := borders[direction][flag]
	if strings.IndexByte(border, lastChar) != -1 && length > 1 {
		parent = adjacent(parent, direction)
	}

	neighbour := baseNeighbours[direction][flag]
	idx := bytes.IndexByte(neighbour, lastChar)
	return fmt.Sprintf("%s%c", parent, base32Bytes[idx])
}

//	northwest(0)	north(1)	northeast(2)
//	west(3)						east(4)
//	southwest(5)	south(6)	southeast(7)
func getNeighbours(code string) []string {
	results := make([]string, 8)

	results[1] = adjacent(code, north) // north
	results[6] = adjacent(code, south) // south
	results[4] = adjacent(code, east)  // east
	results[3] = adjacent(code, west)  // west

	results[0] = adjacent(results[1], west) // northwest
	results[2] = adjacent(results[1], east) // northeast
	results[5] = adjacent(results[6], west) // southwest
	results[7] = adjacent(results[6], east) // southeast

	return results
}
