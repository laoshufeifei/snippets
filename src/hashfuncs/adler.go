package hashfuncs

// https://en.wikipedia.org/wiki/Adler-32
// https://golang.org/src/hash/adler32/adler32.go
// https://mp.weixin.qq.com/s/tg3pDWCOIX36CPk8hbWARA

const (
	// mod is the largest prime that is less than 65536.
	mod = 65521

	// nmax is the largest n such that
	// 255 * n * (n+1) / 2 + (n+1) * (mod-1) <= 2^32-1.
	// It is mentioned in RFC 1950 (search for "5552").
	nmax = 5552
)

func adler32V0(data []byte) uint32 {
	a, b := uint32(1), uint32(0)
	for _, bt := range data {
		a = (a + uint32(bt)) % mod
		b = (b + a) % mod
	}
	return (b << 16) | a
}

func adler32(data []byte) uint32 {
	a, b := uint32(1), uint32(0)
	for len(data) > 0 {
		var next []byte

		// 每次处理数据量为 nmax
		// 处理完之后再 % mod
		// 防止溢出，且尽可能少的执行 mod 运算
		if len(data) > nmax {
			data, next = data[:nmax], data[nmax:]
		}

		for len(data) >= 4 {
			a += uint32(data[0])
			b += a
			a += uint32(data[1])
			b += a
			a += uint32(data[2])
			b += a
			a += uint32(data[3])
			b += a
			data = data[4:]
		}
		for _, x := range data {
			a += uint32(x)
			b += a
		}
		a %= mod
		b %= mod
		data = next
	}

	return b<<16 | a
}
