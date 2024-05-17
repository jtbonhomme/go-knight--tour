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

const (
	North = iota
	East
	South
	West
)

const (
	Left = iota
	Right
)

type Position struct {
	X int
	Y int
}

type Move struct {
	Direction int
	Turn      int
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

func (k *Knight) Moves() []Move {
	moves := []Move{}

	for d := 0; d < 4; d++ {
		for t := 0; t < 2; t++ {
			moves = append(moves, Move{d, t})
		}
	}

	return moves
}

func (k *Knight) Tour() int {
	return k.tour
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

	for {
		// pick randow move
		p = Position{rand.Intn(8), rand.Intn(8)}
		if !k.PositionExists(p) {
			break
		}
	}

	k.Positions = append(k.Positions, p)

	return k.Solve()
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
