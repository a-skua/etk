package widget

import (
	"fmt"
	"image"
	"image/color"
	"testing"
)

var testFillType Widget = NewFill(image.Point{10, 10}, color.White)
var testBoxType Widget = NewBox(testFillType, Margin{}, Padding{})
var testStackType Widget = NewStack(Vertical, testBoxType, testBoxType)
var testImageType Widget = NewImage(image.NewRGBA(image.Rect(0, 0, 100, 100)))
var testLayerType Widget = NewLayer(testFillType, testFillType)

func Test_calcSize(t *testing.T) {
	tests := []struct {
		size    image.Point
		margin  Margin
		padding Padding
		want    image.Point
	}{
		{image.Point{100, 50}, Margin{}, Padding{}, image.Point{100, 50}},
		{image.Point{100, 50}, Margin{10, 0, 10, 0}, Padding{10, 0, 10, 0}, image.Point{140, 50}},
		{image.Point{100, 50}, Margin{0, 10, 0, 10}, Padding{0, 10, 0, 10}, image.Point{100, 90}},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprint(i+1), func(t *testing.T) {
			size := calcSize(tt.size, tt.margin, tt.padding)
			if size != tt.want {
				t.Errorf("calcSize should return %v, but got %v", tt.want, size)
			}
		})
	}
}

func Test_calcPosition(t *testing.T) {
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
			p := calcPosition(tt.margin, tt.padding)
			if p != tt.want {
				t.Errorf("calcPosition should return %v, but got %v", tt.want, p)
			}
		})
	}
}

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
