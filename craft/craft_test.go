package craft

import (
	"fmt"
	"image"
	"image/color"
	"testing"
)

var _ Craft = NewImage(image.NewRGBA(image.Rect(0, 0, 100, 100)))
var _ Craft = NewFill(Size{10, 10}, color.White)
var _ Craft = NewBox(NewFill(Size{10, 10}, color.White), Margin{})
var _ Craft = NewHorizontalStack(
	NewBox(NewFill(Size{10, 10}, color.White), Margin{}),
	NewBox(NewFill(Size{10, 10}, color.White), Margin{}),
)
var _ Craft = NewVerticalStack(
	NewBox(NewFill(Size{10, 10}, color.White), Margin{}),
	NewBox(NewFill(Size{10, 10}, color.White), Margin{}),
)
var _ Craft = NewLayer(
	NewFill(Size{10, 20}, color.White),
	NewFill(Size{20, 10}, color.White),
)

func TestHorizontalStackSize(t *testing.T) {
	tests := []struct {
		widgets []Craft
		want    Size
	}{
		{
			[]Craft{
				NewBox(NewFill(Size{10, 10}, color.White), Margin{}),
				NewBox(NewFill(Size{10, 10}, color.White), Margin{}),
			},
			Size{20, 10},
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
		widgets []Craft
		want    Size
	}{
		{
			[]Craft{
				NewBox(NewFill(Size{10, 10}, color.White), Margin{}),
				NewBox(NewFill(Size{10, 10}, color.White), Margin{}),
			},
			Size{10, 20},
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
		widgets []Craft
		want    Size
	}{
		{
			[]Craft{
				NewFill(Size{10, 30}, color.White),
				NewFill(Size{20, 20}, color.White),
				NewFill(Size{30, 10}, color.White),
			},
			Size{30, 30},
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
