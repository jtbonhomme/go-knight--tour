package knight

import (
	"log"
	"math/rand"
	"time"
)

const (
	Started = iota
	Running
	Paused
	Complete
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
	Positions []Position
	state     int
	speed     int
	tour      int
}

func New(speed int) *Knight {
	return &Knight{
		state:     Started,
		Positions: []Position{},
		speed:     speed,
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

func (k *Knight) OutOfRange(p Position) bool {
	if p.X >= 8 || p.X < 0 || p.Y >= 8 || p.Y < 0 {
		return true
	}

	return false
}

func (k *Knight) PositionExists(p Position) bool {
	for _, kp := range k.Positions {
		if p.X == kp.X && p.Y == kp.Y {
			return true
		}
	}

	return false
}

func (k *Knight) Solve() bool {
	var p Position

	time.Sleep(time.Millisecond * time.Duration(k.speed))
	k.tour++

	log.Printf("knight's tour: %d", k.tour)
	if k.tour == 8*8 {
		return false
	}

	moves := RandomMoves()
	// pick successively random moves
	for _, m := range moves {
		p = k.Positions[len(k.Positions)-1]
		p.X += m.X
		p.Y += m.Y
		if !k.PositionExists(p) && !k.OutOfRange(p) {

			k.Positions = append(k.Positions, p)
			return k.Solve()
		}
	}

	log.Println("no solution")
	return false
}

func (k *Knight) Complete() {
	if k.state != Running && k.state != Paused {
		return
	}

	k.state = Complete
}

func (k *Knight) Run() {
	if k.state != Started && k.state != Paused {
		return
	}

	// pick initial position
	if len(k.Positions) == 0 {
		x := rand.Intn(8)
		y := rand.Intn(8)
		k.Positions = append(k.Positions, Position{x, y})
	}

	go func() {
		log.Println("start knight's tour solver")
		k.Solve()
	}()

	k.state = Running
	log.Printf("knight: run")

}

func (k *Knight) Pause() {
	if k.state != Running {
		return
	}

	k.state = Paused
	log.Printf("knight: pause")

}
