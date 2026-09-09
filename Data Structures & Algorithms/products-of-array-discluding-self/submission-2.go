func productExceptSelf(nums []int) []int {
	L := len(nums)
	postfix := make([]int, L)
	prefix := make([]int, L)
	acc := 1
	for i, n := range nums {
		acc *= n
		prefix[i] = acc
	}
	acc = 1

	for i := range nums {
		in := L - 1 - i
		eleme := nums[in]
		acc *= eleme
		postfix[in] = acc

	}

	ans := make([]int, L)
	ans[0] = postfix[1]
	ans[L-1] = prefix[L-2]

	for i := range nums {
		if i == 0 || i == L-1 {
			continue
		}
		ans[i] = postfix[i+1] * prefix[i-1]
	}

	return ans
}

