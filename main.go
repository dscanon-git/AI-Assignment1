package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type Response struct {
	Init     []int
	Sequence []Step
}
type Step struct {
	Tile      int
	Direction string
}

func main() {
	fmt.Println("Hello AI")
	board := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 0}}
	blank := []int{2, 2}

	fmt.Println("------INIT--------")
	fmt.Println(board)
	fmt.Println(blank)

	fmt.Println("Start Random")
	fmt.Println(randomPuzzle(board, blank))
	fmt.Println("End Random")

	returnToFont(board)

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

func randomPuzzle(board [][]int, b []int) ([][]int, []int) {
	//Random Number
	seed1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(seed1)
	randomnumber := (r1.Intn(100) + 50)
	for i := 0; i < randomnumber; i++ {
		//Random Director
		seed2 := rand.NewSource(time.Now().UnixNano())
		r2 := rand.New(seed2)
		randomdirector := (r2.Intn(100) % 4) + 1
		switch randomdirector {
		case 1:
			if canMove(b, "U") {
				fmt.Println("move blank up")
				board, b = moveU(board, b)
			} else {
				fmt.Println("Can't move")
			}
		case 2:
			if canMove(b, "D") {
				fmt.Println("move blank down")
				board, b = moveD(board, b)
			} else {
				fmt.Println("Can't move")
			}
		case 3:
			if canMove(b, "L") {
				fmt.Println("move blank Left")
				board, b = moveL(board, b)
			} else {
				fmt.Println("Can't move")
			}
		case 4:
			if canMove(b, "R") {
				fmt.Println("move blank Right")
				board, b = moveR(board, b)
			} else {
				fmt.Println("Can't move")
			}
		default:
			fmt.Println("Random move failed")
		}
		printPuzzle(board)
	}

	return board, b
}

func printPuzzle(board [][]int) {
	fmt.Println(board[0])
	fmt.Println(board[1])
	fmt.Println(board[2])
}

func returnToFont(board [][]int) {
	fmt.Println("Enter returnToont")
	initpuzzle := []int{board[0][0], board[0][1], board[0][2],
		board[1][0], board[1][1], board[1][2],
		board[2][0], board[2][1], board[2][2]}

	res := &Response{
		Init: initpuzzle,
		Sequence: []Step{
			Step{
				Tile:      5,
				Direction: "U",
			},
			Step{
				Tile:      7,
				Direction: "U",
			},
		},
	}
	jsonres, _ := json.Marshal(res)
	fmt.Println(string(jsonres))
}
