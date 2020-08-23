package sequence

// https://en.wikipedia.org/wiki/Boyer%E2%80%93Moore%E2%80%93Horspool_algorithm
// https://leetcode.com/problems/implement-strstr/submissions/
// 与 BM 的最主要区别就是失配后，只关注待匹配区域的最后一个字符作为 bad character

func horspool(text, pattern string) int {
	textLen, patternLen := len(text), len(pattern)
	if patternLen == 0 {
		return 0
	}
	if patternLen > textLen {
		return -1
	}

	bs := initBadCharacters(pattern)

	// ti 是待匹配区域的最左索引
	// pi 是 pattern 中的索引
	ti, pi := 0, 0
	for ti <= textLen-patternLen {
		pi = patternLen - 1
		for pi >= 0 && text[ti+pi] == pattern[pi] {
			pi--
		}

		if pi < 0 {
			return ti
		}

		// 只关注待匹配区的最后一个字符
		char := text[ti+patternLen-1]
		ti += bs[char]
	}

	return -1
}
