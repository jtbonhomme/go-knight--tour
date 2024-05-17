package knight

import (
	"log"
	"math/rand"
	"sort"
	"time"
)

type Position struct {
	X int
	Y int
}

type Move struct {
	X         int
	Y         int
	Direction string
}

type Knight struct {
	Positions      []Position
	speed          int
	tour           int
	implementation string
}

func New(speed int, implementation string) *Knight {
	return &Knight{
		Positions:      []Position{},
		speed:          speed,
		implementation: implementation,
	}
}

func Moves() []Move {
	moves := []Move{}

	moves = append(moves, Move{-1, -2, "north left"},
		Move{1, -2, "north right"},
		Move{-1, 2, "south left"},
		Move{1, 2, "south right"},
		Move{2, -1, "east left"},
		Move{2, 1, "east right"},
		Move{-2, 1, "west left"},
		Move{-2, -1, "west right"})

	return moves
}

func RandomMoves() []Move {
	m := Moves()
	rand.Shuffle(len(m), func(i, j int) { m[i], m[j] = m[j], m[i] })
	return m
}

func (k *Knight) Tour() int {
	return k.tour
}

func OutOfRange(p Position) bool {
	if p.X >= 8 || p.X < 0 || p.Y >= 8 || p.Y < 0 {
		return true
	}

	return false
}

func PositionExists(p Position, positions []Position) bool {
	for _, kp := range positions {
		if p.X == kp.X && p.Y == kp.Y {
			return true
		}
	}

	return false
}

func (k *Knight) NaiveSolver(tour int, positions []Position) bool {
	time.Sleep(time.Millisecond * time.Duration(k.speed))

	if tour == 8*8 {
		log.Println("win!")
		return true
	}

	moves := RandomMoves()
	// pick successively random moves
	for _, m := range moves {
		p := positions[len(positions)-1]
		p.X += m.X
		p.Y += m.Y
		if !PositionExists(p, positions) && !OutOfRange(p) {
			positions = append(positions, p)
			k.Update(positions, tour)
			return k.NaiveSolver(tour+1, positions)
		}
	}

	return false
}

func (k *Knight) Update(positions []Position, tour int) {
	k.Positions = positions
	k.tour = tour
}

func (k *Knight) BacktrackingSolver(tour int, positions []Position) bool {
	time.Sleep(time.Millisecond * time.Duration(k.speed))

	if tour == 8*8 {
		log.Println("win!")
		return true
	}

	moves := RandomMoves()
	// pick successively random moves
	for _, m := range moves {
		p := positions[len(positions)-1]
		p.X += m.X
		p.Y += m.Y
		if !PositionExists(p, positions) && !OutOfRange(p) {
			positions = append(positions, p)
			k.Update(positions, tour)
			if k.BacktrackingSolver(tour+1, positions) {
				return true
			}
		}
	}

	return false
}

type ByAccessible []Position

func (a ByAccessible) Len() int           { return len(a) }
func (a ByAccessible) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAccessible) Less(i, j int) bool { return a[i].Accessible() < a[j].Accessible() }

func RankedPositions(pos Position) []Position {
	positions := []Position{}
	moves := Moves()
	for _, m := range moves {
		p := pos
		p.X += m.X
		p.Y += m.Y
		if !PositionExists(p, positions) && !OutOfRange(p) {
			positions = append(positions, p)
		}
	}
	sort.Sort(ByAccessible(positions))
	return positions
}

func (k *Knight) OptimizedSolver(tour int, positions []Position) bool {
	time.Sleep(time.Millisecond * time.Duration(k.speed))
	k.tour = tour
	if tour == 8*8 {
		log.Println("win!")
		return true
	}

	currentPosition := positions[len(positions)-1]
	rankedPositions := RankedPositions(currentPosition)
	// pick best moves
	for _, p := range rankedPositions {
		if !PositionExists(p, positions) && !OutOfRange(p) {
			positions = append(positions, p)
			k.Positions = positions
			if k.OptimizedSolver(tour+1, positions) {
				return true
			}
		}
	}

	return false
}

func (k *Knight) Run() {
	var result bool
	// pick initial position
	x := rand.Intn(8)
	y := rand.Intn(8)
	//x := 4
	//y := 6
	k.Positions = append(k.Positions, Position{x, y})
	k.tour = 1

	go func() {
		log.Println("start knight's tour solver: starting from ", Position{x, y})
		switch k.implementation {
		case "naive":
			result = k.NaiveSolver(k.tour, k.Positions)
		case "backtracking":
			result = k.BacktrackingSolver(k.tour, k.Positions)
		case "optimized":
			result = k.OptimizedSolver(k.tour, k.Positions)
		default:
			log.Fatalf("%s implementation does not exist", k.implementation)
		}
		log.Printf("solver result: %v at position %v", result, k.Positions[len(k.Positions)-1])
	}()

	log.Printf("knight: run")
}

func (p Position) Accessible() int {
	var accessible = [8 * 8]int{
		2, 3, 4, 4, 4, 4, 3, 2,
		3, 4, 6, 6, 6, 6, 4, 3,
		4, 6, 8, 8, 8, 8, 6, 4,
		4, 6, 8, 8, 8, 8, 6, 4,
		4, 6, 8, 8, 8, 8, 6, 4,
		4, 6, 8, 8, 8, 8, 6, 4,
		3, 4, 6, 6, 6, 6, 4, 3,
		2, 3, 4, 4, 4, 4, 3, 2}

	if p.X < 0 || p.Y < 0 || p.X >= 8 || p.Y >= 8 {
		return -1
	}

	return accessible[p.X+p.Y*8]
}
