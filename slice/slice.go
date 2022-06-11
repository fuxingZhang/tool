package slice

// ContainsInt  ContainsInt reports whether val is within ints.
func ContainsInt(ints []int, val int) bool {
	for _, v := range ints {
		if v == val {
			return true
		}
	}
	return false
}

// ContainsStr  ContainsStr reports whether s is within strs.
func ContainsStr(strs []string, s string) bool {
	for _, str := range strs {
		if str == s {
			return true
		}
	}
	return false
}
