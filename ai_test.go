package tic_tac_toe

import (
	"testing"
)

func TestPlay(t *testing.T) {
	b := &Board{}
	b.Start()
	check(Play(b, X), t)
	if b.FreePositions() != 8 {
		t.Errorf("free positions should be 8 but is %d", b.FreePositions())
	}
}

func TestPlayMustUseFreePosition(t *testing.T) {
	b := &Board{}
	b.Start()
	Play(b, X)
	err := Play(b, X)
	if err != nil {
		t.Error(err)
	}
}

func TestPlayStopWhenBoardFull(t *testing.T) {
	b := &Board{}
	b.Start()
	for i := 0; i < 9; i++ {
		err := Play(b, X)
		if err != nil {
			t.Error(err)
		}
	}
	if b.FreePositions() != 0 {
		t.Errorf("board should be full, but is %d", b.FreePositions())
	}
	err := Play(b, X)
	if err == nil {
		t.Error("board is full")
	}

}

func TestFirstPlayMustBeACornerToGuaranteeWinOrDraw(t *testing.T) {
	b := Board{}
	b.Start()

	check(Play(&b, X), t)
	p := b.LastMark()
	if !b.Corner(p.row, p.col) {
		t.Error("first play should be a corner")
	}
}

func TestSecondPlayResponseToCornerWithACenter(t *testing.T) {
	b := Board{}
	b.Start()
	check(b.Put(X, 0, 0), t)

	check(Play(&b, O), t)
	p := b.LastMark()
	if p.row != 1 || p.col != 1 {
		t.Errorf("second play should be a center but got a %d %d", p.row, p.col)
	}
}

func TestSecondPlayResponseToCenterWithACorner(t *testing.T) {
	b := Board{}
	b.Start()
	check(b.Put(X, 1, 1), t)

	check(Play(&b, O), t)
	p := b.LastMark()
	if !b.Corner(p.row, p.col) {
		t.Errorf("second play should be a corner but got a %d %d", p.row, p.col)
	}
}

func TestSecondPlayResponseToEdgeMustBeACenter(t *testing.T) {
	b := Board{}
	b.Start()
	check(b.Put(X, 0, 1), t)
	check(Play(&b, O), t)
	p := b.LastMark()
	if p.row != 1 || p.col != 1 {
		t.Errorf("second play response to edge should be a center but got a %d %d", p.row, p.col)
	}
}

func TestPlayWinIfInFork(t *testing.T) {
	b := Board{}
	b.Start()
	b.Put(X, 0, 0)
	b.Put(X, 1, 1)
	check(Play(&b, X), t)
	p := b.LastMark()
	if p.row != 2 || p.col != 2 {
		t.Errorf("play should be put in position 2, 2 but was %d, %d instead",
			p.row, p.col)
	}
}
