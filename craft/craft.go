package craft

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

// Craft
type Craft interface {
	Image() *ebiten.Image
	Size() Size
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

func (i *Image) Size() Size {
	return Size(i.image.Bounds().Size())
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
	size  Size
	color color.Color
	texts []text
}

func NewFill(size Size, color color.Color) *Fill {
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

func (f *Fill) Size() Size {
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
	craft  Craft
	margin Margin
	texts  []text
}

func NewBox(c Craft, m Margin) *Box {
	return &Box{c, m, []text{}}
}

func (b *Box) Image() *ebiten.Image {
	size := b.Size()

	image := ebiten.NewImage(size.X, size.Y)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(b.margin.Left), float64(b.margin.Top))
	image.DrawImage(b.craft.Image(), op)

	for _, text := range b.texts {
		drawText(image, text.str, text.color)
	}

	return image
}

func (b *Box) Size() Size {
	return calcSize(b.craft.Size(), b.margin)
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

func (s *HorizontalStack) Size() Size {
	width, height := 0, 0
	for _, c := range s.crafts {
		size := c.Size()
		width += size.X
		height = max(height, size.Y)
	}
	return Size{width, height}
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

func (s *VerticalStack) Size() Size {
	width, height := 0, 0
	for _, c := range s.crafts {
		size := c.Size()
		width = max(width, size.X)
		height += size.Y
	}
	return Size{width, height}
}

func (s *VerticalStack) AddText(str string, color color.Color) Craft {
	s.texts = append(s.texts, text{str, color})
	return s
}

func (s *VerticalStack) Const() *Image {
	return &Image{s.Image()}
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

func (l *Layer) Size() Size {
	width, height := 0, 0
	for _, c := range l.crafts {
		img := c.Image()
		size := img.Bounds().Size()
		width = max(width, size.X)
		height = max(height, size.Y)
	}

	return Size{width, height}
}

func (l *Layer) AddText(str string, color color.Color) Craft {
	l.texts = append(l.texts, text{str, color})
	return l
}

func (l *Layer) Const() *Image {
	return &Image{l.Image()}
}
