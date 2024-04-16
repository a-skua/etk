package widget

import (
	"fmt"
	"image"
	"image/color"
	"testing"
)

var _ Widget = NewImage(image.NewRGBA(image.Rect(0, 0, 100, 100)))
var _ Widget = NewFill(image.Point{10, 10}, color.White)
var _ Widget = NewBox(NewFill(image.Point{10, 10}, color.White), Margin{}, Padding{})
var _ Widget = NewStack(
	Vertical,
	NewBox(NewFill(image.Point{10, 10}, color.White), Margin{}, Padding{}),
	NewBox(NewFill(image.Point{10, 10}, color.White), Margin{}, Padding{}),
)
var _ Widget = NewHorizontalStack(
	NewBox(NewFill(image.Point{10, 10}, color.White), Margin{}, Padding{}),
	NewBox(NewFill(image.Point{10, 10}, color.White), Margin{}, Padding{}),
)
var _ Widget = NewVerticalStack(
	NewBox(NewFill(image.Point{10, 10}, color.White), Margin{}, Padding{}),
	NewBox(NewFill(image.Point{10, 10}, color.White), Margin{}, Padding{}),
)
var _ Widget = NewLayer(
	NewFill(image.Point{10, 20}, color.White),
	NewFill(image.Point{20, 10}, color.White),
)

var _ *Image = NewImage(image.NewRGBA(image.Rect(0, 0, 100, 100))).Const().(*Image)
var _ *Image = NewFill(image.Point{10, 10}, color.White).Const().(*Image)
var _ *Image = NewBox(NewFill(image.Point{10, 10}, color.White), Margin{}, Padding{}).Const().(*Image)
var _ *Image = NewHorizontalStack(
	NewBox(NewFill(image.Point{10, 10}, color.White), Margin{}, Padding{}),
	NewBox(NewFill(image.Point{10, 10}, color.White), Margin{}, Padding{}),
).Const().(*Image)
var _ *Image = NewVerticalStack(
	NewBox(NewFill(image.Point{10, 10}, color.White), Margin{}, Padding{}),
	NewBox(NewFill(image.Point{10, 10}, color.White), Margin{}, Padding{}),
).Const().(*Image)
var _ *Image = NewLayer(
	NewFill(image.Point{10, 20}, color.White),
	NewFill(image.Point{20, 10}, color.White),
).Const().(*Image)

func TestBoxPosition(t *testing.T) {
	tests := []struct {
		margin  Margin
		padding Padding
		want    image.Point
	}{
		{Margin{}, Padding{}, image.Point{0, 0}},
		{Margin{10, 0, 0, 0}, Padding{10, 0, 0, 0}, image.Point{20, 0}},
		{Margin{0, 10, 0, 0}, Padding{0, 10, 0, 0}, image.Point{0, 20}},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprint(i+1), func(t *testing.T) {
			b := NewBox(NewFill(image.Point{10, 10}, color.White), tt.margin, tt.padding)
			p := b.Position()
			if p != tt.want {
				t.Errorf("Position should return %v, but got %v", tt.want, p)
			}
		})
	}
}

func TestHorizontalStackSize(t *testing.T) {
	tests := []struct {
		widgets []Widget
		want    image.Point
	}{
		{
			[]Widget{
				NewBox(NewFill(image.Point{10, 10}, color.White), Margin{}, Padding{}),
				NewBox(NewFill(image.Point{10, 10}, color.White), Margin{}, Padding{}),
			},
			image.Point{20, 10},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprint(i+1), func(t *testing.T) {
			s := NewHorizontalStack(tt.widgets...)
			size := s.Size()
			if size != tt.want {
				t.Errorf("Size should return %v, but got %v", tt.want, size)
			}
		})
	}
}

func TestVerticalStackSize(t *testing.T) {
	tests := []struct {
		widgets []Widget
		want    image.Point
	}{
		{
			[]Widget{
				NewBox(NewFill(image.Point{10, 10}, color.White), Margin{}, Padding{}),
				NewBox(NewFill(image.Point{10, 10}, color.White), Margin{}, Padding{}),
			},
			image.Point{10, 20},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprint(i+1), func(t *testing.T) {
			s := NewVerticalStack(tt.widgets...)
			size := s.Size()
			if size != tt.want {
				t.Errorf("Size should return %v, but got %v", tt.want, size)
			}
		})
	}
}

func TestLayerSize(t *testing.T) {
	tests := []struct {
		widgets []Widget
		want    image.Point
	}{
		{
			[]Widget{
				NewFill(image.Point{10, 30}, color.White),
				NewFill(image.Point{20, 20}, color.White),
				NewFill(image.Point{30, 10}, color.White),
			},
			image.Point{30, 30},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprint(i+1), func(t *testing.T) {
			l := NewLayer(tt.widgets...)
			size := l.Size()
			if size != tt.want {
				t.Errorf("Size should return %v, but got %v", tt.want, size)
			}
		})
	}
}
