package main

import (
	"container/heap"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
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
	board    [][]int
	blank    []int
	sol      string
	priority int
	index    int
}

type PQ []*State

func (pq PQ) Len() int { return len(pq) }

func (pq PQ) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority < pq[j].priority
}
func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PQ) Push(x interface{}) {
	n := len(*pq)
	item := x.(*State)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func (pq *PQ) update(item *State, board [][]int, blank []int, sol string, priority int) {
	item.board = board
	item.priority = priority
	heap.Fix(pq, item.index)
}

func main() {
	state := new(State)
	state.board = [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	state.blank = []int{1, 2}
	state.sol = "B"
	state.priority = 3
	state2 := new(State)
	state2.board = [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	state2.blank = []int{1, 2}
	state2.sol = "A"
	state2.priority = 4
	state3 := new(State)
	state3.board = [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	state3.blank = []int{1, 2}
	state3.sol = "C"
	state3.priority = 2
	pq := make(PQ, 0)
	heap.Init(&pq)
	heap.Push(&pq, state)
	heap.Push(&pq, state2)
	heap.Push(&pq, state3)
	for i := 0; i < len(pq); i++ {
		fmt.Println("IN-Q", (pq)[i])
	}

	//	fmt.Println("Hello AI")
	//	http.HandleFunc("/", homeHandler)
	//	http.ListenAndServe(":8000", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home Handler")
	init := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 0}}
	goal := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 0}}
	blank := []int{2, 2}

	fmt.Println("==========START===========")
	init, blank = randomPuzzle(init, blank, 50)
	start := copyBoard(init)
	fmt.Println("===========BFS===========")
	sol := bfs(goal, init, blank)
	fmt.Println("Solution : ", sol.sol)
	step, tile := changeBlanktoTile(init, blank, sol.sol)
	solutionJson := returnToFront(start, step, tile)
	fmt.Println(solutionJson)
	fmt.Println("******Goal******", goal)
	jsonData := struct {
		Json string
	}{
		Json: solutionJson,
	}

	t, _ := template.ParseFiles("app.html")
	t.Execute(w, jsonData)
	//fmt.Fprintf(w, "Hello world!")
}

// BFS Call this
func bfs(goal, init [][]int, blank []int) State { // return []Step
	var stateMap map[string]bool = make(map[string]bool)
	//	var stateQ []State = make([]State, 1)
	stateQ := make(PQ, 1)
	heap.Init(&stateQ)
	state := new(State)
	state.board = init
	state.sol = ""
	state.blank = blank
	state.priority = 1
	fmt.Println("StateQ : ", &stateQ)
	fmt.Println("state :", state)
	heap.Push(&stateQ, state)
	fmt.Println("HEAP HEAP")
	//	stateQ[0].board = init
	//	stateQ[0].sol = ""
	//	stateQ[0].blank = blank
	//fmt.Println(stateQ)
	i := 0
	//for {
	for q := 0; ; q++ {
		// Dequeue, check with goal
		if checkIdentical(stateQ[q].board, goal) {
			//fmt.Println("---------------SUCCESS------------------")
			//fmt.Println(stateQ[q])
			//print(stateQ[q].board)
			return *stateQ[q]
		}
		// Move UDLR , Enqueue
		if u, err := bfsMove(stateQ[q], "U"); err == nil {
			i++
			stateQ = bfsAppend(stateQ, stateMap, u)
		}
		if d, err := bfsMove(stateQ[q], "D"); err == nil {
			i++
			stateQ = bfsAppend(stateQ, stateMap, d)
		}
		if l, err := bfsMove(stateQ[q], "L"); err == nil {
			i++
			stateQ = bfsAppend(stateQ, stateMap, l)
		}
		if r, err := bfsMove(stateQ[q], "R"); err == nil {
			i++
			stateQ = bfsAppend(stateQ, stateMap, r)
		}
	}
	//fmt.Println(stateQ)
	return State{} // For test only must change!!!
}

func bfsAppend(stateQ PQ, hashMap map[string]bool, newState State) PQ {
	//If this state was past before then will not append its
	key := fmt.Sprint(newState.board)
	wasPast := hashMap[key]
	if wasPast == true {
		//		fmt.Println("Was past")
		return stateQ
	}
	hashMap[key] = true
	return append(stateQ, &newState)
}

func bfsMove(s *State, dir string) (State, error) {
	//fmt.Println("--------BFS-Gen :", dir)
	//print(s.board)
	// Check what is last move and not to counter it
	// ie. if last what U then not D (return err)
	if s.sol[len(s.sol):] == "U" && dir == "D" {
		//fmt.Println("No counter move")
		return State{}, errors.New("No counter")
	}
	if s.sol[len(s.sol):] == "D" && dir == "U" {
		//fmt.Println("No counter move")
		return State{}, errors.New("No counter")
	}
	if s.sol[len(s.sol):] == "L" && dir == "R" {
		//fmt.Println("No counter move")
		return State{}, errors.New("No counter")
	}
	if s.sol[len(s.sol):] == "R" && dir == "L" {
		//fmt.Println("No counter move")
		return State{}, errors.New("No counter")
	}
	// Next move
	board := copyBoard(s.board)
	blank := copyBlank(s.blank)
	board, blank, err := move(board, blank, dir)
	if err != nil {
		//fmt.Println("Can't move!!!!!")
		return State{}, errors.New("Can't move")
	}
	sol := s.sol + dir
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
	//fmt.Println("Direction : ", dir)
	//print(newBoard)
	//fmt.Println("Blank :", newBlank)
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
	//fmt.Println("New Board")
	//print(newBoard)
	//fmt.Println("New Blank:", newBlank)
	//fmt.Println("=========================")
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
	//fmt.Println("Enter returnToFront")
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
	///fmt.Println(arrStep)

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
	//fmt.Println(tile)
	//fmt.Println(move)
	return move, tile
}
