package tic_tac_toe

import (
	//"fmt"
	"math/rand"
	"time"
)

func Play(b *Board, m Mark) error {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	running := true
	for running {
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
