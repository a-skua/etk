package util

import (
	"image/color"

	"github.com/a-skua/etk/craft/types"
	"github.com/hajimehoshi/bitmapfont/v3"
	"github.com/hajimehoshi/ebiten/v2"
	ebitenText "github.com/hajimehoshi/ebiten/v2/text"
)

// CalcSize is types.Size + types.Margin
func CalcSize(s types.Size, m types.Margin) types.Size {
	return types.Size{
		X: s.X + m.Left + m.Right,
		Y: s.Y + m.Top + m.Bottom,
	}
}

func DrawText(image *ebiten.Image, str string, color color.Color) {
	ebitenText.Draw(image, str, bitmapfont.Face, 0, 12, color)
}

func Sizein(s types.Size, p types.Position) bool {
	return 0 <= p.X && p.X < s.X &&
		0 <= p.Y && p.Y < s.Y
}
