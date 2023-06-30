package view

import (
	"fmt"
	"minesweeper/game"
	"strconv"

	"github.com/gdamore/tcell"
)

type Cursor struct {
	x int
	y int
}

type View struct {
	cursor *Cursor
	game   *game.Game
	screen tcell.Screen
}

func NewView(g *game.Game) (*View, error) {
	var err error

	screen, err := tcell.NewScreen()
	if err != nil {
		return nil, err
	}
	err = screen.Init()
	if err != nil {
		return nil, err
	}

	screen.Clear()

	view := &View{&Cursor{0, 0}, g, screen}
	return view, nil
}

func (v *View) RunGame() {
gameLoop:
	for {
		event := v.screen.PollEvent()
		switch eventType := event.(type) {
		case *tcell.EventKey:
			if !v.processInput(eventType) {
				break gameLoop
			}
			v.RefreshScreen()
		case *tcell.EventResize:
			v.RefreshScreen()
		}
	}

	v.screen.Fini()
}

func (v *View) processInput(event *tcell.EventKey) bool {
	switch event.Key() {
	case tcell.KeyCtrlC:
		return false
	case tcell.KeyRune:
		switch event.Rune() {
		case 'q':
			return false
		case ' ':
			v.game.OpenCell(v.cursor.x, v.cursor.y)
		case 'x':
			v.game.ToggleMark(v.cursor.x, v.cursor.y)
		case 'r':
			v.screen.Clear()
			v.game.Reset()
		}
	case tcell.KeyLeft:
		v.moveCursor(-1, 0)
	case tcell.KeyRight:
		v.moveCursor(1, 0)
	case tcell.KeyUp:
		v.moveCursor(0, -1)
	case tcell.KeyDown:
		v.moveCursor(0, 1)
	}

	return true
}

func (v *View) RefreshScreen() {
	v.drawText(4, 0, "MineSweeper")
	v.drawText(4, 1, fmt.Sprintf("Mines on a field: %v", v.game.GetMines()))
	v.drawText(4, 2, "Use ←↑↓→ to move cursor, Space to open cell, X to mark cell. Press Q or CTRL+C to quite")
	if v.game.IsOver() {
		outcome := "WIN"
		if v.game.IsLose() {
			outcome = "LOSE"
		}
		v.drawText(4, 3, fmt.Sprintf("Game over! You %s! Press R to reset", outcome))
	}
	v.drawBoard(4, 4)
	v.screen.Show()
}

func (v *View) moveCursor(xOffset int, yOffset int) {
	v.cursor.x += xOffset
	v.cursor.y += yOffset

	if v.cursor.x < 0 {
		v.cursor.x = 0
	} else if v.cursor.x >= v.game.GetWidth() {
		v.cursor.x = v.game.GetWidth() - 1
	}

	if v.cursor.y < 0 {
		v.cursor.y = 0
	} else if v.cursor.y >= v.game.GetHeight() {
		v.cursor.y = v.game.GetHeight() - 1
	}
}

func (v *View) drawBoard(xOffset int, yOffset int) {
	gameOver := v.game.IsOver()
	for x := 0; x < v.game.GetWidth(); x++ {
		for y := 0; y < v.game.GetHeight(); y++ {
			var char rune
			var style tcell.Style
			cell := v.game.GetCell(x, y)

			if gameOver && cell.IsMine() {
				char = '*'
				style = tcell.StyleDefault.Foreground(tcell.ColorRed)
			} else if cell.IsUncovered() {
				style = tcell.StyleDefault.Background(tcell.ColorGreen)
				if cell.NeighborMines() > 0 {
					s := strconv.Itoa(cell.NeighborMines())
					char = []rune(s)[0]
				} else {
					char = ' '
				}
			} else {
				if cell.IsMarked() {
					char = 'x'
				} else {
					char = ' '
				}
				style = tcell.StyleDefault.Background(tcell.ColorGray)
			}

			if v.cursor.x == x && v.cursor.y == y {
				style = tcell.StyleDefault.Background(tcell.ColorBeige).Foreground(tcell.ColorBlack)
			}

			v.screen.SetContent(x+xOffset, y+yOffset, char, nil, style)
		}
	}
}

func (v *View) drawText(x int, y int, text string) {
	for index, char := range text {
		v.screen.SetContent(x+index, y, rune(char), nil, tcell.StyleDefault)
	}
}
