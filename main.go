package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
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

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!")
}

func main() {

	fmt.Println("Hello AI")
	//http.HandleFunc("/", hello)
	//http.ListenAndServe(":8000", nil)
	//	goal := [][]int{{1, 0, 2}, {4, 5, 3}, {7, 8, 6}}

	init := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 0}}
	blank := []int{2, 2}
	print(init)
	fmt.Println("==========START===========")
	randomPuzzle(init, blank, 5)

	//	bfs(goal, init, blank)
}

func move(board [][]int, blank []int, direction string) ([][]int, []int, error) {
	dir := strings.ToUpper(direction)
	fmt.Println("Direction : ", dir)
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
		}
	} else {
		return board, blank, errors.New("Wrong move")
	}
	print(board)
	fmt.Println("Blank:", blank)
	fmt.Println("=========================")
	return board, blank, nil
}

func canMove(blank []int, direction string) bool {
	switch direction {
	case "U":
		if blank[0] <= 0 {
			return false
		}
	case "D":
		if blank[0] >= 2 {
			return false
		}
	case "L":
		if blank[1] <= 0 {
			return false
		}
	case "R":
		if blank[1] >= 2 {
			return false
		}
	default:
		return false
	}
	return true
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

func checkIdentical(b1 [][]int, b2 [][]int) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if b1[i][j] != b2[i][j] {
				return false
			}
		}
	}
	return true
}

//if number parameter is 0 the randomnumber will generate.
func randomPuzzle(board [][]int, b []int, randomnumber int) ([][]int, []int) {
	//Random seed.
	seed1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(seed1)
	if randomnumber == 0 {
		randomnumber = (r1.Intn(100) + 50)
	}

	//keepRune is an array of random direction sequence.
	var directorRune = []rune("UDLR")
	keepRune := make([]rune, randomnumber)
	for i := range keepRune {
		keepRune[i] = directorRune[r1.Intn(len(directorRune))]
	}
	//Move blank follow by sequence of keepRune.
	for _, direct := range keepRune {
		move(board, b, string(direct))
	}

	return board, b
}

func print(board [][]int) {
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

// BFS Call this
func bfs(goal, board [][]int, blank []int) bool { // return []Step
	time.Sleep(1000 * time.Millisecond)
	if checkIdentical(board, goal) {
		fmt.Println("No move require")
		return true
	} else {
		if bfsTv(goal, board, blank, "U") {
			fmt.Println("U")
			return true
		} else if bfsTv(goal, board, blank, "D") {
			fmt.Println("D")
			return true
		} else if bfsTv(goal, board, blank, "L") {
			fmt.Println("L")
			return true
		} else if bfsTv(goal, board, blank, "R") {
			fmt.Println("R")
			return true
		}
	}
	return false
}

// tv=traverse
func bfsTv(goal, board [][]int, blank []int, direction string) bool {
	board_new, blank_new, err := move(board, blank, direction)
	if err != nil {
		return false
	}
	if checkIdentical(board_new, goal) {
		fmt.Println("Success")
		return true
	} else {
		if bfs(goal, board_new, blank_new) {
			return true
		}
	}
	return false
}
