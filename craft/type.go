package craft

import (
	"image"
)

type Size image.Point

type Margin struct {
	Left, Top, Right, Bottom int
}

func MarginAll(v int) Margin {
	return Margin{v, v, v, v}
}
