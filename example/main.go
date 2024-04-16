package main

import (
	"bytes"
	"image"
	"image/color"
	_ "image/jpeg"
	"log"

	"github.com/a-skua/etk"
	"github.com/a-skua/etk/craft"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/images"
)

func main() {
	const (
		screenWidth  = 640
		screenHeight = 360
	)

	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Hello, World!")

	ebitenPng, _, err := image.Decode(bytes.NewReader(images.Ebiten_png))
	if err != nil {
		log.Fatal(err)
	}

	game := etk.New(
		screenWidth,
		screenHeight,
		&etk.DefaultScene{
			Craft: craft.NewVerticalStack(
				craft.NewBox(
					craft.NewFill(craft.Size{X: 10, Y: 10}, color.White),
					craft.MarginAll(10),
				),
				craft.NewBox(
					craft.NewImage(ebitenPng),
					craft.MarginAll(10),
				),
			).Const(),
		},
	)

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
