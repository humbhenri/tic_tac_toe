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

func TestLastMark(t *testing.T) {
	b := &Board{}
	b.Start()
	check(b.Put(O, 1, 2), t)
	p := b.LastMark()
	assertTrue(p.m == O, "last mark failed - mark", t)
	assertTrue(p.col == 2, "last mark failed - col", t)
	assertTrue(p.row == 1, "last mark failed - row", t)

}

func TestCorner(t *testing.T) {
	b := Board{}
	b.Start()
	assertFalse(b.Corner(1, 2), "failed corner 1, 2", t)
	assertTrue(b.Corner(0, 0), "failed corner 0, 0", t)
	assertFalse(b.Corner(1, 1), "failed corner 1, 1", t)
	assertTrue(b.Corner(2, 2), "failed corner 2, 2", t)
	assertTrue(b.Corner(2, 0), "failed corner 2, 0", t)
	assertTrue(b.Corner(0, 2), "failed corner 0, 2", t)
}

func TestEdge(t *testing.T) {
	b := Board{}
	b.Start()
	assertTrue(b.Edge(0, 1), "should be edge 0, 1", t)
	assertTrue(b.Edge(1, 0), "should be edge 1, 0", t)
	assertTrue(b.Edge(1, 2), "should be edge 1, 2", t)
	assertTrue(b.Edge(2, 1), "should be edge 2, 1", t)
	assertFalse(b.Edge(0, 0), "should not be edge 0, 0", t)
}

func TestWin(t *testing.T) {
	b := Board{}
	b.Start()
	b.Put(X, 0, 0)
	b.Put(X, 1, 1)
	b.Put(X, 2, 2)
	assertTrue(b.Win() == X, "X wins", t)

	b = Board{}
	b.Start()
	b.Put(O, 0, 0)
	b.Put(O, 1, 1)
	b.Put(O, 2, 2)
	assertTrue(b.Win() == O, "O wins", t)

	b = Board{}
	b.Start()
	b.Put(O, 0, 0)
	b.Put(X, 1, 2)
	b.Put(O, 0, 1)
	b.Put(O, 0, 2)
	assertTrue(b.Win() == O, "II - O wins", t)

	b = Board{}
	b.Start()
	b.Put(O, 0, 0)
	b.Put(X, 1, 2)
	b.Put(O, 0, 1)
	b.Put(O, 1, 1)
	assertFalse(b.Win() == O, "O should not win", t)

	b = Board{}
	b.Start()
	b.Put(O, 0, 2)
	b.Put(X, 2, 2)
	b.Put(O, 1, 1)
	b.Put(O, 2, 0)
	assertTrue(b.Win() == O, "III - O wins", t)

	b = Board{}
	b.Start()
	assertTrue(b.Win() == None, "beginning of the game", t)
}

func TestFork(t *testing.T) {

	test := func(p1, p2, expected Pos) {
		b := Board{}
		b.Start()
		b.Put(X, p1.row, p1.col)
		b.Put(X, p2.row, p2.col)
		p := b.Fork(X)
		if p == nil {
			t.Error("fork returned nil")
		} else {
			if p.row != expected.row || p.col != expected.col {
				t.Errorf("fork should return %d %d for fork but was %d %d", expected.row,
					expected.col, p.row, p.col)
			}
		}
	}

	test(Pos{X, 0, 0},
		Pos{X, 1, 0}, Pos{X, 2, 0})

	test(Pos{X, 0, 0}, Pos{X, 2, 0}, Pos{X, 1, 0})
	test(Pos{X, 0, 0}, Pos{X, 1, 1}, Pos{X, 2, 2})
	test(Pos{X, 0, 2}, Pos{X, 2, 0}, Pos{X, 1, 1})
	test(Pos{X, 1, 2}, Pos{X, 2, 2}, Pos{X, 0, 2})
}

func TestForkShouldNotReturnOccupiedPos(t *testing.T) {
	b := Board{}
	b.Start()
	b.Put(X, 0, 0)
	b.Put(X, 2, 2)
	b.Put(X, 1, 1)

	p := b.Fork(X)

	if p != nil {
		t.Errorf("p should be nil - occupied position - but was %s", p.String())
	}
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

func check(err error, t *testing.T) {
	if err != nil {
		t.Fatal(err)
	}
}
