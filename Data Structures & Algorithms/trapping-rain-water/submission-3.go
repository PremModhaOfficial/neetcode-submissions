func trap(L []int) int {
	l, r := 0, len(L)-1
	lm, rm := L[l], L[r]
	trpd := 0

	for l < r {
		if lm < rm {
			l++
			lm = max(lm , L[l])
			trpd += lm - L[l]

		} else {
			r--
			rm = max(rm , L[r])
			trpd += rm - L[r]
		}
	}

	return trpd
}
