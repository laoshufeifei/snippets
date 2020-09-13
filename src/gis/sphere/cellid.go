package sphere

// 参考
// https://github.com/golang/geo/blob/master/s2/cellid.go
// https://halfrost.com/go_s2_cellid/
// http://s2.sidewalklabs.com/
// https://laoshufeifei.github.io/hilbert_curve/

import "fmt"

// CellID implements the Sort interface for slices of CellIDs.
type CellID uint64

// Constants related to the bit mangling in the Cell ID.
const (
	lookupBits = 4
	swapMask   = 0x01
	invertMask = 0x02
)

const (
	faceBits = 3
	numFaces = 6

	// This is the number of levels needed to specify a leaf cell.
	maxLevel = 30

	// The extra position bit (61 rather than 60) lets us encode each cell as its
	// Hilbert curve position at the cell center (which is halfway along the
	// portion of the Hilbert curve that fills that cell).
	posBits = 2*maxLevel + 1 // 61

	// The maximum index of a valid leaf cell plus one. The range of valid leaf
	// cell indices is [0..maxSize-1].
	maxSize = 1 << maxLevel // 2 ^ 30

	wrapOffset = uint64(numFaces) << posBits
)

var (
	// 随着 ij 的变化 pos 怎么变化的
	ijToPos = [4][4]int{
		{0, 1, 3, 2}, // canonical order
		{0, 3, 1, 2}, // axes swapped
		{2, 3, 1, 0}, // bits inverted
		{2, 1, 3, 0}, // swapped & inverted
	}
	// 随着 pos 的变化 ij 是怎么变化的
	posToIJ = [4][4]int{
		{0, 1, 3, 2}, // canonical order:    (0,0), (0,1), (1,1), (1,0)
		{0, 2, 3, 1}, // axes swapped:       (0,0), (1,0), (1,1), (0,1)
		{3, 2, 0, 1}, // bits inverted:      (1,1), (1,0), (0,0), (0,1)
		{3, 1, 0, 2}, // swapped & inverted: (1,1), (0,1), (0,0), (1,0)
	}

	posToOrientation = [4]int{swapMask, 0, 0, invertMask | swapMask} // 1, 0, 0, 3

	// lookupIJ 是根据 pos 快速找到 ij 的数组
	lookupIJ [1 << (2*lookupBits + 2)]int // 2 ^ 10 = 1024[]int

	// lookupPos 是为了根据 ij(O) 快速找到 Pos(o) 的数组
	lookupPos [1 << (2*lookupBits + 2)]int // 2 ^ 10 = 1024[]int
)

func init() {
	// 初始化四个方向
	initLookupCell(0, 0, 0, 0, 0, 0)
	initLookupCell(0, 0, 0, swapMask, 0, swapMask)
	initLookupCell(0, 0, 0, invertMask, 0, invertMask)
	initLookupCell(0, 0, 0, swapMask|invertMask, 0, swapMask|invertMask)
}

// initLookupCell initializes the lookupIJ table at init time.
// origOrientation 表示最原始的朝向, orientation 为当前的朝向？
func initLookupCell(level, i, j, origOrientation, pos, orientation int) {
	if level == lookupBits {
		fmt.Printf("i, j is (%d, %d), pos is %d\n", i, j, pos)

		ij := (i << lookupBits) + j
		ijOO := (ij << 2) + origOrientation
		posO := (pos << 2) + orientation
		lookupPos[ijOO] = posO
		fmt.Printf("\ti j O is %04b %04b %02b\n", i, j, origOrientation)
		fmt.Printf("\tpos o is %04b %04b %02b\n", pos>>4, pos&0x0f, orientation)

		posOO := (pos << 2) + origOrientation
		ijO := (ij << 2) + orientation
		lookupIJ[posOO] = ijO
		return
	}

	level++
	i <<= 1
	j <<= 1
	pos <<= 2

	// orientation 表示当前的朝向
	// r 表示当前朝向下的绘制顺序
	// r[0]>>1 取出 i 位， r[0]&1 取出 j 位
	// origOrientation 是 初始化的朝向
	// orientation^posToOrientation[0] 表示下一层的朝向
	r := posToIJ[orientation]
	initLookupCell(level, i+(r[0]>>1), j+(r[0]&1), origOrientation, pos, orientation^posToOrientation[0])
	initLookupCell(level, i+(r[1]>>1), j+(r[1]&1), origOrientation, pos+1, orientation^posToOrientation[1])
	initLookupCell(level, i+(r[2]>>1), j+(r[2]&1), origOrientation, pos+2, orientation^posToOrientation[2])
	initLookupCell(level, i+(r[3]>>1), j+(r[3]&1), origOrientation, pos+3, orientation^posToOrientation[3])
}

// cellIDFromFaceIJ returns a leaf cell given its cube face (range 0..5) and IJ coordinates.
func cellIDFromFaceIJ(f, i, j int) CellID {
	// Note that this value gets shifted one bit to the left at the end
	// of the function.
	n := uint64(f) << (posBits - 1)

	// Alternating faces have opposite Hilbert curve orientations; this
	// is necessary in order for all faces to have a right-handed
	// coordinate system.
	bits := f & swapMask

	// Each iteration maps 4 bits of "i" and "j" into 8 bits of the Hilbert
	// curve position.  The lookup table transforms a 10-bit key of the form
	// "iiiijjjjoo" to a 10-bit value of the form "ppppppppoo", where the
	// letters [ijpo] denote bits of "i", "j", Hilbert curve position, and
	// Hilbert curve orientation respectively.
	for k := 7; k >= 0; k-- {
		mask := (1 << lookupBits) - 1 // 0xf 15

		iv := int((i >> uint(k*lookupBits)) & mask)
		jv := (j >> uint(k*lookupBits)) & mask
		fmt.Printf("fetch i is %04b, j is %04b\n", iv, jv)

		bits += int((i>>uint(k*lookupBits))&mask) << (lookupBits + 2)
		bits += int((j>>uint(k*lookupBits))&mask) << 2
		fmt.Printf("\ti j O is %04b %04b %02b\n", bits>>6, (bits>>2)&0xf, bits&0x3)

		bits = lookupPos[bits]
		fmt.Printf("\tpos o is %04b %04b %02b\n", bits>>6, (bits>>2)&0xf, bits&0x3)

		n |= uint64(bits>>2) << (uint(k) * 2 * lookupBits)
		bits &= (swapMask | invertMask) // 0x3
	}
	return CellID(n<<1 + 1)
}

func cellID2BitString(cid CellID) string {
	bits := make([]byte, 16)
	for i := 0; i < 16; i++ {
		bits[i] = byte(0xf & (cid >> (4 * i)))
	}

	str := ""
	for i := len(bits) - 1; i >= 0; i-- {
		str += fmt.Sprintf("%04b ", bits[i])
	}
	return str[:len(str)-1]
}
