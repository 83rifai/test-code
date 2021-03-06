package sudoku

import (
	"fmt"
	"os"

	"github.com/transdigital/sudoku/solver"
)

// RunSudoku function
func RunSudoku() {
	board, err := solver.GetBoardFrom(os.Stdin)

	if err != nil {
		fmt.Printf("Unable to process the input: %s\n", err)
		os.Exit(1)
	}

	if board.Backtrack() {
		fmt.Println("The Sudoku was solved successfully:")
		fmt.Println(board.String())
	} else {
		fmt.Printf("The Sudoku can't be solved.")
	}
}
