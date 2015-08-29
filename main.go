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
type State struct {
	board [][]int
	blank []int
	sol   string
}

func main() {

	fmt.Println("Hello AI")
	//http.HandleFunc("/", hello)
	//http.ListenAndServe(":8000", nil)
	goal := [][]int{{1, 0, 2}, {4, 5, 3}, {7, 8, 6}}

	init := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 0}}
	blank := []int{2, 2}
	fmt.Println("==========START===========")
	init, blank = randomPuzzle(init, blank, 50)
	start := copyBoard(init)
	fmt.Println("===========BFS===========")
	sol := bfs(goal, init, blank)
	fmt.Println("Solution : ", sol.sol)
	step, tile := changeBlanktoTile(init, blank, sol.sol)
	fmt.Println(returnToFront(start, step, tile))
	fmt.Println("******Goal******", goal)
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!")
}

// BFS Call this
func bfs(goal, init [][]int, blank []int) State { // return []Step
	var stateQ []State = make([]State, 1)
	stateQ[0].board = init
	stateQ[0].sol = ""
	stateQ[0].blank = blank
	fmt.Println(stateQ)
	i := 0
	//for {
	for q := 0; ; q++ {
		// Dequeue, check with goal
		if checkIdentical(stateQ[q].board, goal) {
			fmt.Println("---------------SUCCESS------------------")
			fmt.Println(stateQ[q])
			print(stateQ[q].board)
			return stateQ[q]
		}
		// Move UDLR , Enqueue
		if u, err := bfsMove(&stateQ[q], "U"); err == nil {
			i++
			//stateQ[i] = u
			stateQ = append(stateQ, u)
		}
		if d, err := bfsMove(&stateQ[q], "D"); err == nil {
			i++
			//stateQ[i] = d
			stateQ = append(stateQ, d)
		}
		if l, err := bfsMove(&stateQ[q], "L"); err == nil {
			i++
			//stateQ[i] = l
			stateQ = append(stateQ, l)
		}
		if r, err := bfsMove(&stateQ[q], "R"); err == nil {
			i++
			//stateQ[i] = r
			stateQ = append(stateQ, r)
		}
	}
	fmt.Println(stateQ)
	return State{} // For test only must change!!!
}

func bfsMove(s *State, dir string) (State, error) {
	fmt.Println("--------BFS-Gen :", dir)
	print(s.board)
	// Check what is last move and not to counter it
	// ie. if last what U then not D (return err)
	if s.sol[len(s.sol):] == "U" && dir == "D" {
		fmt.Println("No counter move")
		return State{}, errors.New("No counter")
	}
	if s.sol[len(s.sol):] == "D" && dir == "U" {
		fmt.Println("No counter move")
		return State{}, errors.New("No counter")
	}
	if s.sol[len(s.sol):] == "L" && dir == "R" {
		fmt.Println("No counter move")
		return State{}, errors.New("No counter")
	}
	if s.sol[len(s.sol):] == "R" && dir == "L" {
		fmt.Println("No counter move")
		return State{}, errors.New("No counter")
	}
	// Next move
	board := copyBoard(s.board)
	blank := copyBlank(s.blank)
	board, blank, err := move(board, blank, dir)
	if err != nil {
		fmt.Println("Can't move!!!!!")
		return State{}, errors.New("Can't move")
	}
	sol := s.sol + dir
	// Check is this state was past before. if is was return err
	return State{board: board, blank: blank, sol: sol}, err
}

func copyBlank(blank []int) []int {
	return []int{blank[0], blank[1]}
}

func copyBoard(board [][]int) [][]int {
	newBoard := make([][]int, 3)
	newBoard[0] = make([]int, 3)
	newBoard[1] = make([]int, 3)
	newBoard[2] = make([]int, 3)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			newBoard[i][j] = board[i][j]
		}
	}
	return newBoard
}

func move(board [][]int, blank []int, direction string) ([][]int, []int, error) {
	dir := strings.ToUpper(direction)
	newBoard := copyBoard(board)
	newBlank := copyBlank(blank)
	fmt.Println("Direction : ", dir)
	print(newBoard)
	fmt.Println("Blank :", newBlank)
	if canMove(newBlank, dir) {
		switch dir {
		case "U":
			moveU(newBoard, newBlank)
		case "D":
			moveD(newBoard, newBlank)
		case "L":
			moveL(newBoard, newBlank)
		case "R":
			moveR(newBoard, newBlank)
		}
	} else {
		return newBoard, newBlank, errors.New("Can't move")
	}
	fmt.Println("New Board")
	print(newBoard)
	fmt.Println("New Blank:", newBlank)
	fmt.Println("=========================")
	return newBoard, newBlank, nil
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
	fmt.Printf("RUNE %c\n", keepRune)
	//Move blank follow by sequence of keepRune.
	for _, direct := range keepRune {
		board, b, _ = move(board, b, string(direct))
	}

	return board, b
}

func print(board [][]int) {
	fmt.Println(board[0])
	fmt.Println(board[1])
	fmt.Println(board[2])
}

func returnToFront(board [][]int, step []string, tile []int) string {
	fmt.Println("Enter returnToFront")
	initpuzzle := []int{board[0][0], board[0][1], board[0][2],
		board[1][0], board[1][1], board[1][2],
		board[2][0], board[2][1], board[2][2]}

	arrStep := []Step{}
	for i := 0; i < len(step); i++ {
		move := new(Step)
		move.Tile = tile[i]
		move.Direction = step[i]
		arrStep = append(arrStep, *move)
	}
	fmt.Println(arrStep)

	res := &Response{
		Init:     initpuzzle,
		Sequence: arrStep,
	}
	jsonres, _ := json.Marshal(res)
	//	fmt.Println(string(jsonres))
	return string(jsonres)
}

func changeBlanktoTile(board [][]int, b []int, direct string) ([]string, []int) {
	//change blank move to tile move.
	var tile []int
	var move []string
	for i := 0; i < len(direct); i++ {
		switch string(direct[i]) {
		case "U":
			tile = append(tile, board[b[0]-1][b[1]])
			move = append(move, "D")
			board, b = moveU(board, b)
		case "D":
			tile = append(tile, board[b[0]+1][b[1]])
			move = append(move, "U")
			board, b = moveD(board, b)
		case "L":
			tile = append(tile, board[b[0]][b[1]-1])
			move = append(move, "R")
			board, b = moveL(board, b)
		case "R":
			tile = append(tile, board[b[0]][b[1]+1])
			move = append(move, "L")
			board, b = moveR(board, b)
		}
	}
	fmt.Println(tile)
	fmt.Println(move)
	return move, tile
}
