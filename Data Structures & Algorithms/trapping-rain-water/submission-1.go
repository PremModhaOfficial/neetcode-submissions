func trap(L []int) int {
	traped := 0

	ln := len(L) - 1

	lMax := make([]int, ln+1)
	rMax := make([]int, ln+1)
	for i := range lMax {
		revInd := ln - i
		if i == 0 {
			lMax[i] = L[i]
			rMax[revInd] = L[revInd]
			continue
		}
		lMax[i] = max(L[i], lMax[i-1])
		rMax[revInd] = max(rMax[revInd+1], L[revInd])
	}

	for i := range ln {
		traped += max(min(lMax[i], rMax[i])-L[i], 0)
	}
	return traped
}
