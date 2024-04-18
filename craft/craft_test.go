package craft

import (
	"fmt"
	"image"
	"image/color"
	"testing"

	"github.com/a-skua/etk/craft/types"
)

var _ Craft = NewImage(image.NewRGBA(image.Rect(0, 0, 100, 100)))
var _ Craft = NewFill(types.Size{X: 10, Y: 10}, color.White)
var _ Craft = NewSwitch(
	NewFill(types.Size{X: 10, Y: 10}, color.White),
	NewFill(types.Size{X: 10, Y: 10}, color.White),
	NewFill(types.Size{X: 10, Y: 10}, color.White),
)
var _ Craft = NewBox(NewFill(types.Size{X: 10, Y: 10}, color.White), types.Margin{})
var _ Craft = NewHorizontalStack(
	NewBox(NewFill(types.Size{X: 10, Y: 10}, color.White), types.Margin{}),
	NewBox(NewFill(types.Size{X: 10, Y: 10}, color.White), types.Margin{}),
)
var _ Craft = NewVerticalStack(
	NewBox(NewFill(types.Size{X: 10, Y: 10}, color.White), types.Margin{}),
	NewBox(NewFill(types.Size{X: 10, Y: 10}, color.White), types.Margin{}),
)
var _ Craft = NewLayer(
	NewFill(types.Size{X: 10, Y: 20}, color.White),
	NewFill(types.Size{X: 20, Y: 10}, color.White),
)

func TestSwitchNext(t *testing.T) {
	switcher := NewSwitch(
		NewFill(types.Size{X: 10, Y: 10}, color.White),
		NewFill(types.Size{X: 10, Y: 10}, color.White),
		NewFill(types.Size{X: 10, Y: 10}, color.White),
	)

	switcher.Next()
	if switcher.index != 1 {
		t.Errorf("Next should increment index, but got %d", switcher.index)
	}

	switcher.Next()
	if switcher.index != 2 {
		t.Errorf("Next should increment index, but got %d", switcher.index)
	}

	switcher.Next()
	if switcher.index != 0 {
		t.Errorf("Next should reset index, but got %d", switcher.index)
	}
}

func TestSwitchPrev(t *testing.T) {
	switcher := NewSwitch(
		NewFill(types.Size{X: 10, Y: 10}, color.White),
		NewFill(types.Size{X: 10, Y: 10}, color.White),
		NewFill(types.Size{X: 10, Y: 10}, color.White),
	)

	switcher.Prev()
	if switcher.index != 2 {
		t.Errorf("Prev should decrement index, but got %d", switcher.index)
	}

	switcher.Prev()
	if switcher.index != 1 {
		t.Errorf("Prev should decrement index, but got %d", switcher.index)
	}

	switcher.Prev()
	if switcher.index != 0 {
		t.Errorf("Prev should reset index, but got %d", switcher.index)
	}
}

func TestBoxSize(t *testing.T) {
	tests := []struct {
		craft Craft
		want  types.Size
	}{
		{
			NewBox(NewFill(types.Size{X: 10, Y: 10}, color.White), types.Margin{}),
			types.Size{X: 10, Y: 10},
		},
		{
			NewBox(NewFill(types.Size{X: 10, Y: 10}, color.White), types.Margin{Left: 10, Top: 15, Right: 10, Bottom: 15}),
			types.Size{X: 30, Y: 40},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprint(i+1), func(t *testing.T) {
			size := tt.craft.Size()
			if size != tt.want {
				t.Errorf("types.Size should return %v, but got %v", tt.want, size)
			}
		})
	}
}

func TestBoxUpdate(t *testing.T) {
	box := NewBox(
		&testingCraft{
			updateHandler: func(p types.Position) error {
				want := types.Position{X: 10, Y: 10}
				if p != want {
					t.Errorf("Update should receive %v, but got %v", want, p)
				}
				return nil
			},
		},
		types.MarginAll(10),
	)
	box.Update(types.Position{})
}

