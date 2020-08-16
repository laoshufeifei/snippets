package sequence

// https://www.cnblogs.com/lanxuezaipiao/p/3452579.html
// 坏字符规则：后移位数 = 坏字符的位置 - 搜索词中的上一次出现位置。
// 好后缀规则：后移位数 = 好后缀的位置 - 搜索词中的上一次出现位置
// 1) ABCDAB 中 AB 好后缀，位置是 5，上一次出现的位置是 1
// 2) ABCDEF 中 EF 好后缀，位置是 5，上一次出现的位置是 -1
// 3) EFABCDEF 有多个好后缀: CDEF DEF EF F，优先选取最长切也属于前缀的那一部分，即 EF

// 只考虑 ascii 码
const (
	maxChars = 128
)

func initBadCharacters(pattern string) []int {
	patternLen := len(pattern)
	bcs := make([]int, maxChars)

	for i := 0; i < maxChars; i++ {
		bcs[i] = patternLen
	}

	for i := 0; i < patternLen; i++ {
		idx := pattern[i]
		value := patternLen - 1 - i
		bcs[idx] = value
		// fmt.Printf("set bad character for %c as %d\n", pattern[i], value)
	}

	return bcs
}

func initGoodSuffixes(pattern string) []int {
	patternLen := len(pattern)
	gs := make([]int, patternLen)
	suffixes := initSuffixes(pattern)

	// 第三种情况
	for i := 0; i < patternLen; i++ {
		gs[i] = patternLen
	}

	// 第二种情况
	j := 0
	for i := patternLen - 2; i >= 0; i-- {
		if suffixes[i] == i+1 {
			for ; j < patternLen-1-i; j++ {
				gs[j] = patternLen - 1 - i
			}
		}
	}

	// 第一种情况
	for i := 0; i <= patternLen-2; i++ {
		j = patternLen - 1 - suffixes[i]
		gs[j] = patternLen - 1 - i
	}

	return gs
}

func boyerMoore(text, pattern string) int {
	textLen, patternLen := len(text), len(pattern)
	if patternLen == 0 || textLen == 0 || patternLen > textLen {
		return -1
	}

	bs := initBadCharacters(pattern)
	gs := initGoodSuffixes(pattern)

	// ti 表示在 text 中的移动的索引，pi 是在 pattern 中移动
	pi, ti := 0, 0
	for ti <= textLen-patternLen {
		for pi = patternLen - 1; pi >= 0 && pattern[pi] == text[pi+ti]; pi-- {
		}

		if pi < 0 {
			// j += gs[0]k
			// 暂只做第一次的搜索
			return ti
		}

		tmp := text[pi+ti]
		bcStep := bs[tmp] - patternLen + 1 + pi

		gsStep := gs[pi]
		if bcStep > gsStep {
			ti += bcStep
		} else {
			ti += gsStep
		}
	}

	return -1
}

func initSuffixes(pattern string) []int {
	patternLen := len(pattern)
	suffixes := make([]int, patternLen)

	j, k := 0, 0
	suffixes[patternLen-1] = patternLen
	for i := patternLen - 2; i >= 0; i-- {
		// k = patternLen - 1 - i + j
		// 一开始 i == j 即 k = patternLen - 1
		// 之后 j--, 不就是 k-- 吗？搞的那么复杂
		j, k = i, patternLen-1
		for j >= 0 && pattern[j] == pattern[k] {
			j--
			k--
			// k = patternLen - 1 - i + j
		}
		suffixes[i] = i - j
	}

	return suffixes
}

func initSuffixes2(pattern string) []int {
	m := len(pattern)
	suffixes := make([]int, m)

	f, g := 0, m-1
	suffixes[m-1] = m

	for i := m - 2; i >= 0; i-- {
		tmp := m - 1 - (f - i)
		if i > g && suffixes[tmp] < i-g {
			suffixes[i] = suffixes[tmp]
			continue
		}

		if i < g {
			g = i
		}
		f = i

		tmp = m - 1 - (f - g)
		for g >= 0 && pattern[g] == pattern[tmp] {
			g--
			tmp--
			// tmp = m - 1 - (f - g)
		}

		suffixes[i] = f - g
	}

	return suffixes
}
