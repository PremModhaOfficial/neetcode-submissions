func maxArea(heights []int) int {
	maxAr := 0

	l, r := 0, len(heights)-1
	for l < r {
		maxAr = max(maxAr, (r-l)*(min(heights[l],heights[r])))
		if heights[l] > heights[r] {
			r--
		} else {
			l++
		}
	}

	return maxAr
}
