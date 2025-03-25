package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"follow-the-leader/cmd/game"
)

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Follow the Leader")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	g, err := game.New()
	if err != nil {
		log.Fatal(err)
	}
	
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}