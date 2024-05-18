package knight

func (k *Knight) OutOfRange(p Position) bool {
	if p.X >= 8 || p.X < 0 || p.Y >= 8 || p.Y < 0 {
		return true
	}

	return false
}

func (k *Knight) Invalidate(pos Position) {
	var p Position

	if k.OutOfRange(pos) || k.grid[pos.X+pos.Y*8] == 0 {
		return
	}

	moves := Moves()
	for _, m := range moves {
		p.X = pos.X + m.X
		p.Y = pos.Y + m.Y

		if k.OutOfRange(p) {
			continue
		}

		if k.grid[p.X+p.Y*8] != 0 {
			k.grid[p.X+p.Y*8]--
		}
	}

	k.grid[pos.X+pos.Y*8] = 0
}

func (k *Knight) IsValid(p Position) bool {
	if k.OutOfRange(p) {
		return false
	}

	return k.grid[p.X+p.Y*8] != 0
}
