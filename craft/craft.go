package craft

import (
	"errors"
	"image"
	"image/color"

	"github.com/a-skua/etk/craft/internal/util"
	"github.com/a-skua/etk/craft/types"
	"github.com/hajimehoshi/ebiten/v2"
)

// Craft
type Craft interface {
	Image() *ebiten.Image
	Size() types.Size
	AddText(string, color.Color) Self
	Const() *Image
	Update(types.Position) error
}

type Self = Craft

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

func (i *Image) Size() types.Size {
	return types.Size(i.image.Bounds().Size())
}

func (i *Image) AddText(str string, color color.Color) Self {
	util.DrawText(i.image, str, color)
	return i
}

func (i *Image) Const() *Image {
	return i
}

func (i *Image) Update(p types.Position) error {
	return nil
}

// Fill Craft
type Fill struct {
	size  types.Size
	color color.Color
	texts []types.TextInfo
}

func NewFill(size types.Size, color color.Color) *Fill {
	return &Fill{size, color, []types.TextInfo{}}
}

func (f *Fill) Image() *ebiten.Image {
	image := ebiten.NewImage(f.size.X, f.size.Y)
	image.Fill(f.color)

	for _, t := range f.texts {
		util.DrawText(image, t.Str, t.Color)
	}

	return image
}

func (f *Fill) Size() types.Size {
	return f.size
}

func (f *Fill) AddText(str string, color color.Color) Self {
	f.texts = append(f.texts, types.TextInfo{Str: str, Color: color})
	return f
}

func (f *Fill) Const() *Image {
	return &Image{f.Image()}
}

func (f *Fill) Update(p types.Position) error {
	return nil
}

// Switch Craft
type Switch struct {
	crafts []Craft
	index  int
	texts  []types.TextInfo
}

func NewSwitch(crafts ...Craft) *Switch {
	return &Switch{crafts, 0, []types.TextInfo{}}
}

func (s *Switch) Image() *ebiten.Image {
	return s.crafts[s.index].Image()
}

func (s *Switch) Size() types.Size {
	return s.crafts[s.index].Size()
}

func (s *Switch) AddText(str string, color color.Color) Self {
	s.texts = append(s.texts, types.TextInfo{Str: str, Color: color})
	return s
}

func (s *Switch) Update(p types.Position) error {
	return s.crafts[s.index].Update(p)
}

func (s *Switch) Const() *Image {
	return &Image{s.Image()}
}

func (s *Switch) Next() {
	s.index++
	if s.index >= len(s.crafts) {
		s.index = 0
	}
}

func (s *Switch) Prev() {
	s.index--
	if s.index < 0 {
		s.index = len(s.crafts) - 1
	}
}

// Box Craft
type Box struct {
	craft  Craft
	margin types.Margin
	texts  []types.TextInfo
}

func NewBox(c Craft, m types.Margin) *Box {
	return &Box{c, m, []types.TextInfo{}}
}

func (b *Box) Image() *ebiten.Image {
	size := b.Size()

	image := ebiten.NewImage(size.X, size.Y)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(b.margin.Left), float64(b.margin.Top))
	image.DrawImage(b.craft.Image(), op)

	for _, t := range b.texts {
		util.DrawText(image, t.Str, t.Color)
	}

	return image
}

func (b *Box) Size() types.Size {
	return util.CalcSize(b.craft.Size(), b.margin)
}

func (b *Box) AddText(str string, color color.Color) Self {
	b.texts = append(b.texts, types.TextInfo{Str: str, Color: color})
	return b
}

func (b *Box) Update(p types.Position) error {
	return b.craft.Update(p.Add(b.margin.Pos()))
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
	texts  []types.TextInfo
}

func NewHorizontalStack(crafts ...Craft) *HorizontalStack {
	return &HorizontalStack{crafts, []types.TextInfo{}}
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

	for _, t := range s.texts {
		util.DrawText(image, t.Str, t.Color)
	}

	return image
}

func (s *HorizontalStack) Size() types.Size {
	size := types.Size{}
	for _, c := range s.crafts {
		s := c.Size()
		size.X += s.X
		size.Y = max(size.Y, s.Y)
	}
	return size
}

func (s *HorizontalStack) AddText(str string, color color.Color) Self {
	s.texts = append(s.texts, types.TextInfo{Str: str, Color: color})
	for _, c := range s.crafts {
		c.AddText(str, color)
	}
	return s
}

func (s *HorizontalStack) Update(p types.Position) (err error) {
	for _, c := range s.crafts {
		err = errors.Join(err, c.Update(p))
		p.X += c.Size().X
	}
	return
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
	texts  []types.TextInfo
}

func NewVerticalStack(crafts ...Craft) *VerticalStack {
	return &VerticalStack{crafts, []types.TextInfo{}}
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

	for _, t := range s.texts {
		util.DrawText(image, t.Str, t.Color)
	}

	return image
}

func (s *VerticalStack) Size() types.Size {
	size := types.Size{}
	for _, c := range s.crafts {
		s := c.Size()
		size.X = max(size.X, s.X)
		size.Y += s.Y
	}
	return size
}

func (s *VerticalStack) AddText(str string, color color.Color) Self {
	s.texts = append(s.texts, types.TextInfo{Str: str, Color: color})
	return s
}

func (s *VerticalStack) Const() *Image {
	return &Image{s.Image()}
}

func (s *VerticalStack) Update(p types.Position) (err error) {
	for _, c := range s.crafts {
		err = errors.Join(err, c.Update(p))
		p.Y += c.Size().Y
	}
	return
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
	texts  []types.TextInfo
}

func NewLayer(crafts ...Craft) *Layer {
	return &Layer{crafts, []types.TextInfo{}}
}

func (l *Layer) Image() *ebiten.Image {
	size := l.Size()

	image := ebiten.NewImage(size.X, size.Y)
	for _, w := range l.crafts {
		op := &ebiten.DrawImageOptions{}
		image.DrawImage(w.Image(), op)
	}

	for _, t := range l.texts {
		util.DrawText(image, t.Str, t.Color)
	}

	return image
}

func (l *Layer) Size() types.Size {
	size := types.Size{}
	for _, c := range l.crafts {
		img := c.Image()
		s := img.Bounds().Size()
		size.X = max(size.X, s.X)
		size.Y = max(size.Y, s.Y)
	}

	return size
}

func (l *Layer) AddText(str string, color color.Color) Self {
	l.texts = append(l.texts, types.TextInfo{Str: str, Color: color})
	return l
}

func (l *Layer) Const() *Image {
	return &Image{l.Image()}
}

func (l *Layer) Update(p types.Position) (err error) {
	for _, c := range l.crafts {
		err = errors.Join(err, c.Update(p))
	}
	return
}
