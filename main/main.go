package main

import (
	. "github.com/humbhenri/tic_tac_toe"
)

func main() {
	b := &Board{}
	b.Start()
	player := X
	for {
		ShowStats(b)

		win := b.Win()
		if win != None {
			ShowWinner(win)
			break
		}

		if b.FreePositions() == 0 {
			ShowDraw()
			break
		}

		Play(b, player)

		if player == X {
			player = O
		} else {
			player = X
		}
	}
}
