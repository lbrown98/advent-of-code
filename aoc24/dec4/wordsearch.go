package dec4

import "strings"

// todo: incomplete
func WordSearch(input string) int {
	inputArr := breakIntoArray(input)
	totalXmas := 0

	for i, row := range(inputArr){
		indices := searchForLetter(row, 'X')

		totalXmas += checkHorizontal(row, indices)
		totalXmas += checkVertical(inputArr, i, indices)
		totalXmas += checkDiagonal(inputArr, i, indices)
	}
	return totalXmas
}

func breakIntoArray(input string) []string {
	return strings.Split(input, "\n")
}

func searchForLetter(row string, letter rune) []int {
	var indices []int

	for i, r := range row {
		if r == letter {
			indices = append(indices, i)
		}
	}
	return indices
}

func checkHorizontal(row string, indices []int) int {
	count := 0
	for _, index := range indices {
		if index >= 3 { // check backwards
			substring := row[index-3 : index+1]
			if substring == "SAMX" {
				count += 1
			}
		}
		if index <= len(row)-4 { // check forwards
			substring := row[index: index+4]
			if substring == "XMAS" {
				count += 1
			}
		}
	}
	return count
}

func checkVertical(input []string, row_number int, indices []int) int {
	count := 0

	for _, index := range(indices) {
		if row_number >= 3 { // check up 
			substring := string(input[row_number][index]) + 
					string(input[row_number-1][index]) +
					string(input[row_number-2][index]) + 
					string(input[row_number-3][index])
			if substring == "XMAS" {
				count += 1
			}
		} 
		if row_number <= len(input) - 4 { // check down
			substring := string(input[row_number][index]) + 
						string(input[row_number+1][index]) +
						string(input[row_number+2][index]) + 
						string(input[row_number+3][index])
			if substring == "XMAS" {
				count += 1
			}
		} 
	}
	return count
}

func checkDiagonal(input []string, row_number int, indices []int) int {
	row := input[row_number]
	count := 0

	for _, index := range(indices){
		if row_number >= 3 { // check up 
			if index >= 3 { // check backwards
				substring := string(input[row_number][index]) + 
					string(input[row_number-1][index-1]) +
					string(input[row_number-2][index-2]) + 
					string(input[row_number-3][index-3])
				if substring == "XMAS" {
					count += 1
				}
			}
			if index <= len(row)-4 { // check forwards
				substring := string(input[row_number][index]) + 
					string(input[row_number-1][index+1]) +
					string(input[row_number-2][index+2]) + 
					string(input[row_number-3][index+3])
				if substring == "XMAS" {
					count += 1
				}
			}
		}
		if row_number <= len(input) - 4 { // check down
			if index >= 3 { // check backwards
				substring := string(input[row_number][index]) + 
					string(input[row_number+1][index-1]) +
					string(input[row_number+2][index-2]) + 
					string(input[row_number+3][index-3])
				if substring == "XMAS" {
					count += 1
				}
			}
			if index <= len(row)-4 { // check forwards
				substring := string(input[row_number][index]) + 
					string(input[row_number+1][index+1]) +
					string(input[row_number+2][index+2]) + 
					string(input[row_number+3][index+3])
				if substring == "XMAS" {
					count += 1
				}
			}
		}
	}

	return count
}
