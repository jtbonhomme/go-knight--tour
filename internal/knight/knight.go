package knight

import (
	"log"
	"math/rand"
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

func (k *Knight) Run() {
	var result bool
	// pick initial position
	x := rand.Intn(8)
	y := rand.Intn(8)
	k.Positions = append(k.Positions, Position{x, y})

	go func() {
		log.Println("start knight's tour solver")
		switch k.implementation {
		case "naive":
			result = k.NaiveSolver(1, k.Positions)
		case "backtracking":
			result = k.BacktrackingSolver(1, k.Positions)
		default:
			log.Fatalf("%s implementation does not exist", k.implementation)
		}
		log.Printf("solver result: %v", result)
	}()

	log.Printf("knight: run")
}
