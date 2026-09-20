func dailyTemperatures(temperatures []int) []int {
	dailyT := make([]int, len(temperatures))

	stack := make([][2]int, len(temperatures))
	top := -1

	for curr_ind, curr_temp := range temperatures {
		for top >= 0 && stack[top][1] < curr_temp {
			ind := stack[top][0]
			dailyT[ind] = curr_ind - ind
			top--
		}

		top++
		stack[top] = [2]int{curr_ind, curr_temp}
	}

	return dailyT
}

