package action

import (
	"image/color"
	"log/slog"

	"github.com/a-skua/etk/craft"
	"github.com/a-skua/etk/craft/internal/util"
	"github.com/a-skua/etk/craft/types"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type ActionHandler[T craft.Craft] func(T) error

type MousePressed[T craft.Craft] struct {
	craft   T
	button  ebiten.MouseButton
	handler ActionHandler[T]
}

func NewMousePressed[T craft.Craft](craft T, button ebiten.MouseButton, handler ActionHandler[T]) *MousePressed[T] {
	return &MousePressed[T]{craft, button, handler}
}

func (m *MousePressed[T]) Image() *ebiten.Image {
	return m.craft.Image()
}

func (m *MousePressed[T]) Size() types.Size {
	return m.craft.Size()
}

func (m *MousePressed[T]) AddText(str string, color color.Color) craft.Self {
	return m.craft.AddText(str, color)
}

func (m *MousePressed[T]) Const() *craft.Image {
	return m.craft.Const()
}

func (m *MousePressed[T]) Update(p types.Position) error {
	if !inpututil.IsMouseButtonJustPressed(m.button) {
		return nil
	}

	x, y := ebiten.CursorPosition()
	p = types.Position{X: x, Y: y}.Sub(p)
	if !util.Sizein(m.craft.Size(), p) {
		return nil
	}

	slog.Debug("MousePressed.Update")
	return m.handler(m.craft)
}
