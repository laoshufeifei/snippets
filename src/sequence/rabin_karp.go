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

	// 计算 text[:n] 的 hash 值
	var hashText uint32
	for i := 0; i < n; i++ {
		hashText = hashText*primeRK + uint32(text[i])
	}
	if hashText == hashPattern && text[:n] == pattern {
		return 0
	}

	// 比如说 32123 中找 123
	// 第一次 321 != 123，重点看怎么让 212 参与下一次比较的的
	// 321 * 10 = 3210
	// 3210 + 2 = 3212
	// 3212 - 3 * (10^3) = 3212 - 3000 = 212
	// 这里也可以先用 321 - 3 * (10 ^ 2)(注意就不是使用代码中的 pow 了，而是 primeRK ^ (n-1))
	for i := n; i < m; {
		// 计算下一个子串的 hash 值
		hashText *= primeRK
		hashText += uint32(text[i])
		hashText -= pow * uint32(text[i-n])
		i++

		// 使用 hash 快速排除，最后再精确匹配
		if hashText == hashPattern && text[i-n:i] == pattern {
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

	// 快速幂算法
	// 以 len(str) 为 6 为例来看此函数
	// 6 的二进制是 110
	// 6 = 1*2^2 + 1*2^1 + 0*2^0
	// rk ^ 6 = rk ^ (2^2 + 2^1 + 0*2^1) = rk ^ (2^2) * rk ^ (2^1) * rk ^ 0
	// rk ^ (2^2) = (rk^2) ^ 2 = (rk^2) * (rk^2) 即可以有之前的一个数值来推算出来
	// 每次循环，pow 和 tmp 分别为
	// i: 110  pow: 1  							tmp: rk
	// i: 11   pow: 1  							tmp: rk ^ 2
	// i: 1    pow: 1 * (rk ^ 2)  				tmp: rk ^ 4
	// i: 0    pow: 1* (rk ^ 2) * (rk ^ 4)  	tmp: rk ^ 8
	// pow: 1* (rk ^ 2) * (rk ^ 4) = 1 * (rk ^ 6) 即是 pow(rk, 6)
	var pow, tmp uint32 = 1, primeRK
	for i := len(str); i > 0; i >>= 1 {
		if i&1 != 0 {
			pow *= tmp
		}
		tmp *= tmp
	}
	return hash, pow
}
