
func topKFrequent(nums []int, k int) []int {
	elmFreq := make(map[int]int)
	freqEle := make([]*[]int, len(nums)+1)

	ans := *new([]int)

	for _, n := range nums {
		elmFreq[n] += 1
	}

	for elem, freq := range elmFreq {
		if freqEle[freq] == nil {
			freqEle[freq] = new([]int)
		}

		*freqEle[freq] = append(*freqEle[freq], elem)
	}

	// fmt.Println(freqEle)

outer:
	for i := len(freqEle) - 1; i >= 0; i-- {
		if freqEle[i] == nil {
			continue
		}
		for _,elemenet := range *freqEle[i] {
			ans = append(ans, elemenet)
			if len(ans) == k {
				break outer
			}
		}
	}

	return ans
}
