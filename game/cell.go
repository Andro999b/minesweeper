package game

type Cell struct {
	x              int
	y              int
	isMine         bool
	neighboreMines int
	uncovered      bool
	marked         bool
}

func (c *Cell) IsMine() bool {
	return c.isMine
}

func (c *Cell) IsUncovered() bool {
	return c.uncovered
}

func (c *Cell) NeighborMines() int {
	return c.neighboreMines
}

func (c *Cell) IsMarked() bool {
	return c.marked
}
