package tic_tac_toe

import (
	"fmt"
)

func ShowStats(b *Board) {
	fmt.Print(b.String())
	p := b.LastMark()
	if p != nil {
		fmt.Printf("\t%s player marked [%d, %d]\n", p.m.String(), p.row, p.col)
	}
	fmt.Println("\n\n")
}

func ShowWinner(m Mark) {
	fmt.Printf("Player %s win!\n", m.String())
}

func ShowDraw() {
	fmt.Println("Draw !!")
}
