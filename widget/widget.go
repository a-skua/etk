package widget

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/bitmapfont/v3"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

type Widget interface {
	Image() *ebiten.Image
	Position() image.Point
	Size() image.Point
	AddText(string, color.Color) Widget
}

func calcSize(size image.Point, margin Margin, padding Padding) image.Point {
	width := size.X + margin.Left + margin.Right + padding.Left + padding.Right
	height := size.Y + margin.Top + margin.Bottom + padding.Top + padding.Bottom
	return image.Point{width, height}
}

func calcPosition(margin Margin, padding Padding) image.Point {
	x := margin.Left + padding.Left
	y := margin.Top + padding.Top
	return image.Point{x, y}
}

func drawText(image *ebiten.Image, str string, color color.Color) {
	text.Draw(image, str, bitmapfont.Face, 0, 12, color)
}

// Box Widget
type Box struct {
	image   *ebiten.Image
	margin  Margin
	padding Padding
}

type Margin space

type Padding space

type space struct {
	Left, Top, Right, Bottom int
}

func NewBox(widget Widget, margin Margin, padding Padding) *Box {
	size := calcSize(widget.Size(), margin, padding)

	image := ebiten.NewImage(size.X, size.Y)
	op := &ebiten.DrawImageOptions{}

	position := calcPosition(margin, padding)
	op.GeoM.Translate(float64(position.X), float64(position.Y))
	image.DrawImage(widget.Image(), op)
	return &Box{image, margin, padding}
}

func (b *Box) Position() image.Point {
	return calcPosition(b.margin, b.padding)
}

func (b *Box) Image() *ebiten.Image {
	return b.image
}

func (b *Box) Size() image.Point {
	return b.image.Bounds().Size()
}

func (b *Box) AddText(str string, color color.Color) Widget {
	drawText(b.image, str, color)
	return b
}

// Stack Widget
type Stack struct {
	image *ebiten.Image
}

// Direction is the direction of the stack
type StackDirection int

const (
	// Horizontal stack
	Horizontal StackDirection = iota
	// Vertical stack
	Vertical
)

func newHorizontalStack(widgets ...Widget) *Stack {
	width, height := 0, 0
	images := make([]*ebiten.Image, 0, len(widgets))
	for _, w := range widgets {
		img := w.Image()
		size := img.Bounds().Size()
		width += size.X
		height = max(height, size.Y)
		images = append(images, img)
	}

	image := ebiten.NewImage(width, height)
	x, y := 0.0, 0.0
	for _, img := range images {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(x, y)
		image.DrawImage(img, op)
		x += float64(img.Bounds().Size().X)
	}
	return &Stack{image: image}
}

func newVerticalStack(widgets ...Widget) *Stack {
	width, height := 0, 0
	images := make([]*ebiten.Image, 0, len(widgets))
	for _, w := range widgets {
		img := w.Image()
		size := img.Bounds().Size()
		width = max(width, size.X)
		height += size.Y
		images = append(images, img)
	}

	image := ebiten.NewImage(width, height)
	x, y := 0.0, 0.0
	for _, img := range images {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(x, y)
		image.DrawImage(img, op)
		y += float64(img.Bounds().Size().Y)
	}
	return &Stack{image: image}
}

func NewStack(direction StackDirection, widgets ...Widget) *Stack {
	if direction == Horizontal {
		return newHorizontalStack(widgets...)
	}
	return newVerticalStack(widgets...)
}

func (s *Stack) Image() *ebiten.Image {
	return s.image
}

func (s *Stack) Position() image.Point {
	return image.Point{}
}

func (s *Stack) Size() image.Point {
	return s.image.Bounds().Size()
}

func (s *Stack) AddText(str string, color color.Color) Widget {
	drawText(s.image, str, color)
	return s
}

// Layer Widget
type Layer struct {
	image *ebiten.Image
}

func NewLayer(widgets ...Widget) *Layer {
	width, height := 0, 0
	images := make([]*ebiten.Image, 0, len(widgets))
	for _, w := range widgets {
		img := w.Image()
		size := img.Bounds().Size()
		width = max(width, size.X)
		height = max(height, size.Y)
		images = append(images, img)
	}

	image := ebiten.NewImage(width, height)
	for _, img := range images {
		op := &ebiten.DrawImageOptions{}
		image.DrawImage(img, op)
	}

	return &Layer{image}
}

func (l *Layer) Image() *ebiten.Image {
	return l.image
}

func (l *Layer) Position() image.Point {
	return image.Point{}
}

func (l *Layer) Size() image.Point {
	return l.image.Bounds().Size()
}

func (l *Layer) AddText(str string, color color.Color) Widget {
	drawText(l.image, str, color)
	return l
}

// Image Widget
type Image struct {
	image *ebiten.Image
}

func NewImage(img image.Image) *Image {
	image := ebiten.NewImageFromImage(img)
	return &Image{
		image,
	}
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

func (i *Image) AddText(str string, color color.Color) Widget {
	drawText(i.image, str, color)
	return i
}

// Fill Widget
type Fill struct {
	image *ebiten.Image
}

func NewFill(size image.Point, color color.Color) *Fill {
	image := ebiten.NewImage(size.X, size.Y)
	image.Fill(color)
	return &Fill{image}
}

func (f *Fill) Image() *ebiten.Image {
	return f.image
}

func (f *Fill) Position() image.Point {
	return image.Point{}
}

func (f *Fill) Size() image.Point {
	return f.image.Bounds().Size()
}

func (f *Fill) AddText(str string, color color.Color) Widget {
	drawText(f.image, str, color)
	return f
}
