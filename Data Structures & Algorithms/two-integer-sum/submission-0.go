func twoSum(nums []int, target int) []int {
	set := make(map[int]int) 

    for i, value := range nums {
		need := target - value
		if what, exist := set[need]; exist {
			return []int{what,i}
		}
		set[value] = i 
	}
	return []int{}
}
