package game

import (
	"reflect"
	"testing"
)

func TestMaxHeight(t *testing.T) {
	_, err := NewGame(1, MAX_HEIGHT+1, 10, 1)
	if err == nil || err.Error() != "height to big" {
		t.Errorf("Should fail if height bigger that %v", MAX_HEIGHT)
	}
}

func TestMaxWidth(t *testing.T) {
	_, err := NewGame(MAX_WIDTH+1, 1, 10, 1)
	if err == nil || err.Error() != "width to big" {
		t.Errorf("Should fail if widht bigger that %v", MAX_WIDTH)
	}
}

func TestMaxMines(t *testing.T) {
	_, err := NewGame(1, 1, 1, 1)
	if err == nil || err.Error() != "too much mines" {
		t.Error("Should fail if mines count bigger or equals field size")
	}
}

func TestMinMines(t *testing.T) {
	_, err := NewGame(1, 1, 0, 1)
	if err == nil || err.Error() != "should be atleast 1 mine on a field" {
		t.Error("Should fail if no mines on a field")
	}

	_, err = NewGame(1, 1, -1, 1)
	if err == nil || err.Error() != "should be atleast 1 mine on a field" {
		t.Error("Should fail if no mines is negative")
	}
}

func TestZeroOrNegativeHeight(t *testing.T) {
	_, err := NewGame(1, 0, 1, 1)
	if err == nil || err.Error() != "field to small" {
		t.Error("Should fail if field heigth 0")
	}

	_, err = NewGame(1, -1, 1, 1)
	if err == nil || err.Error() != "field to small" {
		t.Error("Should fail if field heigth negative")
	}
}

func TestZeroOrNegativeWidth(t *testing.T) {
	_, err := NewGame(0, 1, 1, 1)
	if err == nil || err.Error() != "field to small" {
		t.Error("Should fail if field width 0")
	}

	_, err = NewGame(0, 1, 1, 1)
	if err == nil || err.Error() != "field to small" {
		t.Error("Should fail if field width negative")
	}
}

func TestZeroOrNegativeLives(t *testing.T) {
	_, err := NewGame(2, 2, 1, 0)

	if err == nil || err.Error() != "player should have atleast 1 live" {
		t.Error("Should fail if field width 0")
	}

	_, err = NewGame(2, 2, 1, -1)
	if err == nil || err.Error() != "player should have atleast 1 live" {
		t.Error("Should fail if field width negative")
	}
}

