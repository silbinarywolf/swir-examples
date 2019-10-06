package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

var (
	emptyImage *ebiten.Image
)

func init() {
	emptyImage, _ = ebiten.NewImage(1, 1, ebiten.FilterDefault)
	_ = emptyImage.Fill(color.White)
}

func colorScale(clr color.Color) (rf, gf, bf, af float64) {
	r, g, b, a := clr.RGBA()
	if a == 0 {
		return 0, 0, 0, 0
	}

	rf = float64(r) / float64(a)
	gf = float64(g) / float64(a)
	bf = float64(b) / float64(a)
	af = float64(a) / 0xffff
	return
}

// drawRect draws a rectangle on the given destination dst.
//
// drawRect is intended to be used mainly for debugging or prototyping purpose.
func drawRect(dst *ebiten.Image, x, y, width, height float64, clr color.Color) {
	ew, eh := emptyImage.Size()

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(width/float64(ew), height/float64(eh))
	op.GeoM.Translate(x, y)
	op.ColorM.Scale(colorScale(clr))
	// Filter must be 'nearest' filter (default).
	// Linear filtering would make edges blurred.
	_ = dst.DrawImage(emptyImage, op)
}
