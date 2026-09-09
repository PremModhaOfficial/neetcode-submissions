func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	s = strings.Trim(s, "!?,.;:'\" ")
	left, right := 0, len(s)-1

	for left < right {
		for strings.ContainsAny(string(s[left]), "!?,.;:'\" ") {
			left++
		}
		for strings.ContainsAny(string(s[right]), "!?,.;:'\" ") {
			right--
		}

		if s[left] != s[right] {
			return false
		}

		left++
		right--
	}

	return true
}
