package types

import (
	"image"
	"image/color"
)

type TextInfo struct {
	Str   string
	Color color.Color
}

// Size
type Size image.Point

// Margin
type Margin struct {
	Left, Top, Right, Bottom int
}

func MarginAll(v int) Margin {
	return Margin{v, v, v, v}
}

func (m Margin) Pos() Position {
	return Position{m.Left, m.Top}
}

// Position
type Position image.Point

func toPosition(x, y int) Position {
	return Position{X: x, Y: y}
}

func (p Position) Sub(q Position) Position {
	return Position{X: p.X - q.X, Y: p.Y - q.Y}
}

func (p Position) Add(q Position) Position {
	return Position{X: p.X + q.X, Y: p.Y + q.Y}
}
