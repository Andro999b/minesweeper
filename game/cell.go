package game

type Cell struct {
	x             int
	y             int
	isMine        bool
	neighborMines int
	uncovered     bool
	marked        bool
}

func (c *Cell) IsMine() bool {
	return c.isMine
}

func (c *Cell) IsUncovered() bool {
	return c.uncovered
}

func (c *Cell) NeighborMines() int {
	return c.neighborMines
}

func (c *Cell) IsMarked() bool {
	return c.marked
}
