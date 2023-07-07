package game

import "testing"

func TestIsMine(t *testing.T) {
	cell := &Cell{isMine: true}

	if cell.IsMine() != cell.isMine {
		t.Fatal("Cell.IsMine() should returns isMine field value")
	}
}

func TestIsUncovered(t *testing.T) {
	cell := &Cell{uncovered: true}

	if cell.IsUncovered() != cell.uncovered {
		t.Fatal("Cell.IsUncovered() should returns uncovered field value")
	}
}

func TestIsMarked(t *testing.T) {
	cell := &Cell{marked: true}

	if cell.IsMarked() != cell.marked {
		t.Fatal("Cell.IsMarked() should returns marked field value")
	}
}

func TestNeighborMines(t *testing.T) {
	cell := &Cell{neighborMines: 3}

	if cell.NeighborMines() != cell.neighborMines {
		t.Fatal("Cell.NeighborMines() should returns neighborMines field value")
	}
}
