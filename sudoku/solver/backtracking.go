package solver

func (b *Board) Backtrack() bool {
	nextRow, nextCol, hasEmptyCell := b.findEmptyCell()
	if !hasEmptyCell {
		return true
	}

	for candidate := 1; candidate <= N; candidate++ {
		if b.isDigitValid(nextRow, nextCol, candidate) {
			b.Cells[nextRow][nextCol] = candidate

			if b.Backtrack() {
				return true
			}
			// reset the cell
			b.Cells[nextRow][nextCol] = 0
		}
	}

	return false
}

func (b *Board) findEmptyCell() (int, int, bool) {
	for row := 0; row < N; row++ {
		for col := 0; col < N; col++ {
			if b.Cells[row][col] == 0 {
				return row, col, true
			}
		}
	}

	return 0, 0, false
}

func (b *Board) isDigitValid(row, col, digit int) bool {
	startRow := row / 3 * 3
	startCol := col / 3 * 3

	for i := 0; i < N; i++ {
		// check the corresponding row and column
		if b.Cells[row][i] == digit ||
			b.Cells[i][col] == digit ||
			// check the corresponding 3x3 section
			b.Cells[startRow+i/3][startCol+i%3] == digit {
			return false
		}
	}

	return true
}
