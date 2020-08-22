package sequence

// https://en.wikipedia.org/wiki/Boyer%E2%80%93Moore%E2%80%93Horspool_algorithm
// https://leetcode.com/problems/implement-strstr/submissions/

func same(str1, str2 string, len int) bool {
	i := len - 1
	for str1[i] == str2[i] {
		if i == 0 {
			return true
		}
		i--
	}
	return false
}

func horspool(text, pattern string) int {
	textLen, patternLen := len(text), len(pattern)
	if patternLen == 0 {
		return 0
	}
	if patternLen > textLen {
		return -1
	}

	bs := initBadCharacters(pattern)

	skip := 0
	for textLen-skip >= patternLen {
		if same(text[skip:], pattern, len(pattern)) {
			return skip
		}

		char := text[skip+patternLen-1]
		skip += bs[char]
	}

	return -1
}
