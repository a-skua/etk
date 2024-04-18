package util

import (
	"fmt"
	"testing"

	"github.com/a-skua/etk/craft/types"
)

func TestCalcSize(t *testing.T) {
	tests := []struct {
		size   types.Size
		margin types.Margin
		want   types.Size
	}{
		{types.Size{X: 100, Y: 50}, types.Margin{}, types.Size{X: 100, Y: 50}},
		{types.Size{X: 100, Y: 50}, types.Margin{Left: 10, Right: 10}, types.Size{X: 120, Y: 50}},
		{types.Size{X: 100, Y: 50}, types.Margin{Top: 10, Bottom: 10}, types.Size{X: 100, Y: 70}},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprint(i+1), func(t *testing.T) {
			size := CalcSize(tt.size, tt.margin)
			if size != tt.want {
				t.Errorf("calcSize should return %v, but got %v", tt.want, size)
			}
		})
	}
}

func TestSizein(t *testing.T) {
	tests := []struct {
		size types.Size
		p    types.Position
		want bool
	}{
		{types.Size{X: 100, Y: 50}, types.Position{X: 10, Y: 20}, true},
		{types.Size{X: 100, Y: 50}, types.Position{X: 100, Y: 50}, false},
		{types.Size{X: 100, Y: 50}, types.Position{X: 101, Y: 50}, false},
		{types.Size{X: 100, Y: 50}, types.Position{X: 100, Y: 51}, false},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprint(i+1), func(t *testing.T) {
			b := Sizein(tt.size, tt.p)
			if b != tt.want {
				t.Errorf("sizein should return %v, but got %v", tt.want, b)
			}
		})
	}
}
