package etk

import (
	"errors"
	"fmt"
	"image/color"

	"github.com/a-skua/etk/widget"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Scene interface {
	ebiten.Game
	Init()
	Next() bool
	Previous() bool
}

type DefaultScene struct {
	Widget widget.Widget
}

func (DefaultScene) Init() {}

func (DefaultScene) Next() bool {
	return inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft)
}

func (DefaultScene) Previous() bool {
	return false
}

func (DefaultScene) Update() error {
	return nil
}

func (s DefaultScene) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)
	if s.Widget != nil {
		screen.DrawImage(s.Widget.Image(), nil)
	}
}

func (DefaultScene) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

type scene struct {
	current int
	list    []Scene
}

func (s *scene) next() {
	s.current++
	if s.current >= len(s.list) {
		s.current = 0
	}
}

func (s *scene) prev() {
	s.current--
	if s.current < 0 {
		s.current = len(s.list) - 1
	}
}

func (s *scene) Current() Scene {
	return s.list[s.current]
}

type Game struct {
	scene   scene
	width   int
	height  int
	options []interface {
		Update() error
		Draw(screen *ebiten.Image)
	}
}

func New(width, height int, scene Scene, scenes ...Scene) *Game {
	g := &Game{}
	g.width = width
	g.height = height
	g.scene.list = make([]Scene, 0, len(scenes)+1)
	g.scene.list = append(g.scene.list, scene)
	g.scene.list = append(g.scene.list, scenes...)
	g.scene.Current().Init()
	return g
}

func (g *Game) Update() error {
	if g.scene.Current().Next() {
		g.scene.next()
		g.scene.Current().Init()
	} else if g.scene.Current().Previous() {
		g.scene.prev()
		g.scene.Current().Init()
	}

	err := g.scene.Current().Update()
	for _, option := range g.options {
		err = errors.Join(option.Update())
	}
	return err
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.scene.Current().Draw(screen)
	for _, option := range g.options {
		option.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.scene.Current().Layout(g.width, g.height)
}

type debug struct{}

func (d debug) Update() error {
	return nil
}

func (d debug) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, fmt.Sprintf(
		"FPS: %0.2f\nTPS: %0.2f",
		ebiten.ActualFPS(),
		ebiten.ActualTPS(),
	))
}

func (g *Game) Debug() *Game {
	g.options = append(g.options, debug{})
	return g
}
