package types

import (
	"fmt"
	"testing"
)

func TestMarginAll(t *testing.T) {
	tests := []struct {
		margin Margin
		want   Margin
	}{
		{
			MarginAll(10),
			Margin{10, 10, 10, 10},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprint(i+1), func(t *testing.T) {
			if tt.margin != tt.want {
				t.Errorf("MarginAll should return %v, but got %v", tt.want, tt.margin)
			}
		})
	}
}

func TestMargin_pos(t *testing.T) {
	tests := []struct {
		margin Margin
		want   Position
	}{
		{
			Margin{10, 20, 30, 40},
			Position{10, 20},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprint(i+1), func(t *testing.T) {
			p := tt.margin.Pos()
			if p != tt.want {
				t.Errorf("pos should return %v, but got %v", tt.want, p)
			}
		})
	}
}

func TestPosition_sub(t *testing.T) {
	tests := []struct {
		p    Position
		q    Position
		want Position
	}{
		{
			Position{10, 20},
			Position{5, 10},
			Position{5, 10},
		},
		{
			Position{10, 20},
			Position{20, 25},
			Position{-10, -5},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprint(i+1), func(t *testing.T) {
			p := tt.p.Sub(tt.q)
			if p != tt.want {
				t.Errorf("sub should return %v, but got %v", tt.want, p)
			}
		})
	}
}

func TestPosition_add(t *testing.T) {
	tests := []struct {
		p    Position
		q    Position
		want Position
	}{
		{
			Position{10, 20},
			Position{5, 10},
			Position{15, 30},
		},
		{
			Position{10, 20},
			Position{20, 25},
			Position{30, 45},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprint(i+1), func(t *testing.T) {
			p := tt.p.Add(tt.q)
			if p != tt.want {
				t.Errorf("add should return %v, but got %v", tt.want, p)
			}
		})
	}
}
