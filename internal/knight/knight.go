package knight

import (
	"log"
	"math/rand"
)

var accessibility = [8 * 8]int{
	2, 3, 4, 4, 4, 4, 3, 2,
	3, 4, 6, 6, 6, 6, 4, 3,
	4, 6, 8, 8, 8, 8, 6, 4,
	4, 6, 8, 8, 8, 8, 6, 4,
	4, 6, 8, 8, 8, 8, 6, 4,
	4, 6, 8, 8, 8, 8, 6, 4,
	3, 4, 6, 6, 6, 6, 4, 3,
	2, 3, 4, 4, 4, 4, 3, 2}

type Knight struct {
	Positions      []Position
	speed          int
	tour           int
	implementation string
	grid           [8 * 8]int
}

func New(speed int, implementation string) *Knight {
	return &Knight{
		Positions:      []Position{},
		speed:          speed,
		implementation: implementation,
		grid:           accessibility,
	}
}

func (k *Knight) Tour() int {
	return k.tour
}

func (k *Knight) Update(positions []Position, tour int) {
	k.Positions = positions
	k.tour = tour
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

func (p Position) Distance(q Position) int {
	dx := q.X - p.X
	dy := q.Y - p.Y
	if dx < 0 {
		dx = -dx
	}
	if dy < 0 {
		dy = -dy
	}
	return dx + dy
}
