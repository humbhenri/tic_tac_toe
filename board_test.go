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

func TestBoardInBlock(t *testing.T) {

	testBlock := func(pos []Pos) bool {
		b := &Board{}
		b.Start()
		for _, p := range pos {
			err := b.Put(p.m, p.row, p.col)
			if err != nil {
				t.Fatal(err)
			}
		}
		return b.Block(O)
	}

	assertTrue(testBlock([]Pos{Pos{O, 0, 0}, Pos{O, 0, 1}}), "should be in block 0, 0 and 0, 1", t)
	assertTrue(testBlock([]Pos{Pos{O, 0, 0}, Pos{O, 1, 0}}), "should be in block 0, 0 and 1, 0", t)
	assertFalse(testBlock([]Pos{Pos{O, 0, 0}, Pos{O, 1, 2}}), "should not be in block 0, 0 and 1, 2", t)
	assertFalse(testBlock([]Pos{Pos{O, 1, 0}}), "should not be in block 1, 0 ", t)
	assertFalse(testBlock([]Pos{Pos{O, 0, 0}, Pos{X, 1, 0}}), "should not be in block different marks", t)
	assertTrue(testBlock([]Pos{Pos{O, 0, 0}, Pos{O, 1, 0}, Pos{O, 2, 2}}), "should be in block tree marks", t)
	assertTrue(testBlock([]Pos{Pos{O, 0, 0}, Pos{O, 1, 1}}), "should be are in block if same marks in diagonal", t)
	assertTrue(testBlock([]Pos{Pos{O, 0, 2}, Pos{O, 2, 0}}), "should be are in block if same marks in diagonal - II", t)
	assertTrue(testBlock([]Pos{Pos{O, 0, 2}, Pos{O, 1, 1}}), "should be are in block if same marks in diagonal - III", t)
	assertTrue(testBlock([]Pos{Pos{X, 1, 1}, Pos{X, 2, 0}}), "should be are in block if same marks in diagonal - IV", t)
}

func assertTrue(cond bool, msg string, t *testing.T) {
	if !cond {
		t.Error(msg)
	}
}

func assertFalse(cond bool, msg string, t *testing.T) {
	if cond {
		t.Error(msg)
	}
}
