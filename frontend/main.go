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

// Create the `Game` struct
type Game struct {
	ui ebitenui.UI
}

func uiInit() ebitenui.UI {
	// The `hazakura` colorscheme
	hazakura := make(map[string]color.RGBA)
	// The monotone colors
	hazakura["dark_black"] = color.RGBA{0x0f, 0x0f, 0x0d, 0xff}
	hazakura["black"] = color.RGBA{0x15, 0x15, 0x17, 0xff}
	hazakura["dark_gray"] = color.RGBA{0x24, 0x24, 0x26, 0xff}
	hazakura["gray"] = color.RGBA{0x27, 0x27, 0x2b, 0xff}
	hazakura["light_gray"] = color.RGBA{0x45, 0x44, 0x49, 0xff}
	hazakura["overlay"] = color.RGBA{0x5c, 0x5c, 0x61, 0xff}
	hazakura["highlight"] = color.RGBA{0xd9, 0x0, 0xd7, 0xff}
	hazakura["subwhite"] = color.RGBA{0xec, 0xe5, 0xea, 0xff}

	// Actual* colors
	hazakura["red"] = color.RGBA{0xf0, 0x69, 0x69, 0xff}
	hazakura["magenta"] = color.RGBA{0xe8, 0x87, 0xbb, 0xff}
	hazakura["purple"] = color.RGBA{0xa2, 0x92, 0xe8, 0xff}
	hazakura["blue"] = color.RGBA{0x78, 0xb9, 0xc4, 0xff}
	hazakura["cyan"] = color.RGBA{0x7e, 0xe6, 0xae, 0xff}
	hazakura["green"] = color.RGBA{0x91, 0xd6, 0x5c, 0xff}
	hazakura["yellow"] = color.RGBA{0xd9, 0xd5, 0x64, 0xff}

	// Get the window width/height
	window_width, window_height := ebiten.WindowSize()

	// Define the root container
	root := widget.NewContainer(
		// Define the plain color to be used as a default background
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(hazakura["dark_black"])),
	)

	// Define the left bar container
	leftBar := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(hazakura["gray"])),
	)
	// Set the position and size of the left bar
	leftBar.SetLocation(img.Rect(0, 0, int(float64(window_width)/3.5), window_height))
	// Add the left bar to the root container
	root.AddChild(leftBar)

	// Construct the UI
	ui := ebitenui.UI{
		Container: root,
	}

	return ui
}

// Update implements Game
func (g *Game) Update() error {
	g.ui.Update()
	return nil
}

// Draw implements Game
func (g *Game) Draw(screen *ebiten.Image) {
	g.ui.Draw(screen)
}

// Layout implements Game
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
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
