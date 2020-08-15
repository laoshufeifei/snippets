package sequence

func kmpSearchV1(text, pattern string) int {
	textLen := len(text)
	patternLen := len(pattern)
	if textLen == 0 || patternLen == 0 || patternLen > textLen {
		return -1
	}

	// 求一下 next 数组
	next := initNextV1(pattern)
	return kmpImple(text, pattern, next)
}

func kmpSearchV2(text, pattern string) int {
	textLen := len(text)
	patternLen := len(pattern)
	if textLen == 0 || patternLen == 0 || patternLen > textLen {
		return -1
	}

	// 求一下 next 数组
	next := initNextV2(pattern)
	return kmpImple(text, pattern, next)
}

func kmpImple(text, pattern string, next []int) int {
	textLen := len(text)
	patternLen := len(pattern)

	ti, pi := 0, 0
	// for pi < patternLen && ti < textLen {
	// ti - pi 是每一轮比较中 text 中首个比较字符的位置
	maxBegin := textLen - patternLen
	for pi < patternLen && ti-pi <= maxBegin {
		// if pi == -1 || text[ti] == pattern[pi] {
		// 	ti++
		// 	pi++
		// } else {
		// 	pi = next[pi]
		// }

		if text[ti] == pattern[pi] {
			ti++
			pi++
		} else {
			pi = next[pi]
			// 如果 pi == -1 代表需要从 ti 的下一个位置重新比较
			if pi == -1 {
				pi = 0 // pi++
				ti++
			}
		}
	}

	if pi == patternLen {
		return ti - pi
	}

	return -1
}

func initNextV1(pattern string) []int {
	patternLen := len(pattern)
	next := make([]int, patternLen)

	i, n := 0, -1
	next[0] = -1

	for i < patternLen-1 {
		if n == -1 || pattern[i] == pattern[n] {
			i++
			n++
			next[i] = n
		} else {
			// n 退到 next[n] 位置，挣扎下，看看还有机会不
			n = next[n]
		}
	}

	return next
}

func initNextV2(pattern string) []int {
	patternLen := len(pattern)
	next := make([]int, patternLen)

	i, n := 0, -1
	next[0] = -1

	for i < patternLen-1 {
		if n == -1 || pattern[i] == pattern[n] {
			i++
			n++
			if pattern[i] == pattern[n] {
				next[i] = next[n]
			} else {
				next[i] = n
			}
		} else {
			// n 退到 next[n] 位置，挣扎下，看看还有机会不
			n = next[n]
		}
	}

	return next
}
