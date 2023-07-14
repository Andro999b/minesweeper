package main

import (
	"flag"
	"log"
	"minesweeper/game"
	"minesweeper/view"
)

func main() {
	var err error
	var w, h, m, l int

	flag.IntVar(&w, "w", 40, "minefield width")
	flag.IntVar(&h, "h", 10, "minefield height")
	flag.IntVar(&m, "m", 40, "number of mines on a field")
	flag.IntVar(&l, "l", 1, "number of lives")
	flag.Parse()

	g, err := game.NewGame(w, h, m, l)

	if err != nil {
		log.Fatalf("Fail to create a game: %v", err)
	}

	v, err := view.NewView(g)
	if err != nil {
		log.Fatalf("Fail to create a game view: %v", err)
	}

	v.RunGame()
}
