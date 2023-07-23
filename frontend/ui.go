package main

import (
	// Image-related packages
	img "image"
	"image/color"

	// EbitenUI
	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
)

type UI struct {
	base   ebitenui.UI
	colors map[string]color.RGBA
	width, height int
}

// Function for UI initialization
func uiInit(width, height int) UI {
	var ui UI
	// The `hazakura` colorscheme
	{
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

		ui.colors = hazakura
	}

	// Get the window width/height
	ui.width = width
	ui.height = height

	// Define the root container
	root := widget.NewContainer(
		// Define the plain color to be used as a default background
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(ui.colors["dark_black"])),
	)

	// Define the left bar container
	leftBar := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(ui.colors["gray"])),
	)
	// Set the position and size of the left bar
	leftBar.SetLocation(img.Rect(0, 0, int(float64(width)/3.5), height))
	// Add the left bar to the root container
	root.AddChild(leftBar)

	ui.base.Container = root

	return ui
}
