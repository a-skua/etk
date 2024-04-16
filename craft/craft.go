package craft

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Craft interface {
	Image() *ebiten.Image
	Position() image.Point
	Size() image.Point
	AddText(string, color.Color) Craft
	Const() *Image
}

type text struct {
	str   string
	color color.Color
}

// Image Craft
type Image struct {
	image *ebiten.Image
}

func NewImage(img image.Image) *Image {
	image := ebiten.NewImageFromImage(img)
	return &Image{image}
}

func (i *Image) Image() *ebiten.Image {
	return i.image
}

func (i *Image) Position() image.Point {
	return image.Point{}
}

func (i *Image) Size() image.Point {
	return i.image.Bounds().Size()
}

func (i *Image) AddText(str string, color color.Color) Craft {
	drawText(i.image, str, color)
	return i
}

func (i *Image) Const() *Image {
	return i
}

// Fill Craft
type Fill struct {
	size  image.Point
	color color.Color
	texts []text
}

func NewFill(size image.Point, color color.Color) *Fill {
	return &Fill{size, color, []text{}}
}

func (f *Fill) Image() *ebiten.Image {
	image := ebiten.NewImage(f.size.X, f.size.Y)
	image.Fill(f.color)

	for _, text := range f.texts {
		drawText(image, text.str, text.color)
	}

	return image
}

func (f *Fill) Position() image.Point {
	return image.Point{}
}

func (f *Fill) Size() image.Point {
	return f.size
}

func (f *Fill) AddText(str string, color color.Color) Craft {
	f.texts = append(f.texts, text{str, color})
	return f
}

func (f *Fill) Const() *Image {
	return &Image{f.Image()}
}

// Box Craft
type Box struct {
	craft   Craft
	margin  Margin
	padding Padding
	texts   []text
}

type Margin space

type Padding space

type space struct {
	Left, Top, Right, Bottom int
}

func NewBox(craft Craft, margin Margin, padding Padding) *Box {
	return &Box{craft, margin, padding, []text{}}
}

func (b *Box) Position() image.Point {
	return calcPosition(b.margin, b.padding)
}

func (b *Box) Image() *ebiten.Image {
	size := b.Size()

	image := ebiten.NewImage(size.X, size.Y)
	op := &ebiten.DrawImageOptions{}
	position := calcPosition(b.margin, b.padding)
	op.GeoM.Translate(float64(position.X), float64(position.Y))
	image.DrawImage(b.craft.Image(), op)

	for _, text := range b.texts {
		drawText(image, text.str, text.color)
	}

	return image
}

func (b *Box) Size() image.Point {
	return calcSize(b.craft.Size(), b.margin, b.padding)

}

func (b *Box) AddText(str string, color color.Color) Craft {
	b.texts = append(b.texts, text{str, color})
	return b
}

func (b *Box) Const() *Image {
	return &Image{b.Image()}
}

// HorizontalStack Craft
//
// ```
// +---+---+---+
// |   |   |   | --> horizontal
// +---+---+---+
// ```
type HorizontalStack struct {
	crafts []Craft
	texts  []text
}

func NewHorizontalStack(crafts ...Craft) *HorizontalStack {
	return &HorizontalStack{crafts, []text{}}
}

func (s *HorizontalStack) Image() *ebiten.Image {
	size := s.Size()

	image := ebiten.NewImage(size.X, size.Y)
	x, y := 0.0, 0.0
	for _, c := range s.crafts {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(x, y)
		image.DrawImage(c.Image(), op)
		x += float64(c.Size().X)
	}

	for _, text := range s.texts {
		drawText(image, text.str, text.color)
	}

	return image
}

func (s *HorizontalStack) Position() image.Point {
	return image.Point{}
}

func (s *HorizontalStack) Size() image.Point {
	width, height := 0, 0
	for _, c := range s.crafts {
		size := c.Size()
		width += size.X
		height = max(height, size.Y)
	}
	return image.Point{width, height}
}

func (s *HorizontalStack) AddText(str string, color color.Color) Craft {
	s.texts = append(s.texts, text{str, color})
	for _, c := range s.crafts {
		c.AddText(str, color)
	}
	return s
}

func (s *HorizontalStack) Const() *Image {
	return &Image{s.Image()}
}

// VerticalStack Craft
//
// ```
// +---+
// |   |
// +---+
// |   |
// +---+
// ..| vertical
// ..v
//
// ```
type VerticalStack struct {
	crafts []Craft
	texts  []text
}

func NewVerticalStack(crafts ...Craft) *VerticalStack {
	return &VerticalStack{crafts, []text{}}
}

func (s *VerticalStack) Image() *ebiten.Image {
	size := s.Size()

	image := ebiten.NewImage(size.X, size.Y)
	x, y := 0.0, 0.0
	for _, c := range s.crafts {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(x, y)
		image.DrawImage(c.Image(), op)
		y += float64(c.Size().Y)
	}

	for _, text := range s.texts {
		drawText(image, text.str, text.color)
	}

	return image
}

func (s *VerticalStack) Position() image.Point {
	return image.Point{}
}

func (s *VerticalStack) Size() image.Point {
	width, height := 0, 0
	for _, c := range s.crafts {
		size := c.Size()
		width = max(width, size.X)
		height += size.Y
	}
	return image.Point{width, height}
}

func (s *VerticalStack) AddText(str string, color color.Color) Craft {
	s.texts = append(s.texts, text{str, color})
	return s
}

func (s *VerticalStack) Const() *Image {
	return &Image{s.Image()}
}

// Stack Craft
//
// Deprecated: Use HorizontalStack or VerticalStack instead.
type Stack struct {
	Craft
}

// Direction is the direction of the stack
type StackDirection int

const (
	// Horizontal stack
	Horizontal StackDirection = iota
	// Vertical stack
	Vertical
)

// NewStack creates a new stack widget.
//
// Deprecated: Use NewHorizontalStack or NewVerticalStack instead.
func NewStack(direction StackDirection, crafts ...Craft) *Stack {
	if direction == Horizontal {
		return &Stack{NewHorizontalStack(crafts...)}
	}
	return &Stack{NewVerticalStack(crafts...)}
}

// Layer Craft
//
// ```
// +------+
// |      |
// |      |---+
// +------+   |
// ....|      |
// ....+------+ --> layer
// ```
type Layer struct {
	crafts []Craft
	texts  []text
}

func NewLayer(crafts ...Craft) *Layer {
	return &Layer{crafts, []text{}}
}

func (l *Layer) Image() *ebiten.Image {
	size := l.Size()

	image := ebiten.NewImage(size.X, size.Y)
	for _, w := range l.crafts {
		op := &ebiten.DrawImageOptions{}
		image.DrawImage(w.Image(), op)
	}

	for _, text := range l.texts {
		drawText(image, text.str, text.color)
	}

	return image
}

func (l *Layer) Position() image.Point {
	return image.Point{}
}

func (l *Layer) Size() image.Point {
	width, height := 0, 0
	for _, c := range l.crafts {
		img := c.Image()
		size := img.Bounds().Size()
		width = max(width, size.X)
		height = max(height, size.Y)
	}

	return image.Point{width, height}
}

func (l *Layer) AddText(str string, color color.Color) Craft {
	l.texts = append(l.texts, text{str, color})
	return l
}

func (l *Layer) Const() *Image {
	return &Image{l.Image()}
}
