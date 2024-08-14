package array_shashing

func isValidSudoku(board [][]byte) bool {

	var verifyBoard, verifyColumns, verifyRow [9][9]bool

	currentIndex3x3 := 0
	indexBoards3x3 := [81]uint8{
		0, 0, 0, 1, 1, 1, 2, 2, 2,
		0, 0, 0, 1, 1, 1, 2, 2, 2,
		0, 0, 0, 1, 1, 1, 2, 2, 2,
		3, 3, 3, 4, 4, 4, 5, 5, 5,
		3, 3, 3, 4, 4, 4, 5, 5, 5,
		3, 3, 3, 4, 4, 4, 5, 5, 5,
		6, 6, 6, 7, 7, 7, 8, 8, 8,
		6, 6, 6, 7, 7, 7, 8, 8, 8,
		6, 6, 6, 7, 7, 7, 8, 8, 8}

	for indexRow, miniBoard := range board {

		for indexColumn, value := range miniBoard {
			if value == '.' {
				currentIndex3x3++
				continue
			}
			value = value - '0'

			if verifyBoard[indexBoards3x3[currentIndex3x3]][value-1] &&
				verifyRow[indexRow][value-1] &&
				verifyColumns[indexColumn][value-1] {

				return false
			} else {
				verifyBoard[indexBoards3x3[currentIndex3x3]][value-1] = true
				verifyRow[indexRow][value-1] = true
				verifyColumns[indexColumn][value-1] = true
			}

			currentIndex3x3++
		}
	}

	return true
}
