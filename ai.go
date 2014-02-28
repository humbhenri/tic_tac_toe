package tic_tac_toe

import (
	"math/rand"
	"time"
)

const (
	FIRST          = 9
	SECOND         = 8
	SIX_MARKS_FREE = 6
)

// Play make a play in board b with mark m
func Play(b *Board, m Mark) error {
	switch b.FreePositions() {
	case FIRST:
		return markACorner(b, m)
	case SECOND:
		return doSecondPlay(b, m)
	case SIX_MARKS_FREE:
		if b.MarkOf(1, 1) == m && ((b.MarkOf(0, 0) == m.Opposite() && b.MarkOf(2, 2) == m.Opposite()) || (b.MarkOf(0, 2) == m.Opposite() && b.MarkOf(2, 0) == m.Opposite())) {
			return markEdge(b, m)
		}
	}

	if p := b.Fork(m); p != nil {
		return b.Put(m, p.Row, p.Col)
	}

	return markRandom(b, m)
}

func doSecondPlay(b *Board, m Mark) error {
	p := b.LastMark()
	if b.Corner(p.Row, p.Col) || b.Edge(p.Row, p.Col) {
		return markCenter(b, m)
	}
	if p.Row == 1 && p.Col == 1 {
		return markACorner(b, m)
	}

	return markRandom(b, m)
}

func markACorner(b *Board, m Mark) error {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	openning := []int{0, 2}
	return b.Put(m, openning[r.Intn(2)], openning[r.Intn(2)])
}

func markCenter(b *Board, m Mark) error {
	return b.Put(m, 1, 1)
}

func markRandom(b *Board, m Mark) error {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for running := true; running; {
		err := b.Put(m, r.Intn(3), r.Intn(3))
		if err != nil {
			switch err.(type) {
			case FullError, InvalidError:
				return err
			case OccupiedError:
				running = true
			}
		} else {
			running = false
		}
	}
	return nil
}

func markEdge(b *Board, m Mark) error {
	edges := [][]int{{0, 1}, {1, 0}, {1, 2}, {2, 1}}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	p := edges[r.Intn(len(edges))]
	return b.Put(m, p[0], p[1])
}
