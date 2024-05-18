package knight

import (
	"math/rand"
	"sort"
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

type AccessibilitySortable struct {
	positions     []Position
	accessibility [8 * 8]int
}

func (a AccessibilitySortable) Len() int {
	return len(a.positions)
}

func (a AccessibilitySortable) Swap(i, j int) {
	a.positions[i], a.positions[j] = a.positions[j], a.positions[i]
}

func (a AccessibilitySortable) PositionToIndex(i int) int {
	return a.positions[i].X + a.positions[i].Y*8
}

func (a AccessibilitySortable) Less(i, j int) bool {
	return a.accessibility[a.PositionToIndex(i)] < a.accessibility[a.PositionToIndex(j)]
}

func (k *Knight) RankedPositions(pos Position) []Position {
	positions := []Position{}
	moves := Moves()
	for _, m := range moves {
		p := pos
		p.X += m.X
		p.Y += m.Y
		if k.IsValid(p) {
			positions = append(positions, p)
		}
	}

	sort.Sort(AccessibilitySortable{
		positions:     positions,
		accessibility: k.grid,
	})

	return positions
}
