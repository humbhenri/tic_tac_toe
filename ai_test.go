package tic_tac_toe

import "testing"

func TestPlay(t *testing.T) {
	b := &Board{}
	b.Start()
	Play(b, X)
	if b.FreePositions() != 8 {
		t.Error("play")
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
