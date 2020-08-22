package sequence

// https://golang.org/src/internal/bytealg/bytealg.go
// https://segmentfault.com/a/1190000016554961

// primeRK is the prime base used in Rabin-Karp algorithm.
const primeRK = 16777619

func indexRabinKarp(text, pattern string) int {
	m, n := len(text), len(pattern)
	if m == 0 {
		return 0
	}
	if n > m {
		return -1
	}

	// Rabin-Karp search
	// hashPattern 是 pattern 根据上述方法计算出的 hash 值
	// pow 是 primeRK 的 len(pattern) 次幂
	hashPattern, pow := hashStr(pattern)

	// 计算 s[:n] 的 hash 值
	var h uint32
	for i := 0; i < n; i++ {
		h = h*primeRK + uint32(text[i])
	}
	if h == hashPattern && text[:n] == pattern {
		return 0
	}

	for i := n; i < m; {
		// 计算下一个子串的 hash 值
		h *= primeRK
		h += uint32(text[i])
		h -= pow * uint32(text[i-n])
		i++

		// 如果 hash 相等 且子串相等，则返回对应下标
		if h == hashPattern && text[i-n:i] == pattern {
			return i - n
		}
	}
	return -1
}

func hashStr(str string) (uint32, uint32) {
	hash := uint32(0)
	for i := 0; i < len(str); i++ {
		hash = hash*primeRK + uint32(str[i])
	}

	// 下面我用 python 的乘方元素符 ** 代表乘方
	// 我们已 len(str) 为 6 为例来看此函数
	// 6 的二进制是 110
	// 每次循环，pow 和 sq 分别为
	// i: 110  pow: 1  							sq: rk
	// i: 11   pow: 1  							sq: rk ** 2
	// i: 1    pow: 1 * (rk ** 2)  				sq: rk ** 4
	// i: 0    pow: 1* (rk ** 2) * (rk ** 4)  	sq: rk ** 8
	// pow: 1* (rk ** 2) * (rk ** 4) = 1 * (rk ** 6) 即是 pow(rk, 6)
	var pow, sq uint32 = 1, primeRK
	for i := len(str); i > 0; i >>= 1 {
		if i&1 != 0 {
			pow *= sq
		}
		sq *= sq
	}
	return hash, pow
}
