package main

import (
	"image/color"
	"log"
	"time"

	"github.com/a-skua/etk"
	"github.com/hajimehoshi/bitmapfont/v3"
	"github.com/hajimehoshi/ebiten/v2"
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
	text.Draw(screen, "Hello, World!", bitmapfont.Face, s.currentX, s.currentY, color.White)
}

type scene2 struct {
	etk.DefaultScene
}

func (s *scene2) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0xaa, 0xaa, 0xaa, 0x00})
	text.Draw(screen, "Goodbye, 世界!", bitmapfont.Face, 0, 12, color.White)
}

func main() {
	const (
		screenWidth  = 960
		screenHeight = 540
	)

	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Hello, World!")
	scenes := []etk.Scene{
		&etk.DefaultScene{},
		&scene1{
			startX:   0,
			startY:   0,
			moveX:    100,
			moveY:    100,
			moveTime: 5 * time.Second,
		},
		&scene2{},
	}
	if err := ebiten.RunGame(etk.New(screenWidth, screenHeight, scenes[0], scenes[1:]...).Debug()); err != nil {

		log.Fatal(err)
	}
}
