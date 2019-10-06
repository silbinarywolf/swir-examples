package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/silbinarywolf/swir-examples/squaregame/internal/game/input"
)

func Init() {
	input.ButtonInit()
}

func Update(screen *ebiten.Image) error {
	input.ButtonUpdate()
	player.Update()

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	// Fill the screen with #FF0000 color
	screen.Fill(color.NRGBA{0x00, 0x00, 0x00, 0x00})
	drawRect(screen, player.X, player.Y, 32, 32, color.NRGBA{0xff, 0x00, 0x00, 0xff})
	return nil
}
