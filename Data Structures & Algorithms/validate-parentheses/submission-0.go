func isValid(s string) bool {
	state := []rune{}
	pairs := map[rune]rune{')': '(', '}': '{', ']': '['}

	for _, ru := range s {
		if open, isClose := pairs[ru]; isClose {
			if len(state) > 0 && state[len(state)-1] == open {
				state = state[:len(state)-1]
			} else {
				return false
			}
		} else {
			state = append(state, ru)
		}
	}

	return len(state) == 0
}

