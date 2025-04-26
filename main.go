package main

import (
	"dhaka_drift/logic"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

func main() {
	game := logic.NewGame()

	ebiten.SetWindowSize(480, 800)
	ebiten.SetWindowTitle("Dhaka Drift")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
