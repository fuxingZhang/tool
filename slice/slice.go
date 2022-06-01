package slice

// ContainsInt  是否包含整数
func ContainsInt(ints []int, val int) bool {
	for _, v := range ints {
		if v == val {
			return true
		}
	}
	return false
}

// ContainsStr  是否包含字符串
func ContainsStr(strs []string, s string) bool {
	for _, str := range strs {
		if str == s {
			return true
		}
	}
	return false
}
