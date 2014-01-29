package tic_tac_toe

import "testing"

func TestStart(t *testing.T) {
	b := &Board{}
	b.Start()
	if b.FreePositions() != 9 {
		t.Errorf("start failed")
	}

}

func TestMark(t *testing.T) {
	b := &Board{}
	b.Start()
	b.Put(O, 0, 0)
	if b.FreePositions() != 8 {
		t.Error("put failed")
	}
}

func TestPutPositionOccupied(t *testing.T) {
	b := &Board{}
	b.Start()
	b.Put(O, 0, 0)
	err := b.Put(X, 0, 0)
	if err == nil {
		t.Error("test position occupied failed")
	}

	err = b.Put(X, 1, 1)
	if err != nil {
		t.Error(err)
	}
	err = b.Put(X, 1, 1)
	if err == nil {
		t.Error(err)
	}

}

func TestPutCannotPutMoreThan9Marks(t *testing.T) {
	b := &Board{}
	b.Start()
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			err := b.Put(X, i, j)
			if err != nil {
				t.Error(err)
			}
		}
	}

	err := b.Put(X, 10, 0)
	if err == nil {
		t.Error("can put more than 9")
	}

}

func TestPutValidRowAndCol(t *testing.T) {
	b := &Board{}
	b.Start()
	err := b.Put(O, 1, 2)
	if err != nil {
		t.Error("should be ok")
	}
	err = b.Put(O, 1, 4)
	if err == nil {
		t.Error("rows and columns valid should be between 0 and 2")
	}

}
