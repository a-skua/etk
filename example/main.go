package main

import (
	"bytes"
	"image"
	"image/color"
	_ "image/jpeg"
	"log"
	"time"

	"github.com/a-skua/etk"
	"github.com/a-skua/etk/craft"
	"github.com/hajimehoshi/bitmapfont/v3"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/images"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

type scene1 struct {
	etk.DefaultScene
	time     time.Time
	startX   int
	startY   int
	moveX    float32
	moveY    float32
	currentX int
	currentY int
	moveTime time.Duration
}

func (s *scene1) Init() {
	s.time = time.Now()
	s.currentX = s.startX
	s.currentY = s.startY
}

func (s *scene1) Update() error {
	current := time.Since(s.time)
	if current < s.moveTime {
		wait := float32(current) / float32(s.moveTime)
		s.currentX = s.startX + int(float32(s.moveX)*wait)
		s.currentY = s.startY + int(float32(s.moveY)*wait)
	}
	return nil
}

func (s *scene1) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x80, 0x80, 0x80, 0xff})
	text.Draw(screen, "Hello, 世界!", bitmapfont.Face, s.currentX, s.currentY, color.White)
}

type sceneSwitcher struct {
	etk.DefaultScene
}

func (s *sceneSwitcher) Previous() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft)
}

func (s *sceneSwitcher) Next() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeyArrowRight)
}

func main() {
	const (
		screenWidth  = 960
		screenHeight = 540
	)

	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Hello, World!")

	ebitenPng, _, err := image.Decode(bytes.NewReader(images.Ebiten_png))
	if err != nil {
		log.Fatal(err)
	}

	scenes := []etk.Scene{
		&etk.DefaultScene{},
		&scene1{
			startX:   0,
			startY:   0,
			moveX:    100,
			moveY:    100,
			moveTime: 5 * time.Second,
		},
		&sceneSwitcher{
			DefaultScene: etk.DefaultScene{
				Craft: craft.NewBox(
					craft.NewFill(image.Point{100, 32}, color.Gray{0x88}).AddText("Prev: ←\nNext: →", color.White),
					craft.Margin{Top: 100, Left: 100},
					craft.Padding{},
				).Const(),
			},
		},
		&etk.DefaultScene{
			Craft: craft.NewVerticalStack(
				craft.NewBox(
					craft.NewLayer(
						craft.NewBox(
							craft.NewFill(image.Point{100, 100}, color.Gray{0xff}),
							craft.Margin{},
							craft.Padding{},
						),
						craft.NewBox(
							craft.NewFill(image.Point{80, 80}, color.Gray{0x88}),
							craft.Margin{Top: 10, Left: 10},
							craft.Padding{},
						),
						craft.NewBox(
							craft.NewFill(image.Point{60, 60}, color.Gray{0x00}).AddText("Layter", color.White),
							craft.Margin{Top: 20, Left: 20},
							craft.Padding{},
						),
					),
					craft.Margin{Left: 10, Top: 10},
					craft.Padding{},
				),
				craft.NewBox(
					craft.NewFill(image.Point{80, 80}, color.White),
					craft.Margin{Left: 10, Top: 10},
					craft.Padding{},
				),
				craft.NewBox(
					craft.NewFill(image.Point{40, 40}, color.White),
					craft.Margin{Left: 10, Top: 10},
					craft.Padding{},
				),
				craft.NewBox(
					craft.NewFill(image.Point{100, 16}, color.Gray{0x88}).AddText("Vertical", color.White),
					craft.Margin{Left: 10, Top: 10},
					craft.Padding{},
				),
				craft.NewHorizontalStack(
					craft.NewBox(
						craft.NewFill(image.Point{100, 100}, color.White),
						craft.Margin{Left: 10, Top: 10},
						craft.Padding{},
					),
					craft.NewBox(
						craft.NewFill(image.Point{100, 16}, color.Gray{0x88}).AddText("Horizontal", color.White),
						craft.Margin{Left: 10, Top: 10},
						craft.Padding{},
					),
					craft.NewBox(
						craft.NewFill(image.Point{80, 80}, color.White),
						craft.Margin{Left: 10, Top: 10},
						craft.Padding{},
					),
					craft.NewBox(
						craft.NewImage(ebitenPng),
						craft.Margin{Left: 10, Top: 10},
						craft.Padding{},
					),
				),
			).Const(),
		},
	}
	if err := ebiten.RunGame(etk.New(screenWidth, screenHeight, scenes[0], scenes[1:]...).Debug()); err != nil {
		log.Fatal(err)
	}
}
