
func isValidSudoku(board [][]byte) bool {
	sectors := make([]map[string]bool, 9)
	rows := make([]map[string]bool, 9)
	cols := make([]map[string]bool, 9)
	for i := range 9 {
		for j := range 9 {
			// sector
			sectorI := i / 3
			sectorJ := j / 3

			sectorIndex := sectorJ + sectorI*3

			if sectors[sectorIndex] == nil {
				sectors[sectorIndex] = make(map[string]bool)
			}
			if rows[i] == nil {
				rows[i] = make(map[string]bool)
			}
			if cols[j] == nil {
				cols[j] = make(map[string]bool)
			}
			ele := string(board[i][j])
			if ele != "." {
				if sectors[sectorIndex][ele] {
					return false
				}
				sectors[sectorIndex][ele] = true

				if cols[j][ele] {
					return false
				}
				cols[j][ele] = true

				if rows[i][ele] {
					return false
				}
				rows[i][ele] = true
			}
		}
	}

	return true
}
