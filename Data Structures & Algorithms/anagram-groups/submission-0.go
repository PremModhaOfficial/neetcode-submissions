func groupAnagrams(strs []string) [][]string {
	strToIndex := make(map[string][]string)
	arr := new([][]string)

	for _, str := range strs {
		entr := []byte(str)
		sort.Slice(entr, func(i int, j int) bool {
			
						return entr[i] < entr[j]

		})
		strToIndex[string(entr)] = append(strToIndex[string(entr)], str)
	}

		// fmt.Println(strToIndex)


	for _, val := range strToIndex {
		*arr = append(*arr, val)
	}

	return *arr
}

