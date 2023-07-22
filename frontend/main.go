package main

import (
	"log"

	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	img "image"
	"image/color"
)

type Game struct{
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
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("ambition: seriously prideful")
	game := Game{
		ui: uiInit(),
	}
	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
