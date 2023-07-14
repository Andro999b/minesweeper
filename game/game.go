package game

import (
	"container/list"
	"errors"
	"math/rand"
	"time"
)

var getRandSeed = func() int64 {
	return time.Now().UnixNano()
}

const MAX_HEIGHT = 20
const MAX_WIDTH = 100

type neighborVisitor func(*Cell)

type Game struct {
	field     [][]*Cell
	h         int
	w         int
	lives     int
	livesLeft int
	mines     int
	uncovered int
}

func (g *Game) GetWidth() int {
	return g.w
}

func (g *Game) GetHeight() int {
	return g.h
}

func (g *Game) GetMines() int {
	return g.mines
}

func (g *Game) GetLivesLeft() int {
	return g.livesLeft
}

func (g *Game) GetCell(x int, y int) *Cell {
	if x < 0 || y < 0 || x >= g.w || y >= g.h {
		return nil
	}

	return g.field[x][y]
}

func (g *Game) ToggleMark(x int, y int) {
	if g.IsOver() {
		return
	}

	cell := g.GetCell(x, y)
	if cell == nil || cell.uncovered {
		return
	}

	cell.marked = !cell.marked
}

func (g *Game) OpenCell(x int, y int) {
	if g.IsOver() {
		return
	}

	cell := g.GetCell(x, y)
	if cell == nil || cell.uncovered || cell.marked {
		return
	}

	if cell.isMine {
		cell.marked = false
		cell.uncovered = true
		g.livesLeft--
		return
	}

	g.uncoverCell(cell)
	if cell.neighborMines > 0 {
		return
	}

	q := list.New()
	q.PushBack(cell)
	for q.Len() > 0 {
		front := q.Front()
		cur := front.Value.(*Cell)
		q.Remove(front)

		g.forEachCellNeighbor(cur, func(neighbor *Cell) {
			if neighbor.uncovered {
				return
			}

			g.uncoverCell(neighbor)
			if neighbor.neighborMines > 0 {
				return
			}

			q.PushBack(neighbor)
		})
	}
}

func (g *Game) IsLose() bool {
	return g.livesLeft == 0
}

func (g *Game) IsOver() bool {
	return g.IsLose() || g.mines+g.uncovered == g.w*g.h
}

func (g *Game) Reset() {
	g.livesLeft = g.lives
	g.uncovered = 0
	g.generateField()
}

func (g *Game) generateField() {
	field := make([][]*Cell, g.w)
	for i := range field {
		field[i] = make([]*Cell, g.h)
		for j := range field[i] {
			field[i][j] = &Cell{x: i, y: j}
		}
	}

	g.field = field

	cellsCount := g.w * g.h
	poss := make([]struct {
		x int
		y int
	}, cellsCount)

	// generate array of all positions
	for x := 0; x < g.w; x++ {
		for y := 0; y < g.h; y++ {
			i := x*g.h + y
			poss[i].x = x
			poss[i].y = y
		}
	}

	rnd := rand.New(rand.NewSource(getRandSeed()))

	// for each mine
	for m := 0; m < g.mines; m++ {
		// pick random position form positions array
		posIndex := rnd.Intn(len(poss))
		pos := poss[posIndex]

		// mark cell as a mine
		cell := g.field[pos.x][pos.y]
		cell.isMine = true

		// inc count of mines for each neighbor cells
		g.forEachCellNeighbor(cell, func(neighbor *Cell) {
			neighbor.neighborMines++
		})

		// remove picked position
		poss = append(poss[:posIndex], poss[posIndex+1:]...)
	}
}

func (g *Game) forEachCellNeighbor(cell *Cell, cb neighborVisitor) {
	for xo := -1; xo <= 1; xo++ {
		for yo := -1; yo <= 1; yo++ {
			// skip current cell
			if xo == 0 && yo == 0 {
				continue
			}

			nx := cell.x - xo
			ny := cell.y - yo

			// check if position in range
			if nx < 0 || ny < 0 || nx >= g.w || ny >= g.h {
				continue
			}

			cb(g.field[nx][ny])
		}
	}
}

func (g *Game) uncoverCell(cell *Cell) {
	cell.marked = false
	cell.uncovered = true
	g.uncovered++
}

func NewGame(w int, h int, mines int, lives int) (*Game, error) {
	if w < 1 || h < 1 {
		return nil, errors.New("field to small")
	}

	if mines < 1 {
		return nil, errors.New("should be atleast 1 mine on a field")
	}

	if h > MAX_HEIGHT {
		return nil, errors.New("height to big")
	}

	if w > MAX_WIDTH {
		return nil, errors.New("width to big")
	}

	if w*h <= mines {
		return nil, errors.New("too much mines")
	}

	if lives < 0 {
		return nil, errors.New("lives should be bigger that 0")
	}

	game := &Game{
		w:         w,
		h:         h,
		mines:     mines,
		lives:     lives,
		livesLeft: lives,
	}

	game.generateField()

	return game, nil
}
