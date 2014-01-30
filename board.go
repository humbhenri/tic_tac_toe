package tic_tac_toe

import (
	"strconv"
)

type Mark int

func (m Mark) String() string {
	switch m {
	case O:
		return "O"
	case X:
		return "X"
	}
	return "ERROR"
}

const (
	O Mark = iota
	X
)

type Pos struct {
	m   Mark
	row int
	col int
}

type Board struct {
	free int
	pos  []Pos
}

func (b *Board) Start() {
	b.free = 9
	b.pos = []Pos{}
}

func (b *Board) FreePositions() int {
	return b.free
}

func (b *Board) Put(m Mark, i, j int) error {
	if i < 0 || i > 2 || j < 0 || j > 2 {
		return InvalidError{i, j}
	}
	if b.free == 0 {
		return FullError{}
	}
	for _, p := range b.pos {
		if p.row == i && p.col == j {
			return OccupiedError{m, i, j}
		}
	}
	b.pos = append(b.pos, Pos{m, i, j})
	b.free--
	return nil
}

func (b *Board) Block(m Mark) bool {

	for i := 0; i < len(b.pos); i++ {
		for j := 0; j < len(b.pos); j++ {
			if i != j {
				p1 := b.pos[i]
				p2 := b.pos[j]
				if p1.m == p2.m &&
					(p1.col == p2.col ||
						p1.row == p2.row ||
						(p1.row == p1.col && p2.row == p2.col) ||
						(p1.row == p2.col && p1.col == p2.row) ||
						(p1.row == 1 && p1.col == 1 && p2.row == 0 && p2.col == 2) ||
						(p1.row == 1 && p1.col == 1 && p2.row == 2 && p2.col == 0)) {
					return true
				}
			}
		}
	}

	return false

}

type OccupiedError struct {
	m   Mark
	row int
	col int
}

func (e OccupiedError) Error() string {
	return "Put " + e.m.String() + " at row " + strconv.Itoa(e.row) + " and column " + strconv.Itoa(e.col) + ": already occupied"
}

type FullError struct {
}

func (e FullError) Error() string {
	return "Board is full"
}

type InvalidError struct {
	row int
	col int
}

func (e InvalidError) Error() string {
	return "Invalid move: row = " + strconv.Itoa(e.row) + ", col = " +
		strconv.Itoa(e.col)
}
