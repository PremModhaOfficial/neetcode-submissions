func twoSum(numbers []int, target int) []int {
	ans := []int{}

	set := map[int]int{}

	for i, n := range numbers {
		diff := target - n

		if indx, exist := set[diff]; exist {
			return []int{indx+1, i+1}
		}
		set[n] = i

	}

	return ans
}

