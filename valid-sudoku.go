package main

func isValidSudoku(board [][]byte) bool {

	for _, miniBoard := range board {
		verifyBoard := [9]bool{false, false, false, false, false, false, false, false, false}

		for _, value := range miniBoard {
			if value >= 1 && value <= 9 {
				if !verifyBoard[value-1] {
					verifyBoard[value-1] = true
				} else {
					return false
				}
			} else if value == '.' {
				continue
			} else {
				return false
			}
		}
	}

	return true
}
