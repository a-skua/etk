package main

import (
	"bytes"
	"image"
	"image/color"
	_ "image/jpeg"
	"log"
	"time"

	"github.com/a-skua/etk"
	"github.com/a-skua/etk/widget"
	"github.com/hajimehoshi/bitmapfont/v3"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/images"
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
		&etk.DefaultScene{
			Widget: widget.NewStack(
				widget.Vertical,
				widget.NewBox(
					widget.NewLayer(
						widget.NewBox(
							widget.NewFill(image.Point{100, 100}, color.Gray{0xff}),
							widget.Margin{},
							widget.Padding{},
						),
						widget.NewBox(
							widget.NewFill(image.Point{80, 80}, color.Gray{0x88}),
							widget.Margin{Top: 10, Left: 10},
							widget.Padding{},
						),
						widget.NewBox(
							widget.NewFill(image.Point{60, 60}, color.Gray{0x00}).AddText("Layter", color.White),
							widget.Margin{Top: 20, Left: 20},
							widget.Padding{},
						),
					),
					widget.Margin{Left: 10, Top: 10},
					widget.Padding{},
				),
				widget.NewBox(
					widget.NewFill(image.Point{80, 80}, color.White),
					widget.Margin{Left: 10, Top: 10},
					widget.Padding{},
				),
				widget.NewBox(
					widget.NewFill(image.Point{40, 40}, color.White),
					widget.Margin{Left: 10, Top: 10},
					widget.Padding{},
				),
				widget.NewBox(
					widget.NewFill(image.Point{100, 16}, color.Gray{0x88}).AddText("Vertical", color.White),
					widget.Margin{Left: 10, Top: 10},
					widget.Padding{},
				),
				widget.NewStack(
					widget.Horizontal,
					widget.NewBox(
						widget.NewFill(image.Point{100, 100}, color.White),
						widget.Margin{Left: 10, Top: 10},
						widget.Padding{},
					),
					widget.NewBox(
						widget.NewFill(image.Point{100, 16}, color.Gray{0x88}).AddText("Horizontal", color.White),
						widget.Margin{Left: 10, Top: 10},
						widget.Padding{},
					),
					widget.NewBox(
						widget.NewFill(image.Point{80, 80}, color.White),
						widget.Margin{Left: 10, Top: 10},
						widget.Padding{},
					),
					widget.NewBox(
						widget.NewImage(ebitenPng),
						widget.Margin{Left: 10, Top: 10},
						widget.Padding{},
					),
				),
			),
		},
	}
	if err := ebiten.RunGame(etk.New(screenWidth, screenHeight, scenes[0], scenes[1:]...).Debug()); err != nil {
		log.Fatal(err)
	}
}
