func isValidSudoku(board [][]byte) bool {
	sectors := [3][3]map[string]bool{}
	cols := [9]map[string]bool{}
	rows := [9]map[string]bool{}

	for i := range 9 {
		for j := range 9 {
			if sectors[i/3][j/3] == nil {
				sectors[i/3][j/3] = make(map[string]bool)
			}
			if rows[i] == nil {
				rows[i] = make(map[string]bool)
			}
			if cols[j] == nil {
				cols[j] = make(map[string]bool)
			}

			element := string(board[i][j])
			if element == "." {
				continue
			}
			if sectors[i/3][j/3][element] {
				return false
			}
			sectors[i/3][j/3][element] = true

			if cols[j][element] {
				return false
			}
			cols[j][element] = true

			if rows[i][element] {
				return false
			}
			rows[i][element] = true

		}
	}

	return true
}
