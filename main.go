package main

import "fmt"

func main() {
	//TEST function checkidentical
	puzzle := [3][3]int{}
	puzzle[0][0] = 1
	puzzle[0][1] = 2
	puzzle[0][2] = 3
	puzzle[1][0] = 4
	puzzle[1][1] = 5
	puzzle[1][2] = 6
	puzzle[2][0] = 7
	puzzle[2][1] = 8
	puzzle[2][2] = 9

	puzzle2 := [3][3]int{}
	puzzle2 = puzzle
	puzzle2[0][0] = 5
	fmt.Println(checkidentical(puzzle2, puzzle))

}

func checkidentical(temp1 [3][3]int, temp2 [3][3]int) bool {
	if len(temp1) != len(temp2) {
		return false
	}
	for i, v := range temp1 {
		if v != temp2[i] {
			return false
		}
	}
	return true

}
