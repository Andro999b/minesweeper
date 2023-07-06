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
		t.Errorf("Should fail if mines count bigger or equals field size")
	}
}

func TestMinMines(t *testing.T) {
	_, err := NewGame(1, 1, 0)
	if err.Error() != "should be atleast 1 mine on a field" {
		t.Errorf("Should fail if no mines on a field")
	}

	_, err = NewGame(1, 1, -1)
	if err.Error() != "should be atleast 1 mine on a field" {
		t.Errorf("Should fail if no mines is negative")
	}
}

func TestZeroOrNegativeHeight(t *testing.T) {
	_, err := NewGame(1, 0, 1)
	if err.Error() != "field to small" {
		t.Errorf("Should fail if field heigth 0")
	}

	_, err = NewGame(1, -1, 1)
	if err.Error() != "field to small" {
		t.Errorf("Should fail if field heigth negative")
	}
}

func TestZeroOrNegativeWidth(t *testing.T) {
	_, err := NewGame(0, 1, 1)
	if err.Error() != "field to small" {
		t.Errorf("Should fail if field width 0")
	}

	_, err = NewGame(0, 1, 1)
	if err.Error() != "field to small" {
		t.Errorf("Should fail if field width negative")
	}
}
