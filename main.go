package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello AI")
	board := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 0}}
	blank := []int{2, 2}

	fmt.Println("------INIT--------")
	fmt.Println(board)
	fmt.Println(blank)
	move(board, blank, "l")
}

func move(board [][]int, blank []int, direction string) [][]int {
	dir := strings.ToUpper(direction)
	fmt.Println("===========")
	fmt.Println("Direction", direction)
	fmt.Println("Blank", blank)
	fmt.Println(board[0])
	fmt.Println(board[1])
	fmt.Println(board[2])

	if canMove(blank, dir) {

		switch dir {
		case "U":
			board, blank = moveU(board, blank)
		case "D":
			board, blank = moveD(board, blank)
		case "L":
			board, blank = moveL(board, blank)
		case "R":
			board, blank = moveR(board, blank)
		default:
			panic("Undefined direction")
		}
	}
	fmt.Println("----------AFTER--------")
	fmt.Println("Blank", blank)
	fmt.Println(board[0])
	fmt.Println(board[1])
	fmt.Println(board[2])
	return board
}

func canMove(blank []int, direction string) bool {
	switch direction {
	case "U":
		if blank[0] == 0 {
			return false
		}
		return true
	case "D":
		if blank[0] == 2 {
			return false
		}
		return true
	case "L":
		if blank[1] == 0 {
			return false
		}
		return true
	case "R":
		if blank[1] == 2 {
			return false
		}
		return true
	default:
		return false
	}
}

// b is blank
func moveU(board [][]int, b []int) ([][]int, []int) {
	board[b[0]][b[1]] = board[b[0]-1][b[1]]
	board[b[0]-1][b[1]] = 0
	b[0] = b[0] - 1
	return board, b
}
func moveD(board [][]int, b []int) ([][]int, []int) {
	board[b[0]][b[1]] = board[b[0]+1][b[1]]
	board[b[0]+1][b[1]] = 0
	b[0] = b[0] + 1
	return board, b
}
func moveL(board [][]int, b []int) ([][]int, []int) {
	board[b[0]][b[1]] = board[b[0]][b[1]-1]
	board[b[0]][b[1]-1] = 0
	b[1] = b[1] - 1
	return board, b
}
func moveR(board [][]int, b []int) ([][]int, []int) {
	board[b[0]][b[1]] = board[b[0]][b[1]+1]
	board[b[0]][b[1]+1] = 0
	b[1] = b[1] + 1
	return board, b
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
