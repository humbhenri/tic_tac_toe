package tic_tac_toe

import (
	"math/rand"
	"time"
)

func Play(b *Board, m Mark) error {
	switch b.FreePositions() {
	case 9:
		return markACorner(b, m)
	case 8:
		return doSecondPlay(b, m)
	}

	if p := b.Fork(m); p != nil {
		return b.Put(m, p.row, p.col)
	}

	return markRandom(b, m)
}

func doSecondPlay(b *Board, m Mark) error {
	p := b.LastMark()
	if b.Corner(p.row, p.col) || b.Edge(p.row, p.col) {
		return markCenter(b, m)
	}
	if p.row == 1 && p.col == 1 {
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
