package main

import (
	"fmt"

	"github.com/transdigital/recursion"
	"github.com/transdigital/sudoku"

	"github.com/transdigital/search-array"
)

func main() {
	fmt.Println("== Without Recursion Done {-1.5, 0, 4.4, 5, 6, 7} ==")
	ints := []float64{-1.5, 0, 4.4, 5, 6, 7}
	fmt.Println(search.SearchInArray(4.5, ints))
	fmt.Println(search.SearchInArray(5.5, ints))
	fmt.Println("== Search Array Done ==")
	fmt.Println("== Without Recursion Start (1071, 1029) ==")
	fmt.Println(recursion.WithoutRecursion(1071, 1029))
	fmt.Println("== Without Recursion Done ==")
	fmt.Println("== Run Sudoku ==")
	sudoku.RunSudoku()
	fmt.Println("== Sudoku End ==")
}
