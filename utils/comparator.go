package utils

// Comparator is a function compare two values.
// if a == b, return 0;
// if a > b, return 1;
// if a < b, return -1;
type Comparator func(a, b interface{}) int

// IntComparator compare two int
func IntComparator(a, b interface{}) int {
	aValue := a.(int)
	bValue := b.(int)
	switch {
	case aValue == bValue:
		return 0
	case aValue > bValue:
		return 1
	default:
		return -1
	}
}

// StringComparator compare two string
func StringComparator(a, b interface{}) int {
	s1 := a.(string)
	s2 := b.(string)

	len1 := len(s1)
	len2 := len(s2)
	minLen := len2
	if len1 < len2 {
		minLen = len1
	}

	diff := 0
	for i := 0; i < minLen; i++ {
		diff = int(s1[i]) - int(s2[i])
		if diff != 0 {
			break
		}
	}

	if diff == 0 {
		diff = len1 - len2
	}

	switch {
	case diff > 0:
		return 1
	case diff < 0:
		return -1
	default:
		return 0
	}
}