func TestGetCell(t *testing.T) {
	game, _ := NewGame(2, 2, 1, 1)

	if game.GetCell(-1, 0) != nil {
		t.Error("GetCell should return nil for negatives")
	}

	if game.GetCell(0, -1) != nil {
		t.Error("GetCell should return nil for negatives")
	}

	if game.GetCell(0, 0) == nil {
		t.Error("GetCell should return cell for valida position")
	}

	if game.GetCell(0, 999) != nil {
		t.Error("GetCell should nil for out of bondries position")
	}

	if game.GetCell(999, 0) != nil {
		t.Error("GetCell should nil for out of bondries position")
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
	g := Game{livesLeft: 1}

	if g.IsLose() {
		t.Error("Game should not be lost untile mine hit")
	}

	g.livesLeft = 0

	if !g.IsLose() {
		t.Error("Game should be lost if mine hit")
	}

	if !g.IsOver() {
		t.Error("Game should be over if player lose")
	}
}

func TestGenerateField(t *testing.T) {
	actualGetRandSeed := getRandSeed
	getRandSeed = func() int64 { return 1 }
	defer func() {
		getRandSeed = actualGetRandSeed
	}()

	game, err := NewGame(3, 3, 2, 1)

	if err != nil {
		t.Error("Should creates new game without error")
	}

	actualXPos := flatField[int](game, func(c *Cell) int { return c.x })
	expectedXPos := []int{
		0, 1, 2,
		0, 1, 2,
		0, 1, 2,
	}
	if !reflect.DeepEqual(actualXPos, expectedXPos) {
		t.Error("Incorrect cells x pos")
	}

	actualYPos := flatField[int](game, func(c *Cell) int { return c.y })
	expectedYPos := []int{
		0, 0, 0,
		1, 1, 1,
		2, 2, 2,
	}
	if !reflect.DeepEqual(actualYPos, expectedYPos) {
		t.Error("Incorrect cells y pos")
	}

	actualMinesPos := flatField[bool](game, func(c *Cell) bool { return c.isMine })
	expectedMinesPos := []bool{
		false, false, false,
		false, false, false,
		false, true, true,
	}
	if !reflect.DeepEqual(actualMinesPos, expectedMinesPos) {
		t.Error("Incorrect mines placement")
	}

	actualNeighborMines := flatField[int](game, func(c *Cell) int { return c.neighborMines })
	expectedNeighborMines := []int{
		0, 0, 0,
		1, 2, 2,
		1, 1, 1,
	}
	if !reflect.DeepEqual(actualNeighborMines, expectedNeighborMines) {
		t.Error("Incorrect mines neigbor count")
	}
}

func TestToggleMark(t *testing.T) {
	actualGetRandSeed := getRandSeed
	getRandSeed = func() int64 { return 1 }
	defer func() {
		getRandSeed = actualGetRandSeed
	}()

	game, err := NewGame(3, 3, 2, 1)

	if err != nil {
		t.Error("Should creates new game without error")
	}

	game.ToggleMark(0, 0)
	actualMarked := flatField[bool](game, func(c *Cell) bool { return c.marked })
	expectMarked := []bool{
		true, false, false,
		false, false, false,
		false, false, false,
	}

	if !reflect.DeepEqual(actualMarked, expectMarked) {
		t.Error("Incorrect marked cells")
	}
}

func TestOpenCell(t *testing.T) {
	actualGetRandSeed := getRandSeed
	getRandSeed = func() int64 { return 1 }
	defer func() {
		getRandSeed = actualGetRandSeed
	}()

	game, err := NewGame(3, 3, 2, 1)

	if err != nil {
		t.Error("Should creates new game without error")
	}

	// should open single cell
	game.OpenCell(0, 2)

	actualUncovered := flatField[bool](game, func(c *Cell) bool { return c.uncovered })
	expectUncovered := []bool{
		false, false, false,
		false, false, false,
		true, false, false,
	}

	if !reflect.DeepEqual(actualUncovered, expectUncovered) {
		t.Error("Incorrect uncovered cells")
	}

	game.OpenCell(0, 0)

	actualUncovered = flatField[bool](game, func(c *Cell) bool { return c.uncovered })
	expectUncovered = []bool{
		true, true, true,
		true, true, true,
		true, false, false,
	}

	if !reflect.DeepEqual(actualUncovered, expectUncovered) {
		t.Error("Incorrect uncovered cells")
	}

	if !game.IsOver() {
		t.Error("Game should be over")
	}

	if game.IsLose() {
		t.Error("Player should win")
	}
}

func TestGameLose(t *testing.T) {
	actualGetRandSeed := getRandSeed
	getRandSeed = func() int64 { return 1 }
	defer func() {
		getRandSeed = actualGetRandSeed
	}()

	game, err := NewGame(3, 3, 2, 1)
	if err != nil {
		t.Error("Should creates new game without error")
	}

	game.OpenCell(2, 2)

	if !game.IsOver() {
		t.Error("Game should be over")
	}

	if !game.IsLose() {
		t.Error("Player should lose")
	}

	if game.livesLeft != 0 {
		t.Error("Game livesLeft should be 0")
	}
}

func TestLivesDec(t *testing.T) {
	actualGetRandSeed := getRandSeed
	getRandSeed = func() int64 { return 1 }
	defer func() {
		getRandSeed = actualGetRandSeed
	}()

	game, err := NewGame(3, 3, 2, 2)
	if err != nil {
		t.Error("Should creates new game without error")
	}

	if game.GetLivesLeft() != 2 {
		t.Error("GetLivesLeft should retrun correct value")
	}

	game.OpenCell(2, 2)

	if game.IsOver() {
		t.Error("Game should not be over")
	}

	if game.IsLose() {
		t.Error("Player should not lose")
	}

	if game.GetLivesLeft() != 1 {
		t.Error("GetLivesLeft should retrun correct value")
	}
}

func flatField[T any](game *Game, fn func(*Cell) T) []T {
	slice := make([]T, game.w*game.h)

	for x := range game.field {
		for y := range game.field[x] {
			i := x*game.h + y
			slice[i] = fn(game.field[y][x])
		}
	}

	return slice
}
