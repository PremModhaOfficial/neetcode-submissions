func carFleet(target int, position []int, speed []int) int {
	stack := []float32{}
	mlist := make([][2]int, len(position))

	for i := range position {
		mlist[i] = [2]int{position[i], speed[i]}
	}

	sort.Slice(mlist, func(i, j int) bool {
		return mlist[i][0] < mlist[j][0]
	})

	n := len(position)
	for i := range position {
		last := n - 1 - i
		t := float32(target-mlist[last][0]) / float32(mlist[last][1])
		if len(stack) > 0 && stack[len(stack)-1] >= t {
		} else {
			stack = append(stack, t)
		}
	}

	return len(stack)
}