func TestHorizontalStackSize(t *testing.T) {
	tests := []struct {
		widgets []Craft
		want    types.Size
	}{
		{
			[]Craft{
				NewBox(NewFill(types.Size{X: 10, Y: 10}, color.White), types.Margin{}),
				NewBox(NewFill(types.Size{X: 10, Y: 10}, color.White), types.Margin{}),
			},
			types.Size{X: 20, Y: 10},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprint(i+1), func(t *testing.T) {
			s := NewHorizontalStack(tt.widgets...)
			size := s.Size()
			if size != tt.want {
				t.Errorf("types.Size should return %v, but got %v", tt.want, size)
			}
		})
	}
}

func TestHorizontalStackUpdate(t *testing.T) {
	stack := NewHorizontalStack(
		&testingCraft{
			updateHandler: func(p types.Position) error {
				want := types.Position{X: 0, Y: 0}
				if p != want {
					t.Errorf("Update should receive %v, but got %v", want, p)
				}
				return nil
			},
			size: types.Size{X: 10, Y: 10},
		},
		&testingCraft{
			updateHandler: func(p types.Position) error {
				want := types.Position{X: 10, Y: 0}
				if p != want {
					t.Errorf("Update should receive %v, but got %v", want, p)
				}
				return nil
			},
			size: types.Size{X: 10, Y: 10},
		},
		&testingCraft{
			updateHandler: func(p types.Position) error {
				want := types.Position{X: 20, Y: 0}
				if p != want {
					t.Errorf("Update should receive %v, but got %v", want, p)
				}
				return nil
			},
			size: types.Size{X: 10, Y: 10},
		},
	)
	stack.Update(types.Position{})
}

func TestVerticalStackSize(t *testing.T) {
	tests := []struct {
		widgets []Craft
		want    types.Size
	}{
		{
			[]Craft{
				NewBox(NewFill(types.Size{X: 10, Y: 10}, color.White), types.Margin{}),
				NewBox(NewFill(types.Size{X: 10, Y: 10}, color.White), types.Margin{}),
			},
			types.Size{X: 10, Y: 20},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprint(i+1), func(t *testing.T) {
			s := NewVerticalStack(tt.widgets...)
			size := s.Size()
			if size != tt.want {
				t.Errorf("types.Size should return %v, but got %v", tt.want, size)
			}
		})
	}
}

func TestVerticalStackUpdate(t *testing.T) {
	stack := NewVerticalStack(
		&testingCraft{
			updateHandler: func(p types.Position) error {
				want := types.Position{X: 0, Y: 0}
				if p != want {
					t.Errorf("Update should receive %v, but got %v", want, p)
				}
				return nil
			},
			size: types.Size{X: 10, Y: 10},
		},
		&testingCraft{
			updateHandler: func(p types.Position) error {
				want := types.Position{X: 0, Y: 10}
				if p != want {
					t.Errorf("Update should receive %v, but got %v", want, p)
				}
				return nil
			},
			size: types.Size{X: 10, Y: 10},
		},
		&testingCraft{
			updateHandler: func(p types.Position) error {
				want := types.Position{X: 0, Y: 20}
				if p != want {
					t.Errorf("Update should receive %v, but got %v", want, p)
				}
				return nil
			},
			size: types.Size{X: 10, Y: 10},
		},
	)
	stack.Update(types.Position{})
}

func TestLayerSize(t *testing.T) {
	tests := []struct {
		widgets []Craft
		want    types.Size
	}{
		{
			[]Craft{
				NewFill(types.Size{X: 10, Y: 30}, color.White),
				NewFill(types.Size{X: 20, Y: 20}, color.White),
				NewFill(types.Size{X: 30, Y: 10}, color.White),
			},
			types.Size{X: 30, Y: 30},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprint(i+1), func(t *testing.T) {
			l := NewLayer(tt.widgets...)
			size := l.Size()
			if size != tt.want {
				t.Errorf("types.Size should return %v, but got %v", tt.want, size)
			}
		})
	}
}
