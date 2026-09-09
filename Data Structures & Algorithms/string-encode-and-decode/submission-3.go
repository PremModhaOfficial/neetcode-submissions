

type Solution struct{}

const CR = "/n/r"

func (s *Solution) Encode(strs []string) string {
	var strLenghts strings.Builder
	var strsAp strings.Builder

	for _, st := range strs {
		stl := len(st)
		strLenghts.WriteString((" " + strconv.Itoa(stl)))
		strsAp.WriteString(st)
	}

	return strLenghts.String() + CR + strsAp.String()
}

func (s *Solution) Decode(encoded string) []string {
	spits := strings.Split(encoded, CR)
	//  []int for sizes

	splitAtsEn := spits[0]
	if splitAtsEn == "" {
		return *new([]string)
	}
	payloadEn := spits[1]

	// decode the splitAts
	splitAts := *new([]int)
	for _, strSplitAt := range strings.Split(strings.Trim(splitAtsEn, " "), " ") {
		in, _ := strconv.Atoi(strSplitAt)
		splitAts = append(splitAts, in)
	}


	accumilator := 0
	bytez := []byte(payloadEn)
	ret := *new([]string)
	for _, jump := range splitAts {
		ret = append(ret, string(bytez[accumilator:accumilator+ jump]))
		accumilator += jump
	}

	return ret
}
