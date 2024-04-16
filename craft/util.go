package craft

import (
	"image/color"

	"github.com/hajimehoshi/bitmapfont/v3"
	"github.com/hajimehoshi/ebiten/v2"
	ebitenText "github.com/hajimehoshi/ebiten/v2/text"
)

// calcSize is Size + Margin
func calcSize(s Size, m Margin) Size {
	width := s.X + m.Left + m.Right
	height := s.Y + m.Top + m.Bottom
	return Size{width, height}
}

func drawText(image *ebiten.Image, str string, color color.Color) {
	ebitenText.Draw(image, str, bitmapfont.Face, 0, 12, color)
}
