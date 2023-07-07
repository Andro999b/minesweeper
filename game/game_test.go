package game

import "testing"

func TestMaxHeight(t *testing.T) {
	_, err := NewGame(1, MAX_HEIGHT+1, 10)
	if err.Error() != "height to big" {
		t.Errorf("Should fail if height bigger that %v", MAX_HEIGHT)
	}
}

func TestMaxWidth(t *testing.T) {
	_, err := NewGame(MAX_WIDTH+1, 1, 10)
	if err.Error() != "width to big" {
		t.Errorf("Should fail if widht bigger that %v", MAX_WIDTH)
	}
}

func TestMaxMines(t *testing.T) {
	_, err := NewGame(1, 1, 1)
	if err.Error() != "too much mines" {
		t.Error("Should fail if mines count bigger or equals field size")
	}
}

func TestMinMines(t *testing.T) {
	_, err := NewGame(1, 1, 0)
	if err.Error() != "should be atleast 1 mine on a field" {
		t.Error("Should fail if no mines on a field")
	}

	_, err = NewGame(1, 1, -1)
	if err.Error() != "should be atleast 1 mine on a field" {
		t.Error("Should fail if no mines is negative")
	}
}

func TestZeroOrNegativeHeight(t *testing.T) {
	_, err := NewGame(1, 0, 1)
	if err.Error() != "field to small" {
		t.Error("Should fail if field heigth 0")
	}

	_, err = NewGame(1, -1, 1)
	if err.Error() != "field to small" {
		t.Error("Should fail if field heigth negative")
	}
}

func TestZeroOrNegativeWidth(t *testing.T) {
	_, err := NewGame(0, 1, 1)
	if err.Error() != "field to small" {
		t.Error("Should fail if field width 0")
	}

	_, err = NewGame(0, 1, 1)
	if err.Error() != "field to small" {
		t.Error("Should fail if field width negative")
	}
}

func TestGetWidth(t *testing.T) {
	g := Game{w: 10}

	if g.GetWidth() != g.w {
		t.Error("GetWidth should return w field value")
	}
}

func TestGetHeight(t *testing.T) {
	g := Game{h: 10}

	if g.GetHeight() != g.h {
		t.Error("GetHeight should return h field value")
	}
}

func TestGetMines(t *testing.T) {
	g := Game{mines: 10}

	if g.GetMines() != g.mines {
		t.Error("GetMines should return mines field value")
	}
}

func TestIsLose(t *testing.T) {
	g := Game{}

	if g.IsLose() {
		t.Error("Game should not be lost untile mine hit")
	}

	g.minehit = true

	if !g.IsLose() {
		t.Error("Game should be lost if mine hit")
	}

	if !g.IsOver() {
		t.Error("Game should be over if player  lose")
	}
}
