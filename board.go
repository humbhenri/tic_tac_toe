package tic_tac_toe

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

type Mark int

func (m Mark) String() string {
	switch m {
	case O:
		return "O"
	case X:
		return "X"
	}
	return " "
}

const (
	None Mark = iota
	O
	X
)

type Pos struct {
	m   Mark
	row int
	col int
}

func (p *Pos) String() string {
	return fmt.Sprintf("[%s, %d, %d]", p.m.String(), p.row, p.col)
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

func (b *Board) LastMark() *Pos {
	if len(b.pos) == 0 {
		return nil
	}
	return &b.pos[len(b.pos)-1]
}

func (b *Board) Corner(i, j int) bool {
	return (i == 0 && (j == 0 || j == 2)) || (i == 2 && (j == 0 || j == 2))
}

func (b *Board) String() string {
	var buf bytes.Buffer
	var grid [3][3]string
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			grid[i][j] = None.String()
		}
	}
	for _, p := range b.pos {
		grid[p.row][p.col] = p.m.String()
	}
	var rows []string
	for i := 0; i < 3; i++ {
		rows = append(rows, strings.Join(grid[i][:], " | "))
	}
	buf.WriteString(strings.Join(rows, "\n---------\n"))
	return buf.String()
}

func (b *Board) Edge(i, j int) bool {
	return (i == 0 && j == 1) ||
		(i == 1 && (j == 0 || j == 2)) ||
		(i == 2 && j == 1)
}

func (b *Board) Win() Mark {
	for i := 0; i < len(b.pos); i++ {
		for j := 0; j < len(b.pos); j++ {
			if i == j {
				continue
			}
			for k := 0; k < len(b.pos); k++ {
				if j == k || i == k {
					continue
				}
				p1 := b.pos[i]
				p2 := b.pos[j]
				p3 := b.pos[k]
				if p1.m == p2.m && p1.m == p3.m {
					if p1.row == p2.row && p1.row == p3.row {
						return p1.m
					}
					if p1.col == p2.col && p1.col == p3.col {
						return p1.m
					}
					if p1.row == 0 && p1.col == 0 && p2.row == 1 && p2.col == 1 && p3.row == 2 && p3.col == 2 {
						return p1.m
					}
					if p1.row == 0 && p1.col == 2 && p2.row == 1 && p2.col == 1 && p3.row == 2 && p3.col == 0 {
						return p1.m
					}
				}

			}
		}
	}

	return None
}

func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func (b *Board) Fork(m Mark) *Pos {
	var diff [3][3]int
	diff[1][2] = 0
	diff[2][1] = 0
	diff[0][2] = 1
	diff[2][0] = 1
	diff[0][1] = 2
	diff[1][0] = 2

	for i := 0; i < len(b.pos); i++ {
		for j := 0; j < len(b.pos); j++ {
			if i == j {
				continue
			}
			p1 := b.pos[i]
			p2 := b.pos[j]
			if p1.m == p2.m {
				if p1.row == p2.row {
					return &Pos{p1.m, p1.row, diff[p1.col][p2.col]}
				}
				if p1.col == p2.col {
					return &Pos{p1.m, diff[p1.row][p2.row], p1.col}
				}
				if Abs(p1.row-p2.row) == Abs(p1.col-p2.col) {
					return &Pos{p1.m, diff[p1.row][p2.row], diff[p1.col][p2.col]}
				}
			}
		}
	}
	return nil
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
