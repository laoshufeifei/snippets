package sequence

// 目前只处理简单字符
func brustForceSearch1(text, pattern string) int {
	textLen := len(text)
	patternLen := len(pattern)
	if patternLen == 0 {
		return 0
	}
	if patternLen > textLen {
		return -1
	}

	ti, pi := 0, 0
	// for ti < textLen && pi < patternLen {
	// ti - pi 是每一轮比较中 text 中首个比较字符的位置
	maxBegin := textLen - patternLen
	for pi < patternLen && ti-pi <= maxBegin {
		if text[ti] == pattern[pi] {
			ti++
			pi++
		} else {
			// pattern 开始的位置的下一个
			// ti = ti - pi + 1
			ti -= pi - 1
			pi = 0
		}
	}

	if pi == patternLen {
		return ti - pi
	}

	return -1
}

func brustForceSearch2(text, pattern string) int {
	textLen := len(text)
	patternLen := len(pattern)
	if textLen == 0 || patternLen == 0 || patternLen > textLen {
		return -1
	}

	// ti: [0, tMax]
	// pi: [0, patternLen)
	tMax := textLen - patternLen
	for ti := 0; ti <= tMax; ti++ {
		pi := 0
		for ; pi < patternLen; pi++ {
			ch1 := text[ti+pi]
			ch2 := pattern[pi]
			if ch1 != ch2 {
				break
			}
		}

		if pi == patternLen {
			return ti
		}
	}

	return -1
}
