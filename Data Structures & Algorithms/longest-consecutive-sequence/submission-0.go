
func longestConsecutive(nums []int) int {
	set := map[int]bool{}
	maxSeq := 0

	for _, e := range nums {
		set[e] = true
	}

	for _, e := range nums {
		// does e-1 exist?
		if set[e-1] {
			continue
		}

		seq := 1

		for set[e+seq] {
			seq++
		}

		maxSeq = max(maxSeq, seq)

	}

	return maxSeq
}

