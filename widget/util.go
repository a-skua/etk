package widget

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/bitmapfont/v3"
	"github.com/hajimehoshi/ebiten/v2"
	ebitenText "github.com/hajimehoshi/ebiten/v2/text"
)

func calcSize(size image.Point, margin Margin, padding Padding) image.Point {
	width := size.X + margin.Left + margin.Right + padding.Left + padding.Right
	height := size.Y + margin.Top + margin.Bottom + padding.Top + padding.Bottom
	return image.Point{width, height}
}

func calcPosition(margin Margin, padding Padding) image.Point {
	x := margin.Left + padding.Left
	y := margin.Top + padding.Top
	return image.Point{x, y}
}

func drawText(image *ebiten.Image, str string, color color.Color) {
	ebitenText.Draw(image, str, bitmapfont.Face, 0, 12, color)
}
