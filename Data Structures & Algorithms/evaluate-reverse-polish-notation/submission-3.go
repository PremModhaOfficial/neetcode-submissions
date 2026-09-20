func evalRPN(tokens []string) int {
	op := map[string]bool{"+": true, "-": true, "/": true, "*": true}
	stack := []int{}

	for _, tok := range tokens {
		if _, exist := op[tok]; exist {
			last1st := stack[len(stack)-1]
			last2nd := stack[len(stack)-2]

			stack = stack[:len(stack)-2]

			switch tok {
			case "*":
				stack = append(stack, last1st*last2nd)
			case "-":
				stack = append(stack, last2nd-last1st)
			case "/":
				stack = append(stack, last2nd/last1st)
			case "+":
				stack = append(stack, last1st+last2nd)
			}
		} else {
			n, _ := strconv.Atoi(tok)
			stack = append(stack, n)
		}
	}

	return stack[0]
}


