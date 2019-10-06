package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/silbinarywolf/swir-examples/squaregame/internal/game"
)

func main() {
	game.Init()
	if err := ebiten.Run(game.Update, 320, 240, 2, "Hello world!"); err != nil {
		panic(err)
	}
}
