package main

import (
	"bytes"
	"image"
	"image/color"
	_ "image/jpeg"
	"log"
	"log/slog"

	"github.com/a-skua/etk"
	"github.com/a-skua/etk/craft"
	"github.com/a-skua/etk/craft/action"
	"github.com/a-skua/etk/craft/types"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/images"
)

func init() {
	slog.SetLogLoggerLevel(slog.LevelDebug)
}

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
					action.NewMousePressed(
						craft.NewFill(types.Size{X: 10, Y: 10}, color.White),
						ebiten.MouseButtonLeft,
						func(f *craft.Fill) error {
							*f = *craft.NewFill(f.Size(), color.RGBA{255, 0, 0, 255})
							return nil
						},
					),
					types.MarginAll(10),
				),
				craft.NewBox(
					craft.NewImage(ebitenPng),
					types.MarginAll(10),
				).Const(),
			),
		},
	)

	if err := ebiten.RunGame(game.Debug()); err != nil {
		log.Fatal(err)
	}
}
