package tic_tac_toe

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

// Mark datatype
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

// Mark types
const (
	None Mark = iota
	O
	X
)

// Pos represents a mark in the board
type Pos struct {
	M   Mark
	Row int
	Col int
}

func (p *Pos) String() string {
	return fmt.Sprintf("[%s, %d, %d]", p.M.String(), p.Row, p.Col)
}

// Board represents the board game
type Board struct {
	Free int
	Pos  []Pos
}

// Start resets the board
func (b *Board) Start() {
	b.Free = 9
	b.Pos = []Pos{}
}

// FreePositions return the number of free positions of the board
func (b *Board) FreePositions() int {
	return b.Free
}

// Put places the mark m in the row in and column j
func (b *Board) Put(m Mark, i, j int) error {
	if i < 0 || i > 2 || j < 0 || j > 2 {
		return InvalidError{i, j}
	}
	if b.Free == 0 {
		return FullError{}
	}
	if b.Occupied(i, j) {
		return OccupiedError{m, i, j}
	}
	b.Pos = append(b.Pos, Pos{m, i, j})
	b.Free--
	return nil
}

// Occupied returns true if there is a mark different of None in row i and column j
func (b *Board) Occupied(i, j int) bool {
	for _, p := range b.Pos {
		if p.Row == i && p.Col == j && p.M != None {
			return true
		}
	}
	return false
}

// LastMark return the last position occupied
func (b *Board) LastMark() *Pos {
	if len(b.Pos) == 0 {
		return nil
	}
	return &b.Pos[len(b.Pos)-1]
}

// Corner returns true if row i and column j represents a corner of the board
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
	for _, p := range b.Pos {
		grid[p.Row][p.Col] = p.M.String()
	}
	var rows []string
	for i := 0; i < 3; i++ {
		rows = append(rows, strings.Join(grid[i][:], " | "))
	}
	buf.WriteString(strings.Join(rows, "\n---------\n"))
	return buf.String()
}

// Edge returns true if row i and column j represents an edge of the board
func (b *Board) Edge(i, j int) bool {
	return (i == 0 && j == 1) ||
		(i == 1 && (j == 0 || j == 2)) ||
		(i == 2 && j == 1)
}

// Win returns the mark of the player if he/she wins, or None otherwise
func (b *Board) Win() Mark {
	for i := 0; i < len(b.Pos); i++ {
		for j := 0; j < len(b.Pos); j++ {
			if i == j {
				continue
			}
			for k := 0; k < len(b.Pos); k++ {
				if j == k || i == k {
					continue
				}
				p1 := b.Pos[i]
				p2 := b.Pos[j]
				p3 := b.Pos[k]
				if p1.M == p2.M && p1.M == p3.M {
					if p1.Row == p2.Row && p1.Row == p3.Row {
						return p1.M
					}
					if p1.Col == p2.Col && p1.Col == p3.Col {
						return p1.M
					}
					if p1.Row == 0 && p1.Col == 0 && p2.Row == 1 && p2.Col == 1 && p3.Row == 2 && p3.Col == 2 {
						return p1.M
					}
					if p1.Row == 0 && p1.Col == 2 && p2.Row == 1 && p2.Col == 1 && p3.Row == 2 && p3.Col == 0 {
						return p1.M
					}
				}

			}
		}
	}

	return None
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

// Fork return the position if it's possible to win, or nil otherwise
func (b *Board) Fork(m Mark) (p *Pos) {
	var diff [3][3]int
	diff[1][2] = 0
	diff[2][1] = 0
	diff[0][2] = 1
	diff[2][0] = 1
	diff[0][1] = 2
	diff[1][0] = 2

	for i := 0; i < len(b.Pos); i++ {
		for j := 0; j < len(b.Pos); j++ {
			if i == j {
				continue
			}
			p1 := b.Pos[i]
			p2 := b.Pos[j]
			if p1.M == p2.M {
				if p1.Row == p2.Row {
					p = &Pos{p1.M, p1.Row, diff[p1.Col][p2.Col]}
					if !b.Occupied(p.Row, p.Col) {
						return
					}
				}
				if p1.Col == p2.Col {
					p = &Pos{p1.M, diff[p1.Row][p2.Row], p1.Col}
					if !b.Occupied(p.Row, p.Col) {
						return
					}
				}
				if abs(p1.Row-p2.Row) == abs(p1.Col-p2.Col) {
					p = &Pos{p1.M, diff[p1.Row][p2.Row], diff[p1.Col][p2.Col]}
					if !b.Occupied(p.Row, p.Col) {
						return
					}
				}
			}
		}
	}

	return nil
}

// OccupiedError error if position occupied
type OccupiedError struct {
	m   Mark
	row int
	col int
}

func (e OccupiedError) Error() string {
	return "Put " + e.m.String() + " at row " + strconv.Itoa(e.row) + " and column " + strconv.Itoa(e.col) + ": already occupied"
}

// FullError error when board is full
type FullError struct {
}

func (e FullError) Error() string {
	return "Board is full"
}

// InvalidError error when the move is invalid
type InvalidError struct {
	row int
	col int
}

func (e InvalidError) Error() string {
	return "Invalid move: row = " + strconv.Itoa(e.row) + ", col = " +
		strconv.Itoa(e.col)
}
