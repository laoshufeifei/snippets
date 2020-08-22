package sequence

// https://www.cnblogs.com/lanxuezaipiao/p/3452579.html
// https://ethsonliu.com/2019/11/boyer-moore.html
// https://dwnusbaum.github.io/boyer-moore-demo/
// 坏字符规则：后移位数 = 坏字符的位置 - 搜索词中的上一次出现位置。
// 好后缀规则：后移位数 = 好后缀的位置 - 搜索词中的上一次出现位置
// 如果"好后缀"有多个，则除了最长的那个"好后缀"，其他"好后缀"的上一次出现位置必须在头部
// 举例:
// ABCDAB 中 AB 好后缀，位置是 5，上一次出现的位置是 1
// ABCDEF 中 EF 好后缀，位置是 5，上一次出现的位置是 -1
// EFABCDEF 有多个好后缀: CDEF DEF EF F，优先选取最长且也属于前缀的那一部分，即 EF
// BABCDAB 有多个好后缀: DAB AB B，采用 B

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

	for i := 0; i < patternLen-1; i++ {
		char := pattern[i]
		value := patternLen - 1 - i
		bcs[char] = value
		// fmt.Printf("set bad character for %c as %d\n", pattern[i], value)
	}

	return bcs
}

func initGoodSuffixes(pattern string) []int {
	patternLen := len(pattern)
	gs := make([]int, patternLen)
	suffixes := initSuffixes(pattern)

	// 第三种情况
	// 如果完全不存在和 "好后缀" 匹配的子串，则右移整个搜索词。
	for i := 0; i < patternLen; i++ {
		gs[i] = patternLen
	}

	// 第二种情况
	// 如果不存在和 "最长好后缀" 完全匹配的子串，则选取长度最长且也属于前缀的那个 "真好后缀"
	j := 0
	for i := patternLen - 2; i >= 0; i-- {
		if suffixes[i] == i+1 {
			for ; j < patternLen-1-i; j++ {
				gs[j] = patternLen - 1 - i
			}
		}
	}

	// 第一种情况
	// 搜索词中有子串和 "最长好后缀" 完全匹配，则将最靠右的那个完全匹配的子串移动到 "最长好后缀" 的位置继续进行匹配
	for i := 0; i <= patternLen-2; i++ {
		j = patternLen - 1 - suffixes[i]
		// fmt.Printf("%d change gs[%d] from %d to %d\n", i, j, gs[j], patternLen-1-i)
		gs[j] = patternLen - 1 - i
	}

	return gs
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

func boyerMoore(text, pattern string) int {
	textLen, patternLen := len(text), len(pattern)
	if patternLen == 0 {
		return 0
	}
	if patternLen > textLen {
		return -1
	}

	bs := initBadCharacters(pattern)
	gs := initGoodSuffixes(pattern)

	// pi 是在 pattern 中每次比较最右的位置
	// ti 表示在 text 中每次移动后最左的位置，ti+pi 才是参与比较的字符
	ti, pi := 0, 0
	for ti <= textLen-patternLen {
		pi = patternLen - 1
		for pi >= 0 && text[pi+ti] == pattern[pi] {
			pi--
		}

		if pi < 0 {
			// j += gs[0]k
			// 暂只做第一次的搜索
			return ti
		}

		tmp := text[pi+ti]
		badStep := bs[tmp] - patternLen + 1 + pi
		goodStep := gs[pi]
		// fmt.Printf("bad step %d(%c), good step %d\n", badStep, tmp, goodStep)

		if badStep > goodStep {
			ti += badStep
		} else {
			ti += goodStep
		}
		// fmt.Printf("now ti is %d(%c)\n", ti, text[ti])
	}

	return -1
}
