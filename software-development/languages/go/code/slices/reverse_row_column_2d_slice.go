package main

import "fmt"

func main() {
	arr := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}

	fmt.Println(arr)
	// arr = reverseRow(arr, 1)
	arr = reverseColumn(arr, 2)
	fmt.Println(arr)

}

// Note: this function assumes that all rows have the same length
func reverseColumn(s [][]int, colIndex int) [][]int {
	if colIndex < 0 || len(s) == 0 || colIndex >= len(s[0]) {
		return s
	}

	for i := 0; i < len(s)/2; i++ {
		opI := len(s) - 1 - i
		s[i][colIndex], s[opI][colIndex] = s[opI][colIndex], s[i][colIndex]
	}

	return s
}

func reverseRow(s [][]int, rowIndex int) [][]int {
	if rowIndex < 0 || rowIndex >= len(s) {
		return s
	}

	for i := 0; i < len(s[rowIndex])/2; i++ {
		opI := len(s[rowIndex]) - 1 - i
		s[rowIndex][i], s[rowIndex][opI] = s[rowIndex][opI], s[rowIndex][i]
	}

	return s
}
