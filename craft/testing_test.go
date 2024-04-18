package craft

import (
	"image/color"

	"github.com/a-skua/etk/craft/types"
	"github.com/hajimehoshi/ebiten/v2"
)

// Craft
type testingCraft struct {
	size          types.Size
	updateHandler func(types.Position) error
}

func (t *testingCraft) Image() *ebiten.Image {
	return &ebiten.Image{}
}

func (t *testingCraft) Size() types.Size {
	return t.size
}

func (t *testingCraft) AddText(s string, c color.Color) Self {
	return t
}

func (t *testingCraft) Const() *Image {
	return &Image{}
}

func (t *testingCraft) Update(p types.Position) error {
	return t.updateHandler(p)
}
