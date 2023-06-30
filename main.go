package main

import (
	"flag"
	"log"
	"minesweeper/game"
	"minesweeper/view"
)

func main() {
	var err error
	var w, h, m int

	flag.IntVar(&w, "w", 40, "minefield width")
	flag.IntVar(&h, "h", 10, "minefield height")
	flag.IntVar(&m, "m", 20, "number of mines on a field")
	flag.Parse()

	g, err := game.NewGame(w, h, m)

	if err != nil {
		log.Fatalf("Fail to create a game: %v", err)
	}

	v, err := view.NewView(g)
	if err != nil {
		log.Fatalf("Fail to create a game view: %v", err)
	}

	v.RunGame()
}
