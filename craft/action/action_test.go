package action

import (
	"image/color"

	"github.com/a-skua/etk/craft"
	"github.com/a-skua/etk/craft/types"
	"github.com/hajimehoshi/ebiten/v2"
)

var _ craft.Craft = NewMousePressed(
	craft.NewFill(types.Size{X: 10, Y: 10}, color.White),
	ebiten.MouseButtonLeft,
	func(f *craft.Fill) error {
		return nil
	},
)
