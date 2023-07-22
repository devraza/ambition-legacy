package main

import (
	// Image-related packages
	img "image"
	"image/color"

	// Random
	"math/rand"

	// Logs
	"log"

	// Ebitengine
	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	ui ebitenui.UI
}

func uiInit() ebitenui.UI {
	ui := ebitenui.UI{
		Container: widget.NewContainer(widget.ContainerOpts.Layout(widget.NewAnchorLayout())),
	}
	leftBar := widget.NewContainer(widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(color.White)))
	leftBar.SetLocation(img.Rect(0, 0, 1, 1))
	ui.Container.AddChild(leftBar)

	return ui
}

func (g *Game) Update() error {
	g.ui.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.ui.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	window_titles := []string{
		"coding into the online cyberframe",
		"seriously prideful",
		"[REDACTED]",
		"just another day of shooting down bevies",
		"mud is delicious",
	}
	window_title := "Ambition: " + window_titles[rand.Intn(len(window_titles))]

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle(window_title)
	game := Game{
		ui: uiInit(),
	}
	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
