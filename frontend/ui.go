package main

import (
	// Image-related packages
	img "image"
	"image/color"

	// String convert
	// "strconv"

	// EbitenUI
	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"

	// Fonts
	"github.com/devraza/ambition/assets/fonts"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

// The UI struct
type UI struct {
	base          ebitenui.UI
	colors        map[string]color.RGBA
	width, height int
}

// The `hazakura` colorscheme (default)
var hazakura = map[string]color.RGBA{
	// The monotone colors
	"dark_black": color.RGBA{0x0f, 0x0f, 0x0d, 0xff},
	"black":      color.RGBA{0x15, 0x15, 0x17, 0xff},
	"dark_gray":  color.RGBA{0x24, 0x24, 0x26, 0xff},
	"gray":       color.RGBA{0x27, 0x27, 0x2b, 0xff},
	"light_gray": color.RGBA{0x45, 0x44, 0x49, 0xff},
	"overlay":    color.RGBA{0x5c, 0x5c, 0x61, 0xff},
	"highlight":  color.RGBA{0xd9, 0x0, 0xd7, 0xff},
	"subwhite":   color.RGBA{0xec, 0xe5, 0xea, 0xff},

	// Actual* colors
	"red":     color.RGBA{0xf0, 0x69, 0x69, 0xff},
	"magenta": color.RGBA{0xe8, 0x87, 0xbb, 0xff},
	"purple":  color.RGBA{0xa2, 0x92, 0xe8, 0xff},
	"blue":    color.RGBA{0x78, 0xb9, 0xc4, 0xff},
	"cyan":    color.RGBA{0x7e, 0xe6, 0xae, 0xff},
	"green":   color.RGBA{0x91, 0xd6, 0x5c, 0xff},
	"yellow":  color.RGBA{0xd9, 0xd5, 0x64, 0xff},
}

// Function for UI initialization
func uiInit(width, height int) UI {
	var ui UI

	// Define the UI colors
	ui.colors = hazakura

	// Load the images for the button states
	buttonImage, _ := loadButtonImage()

	// Get the window width/height
	ui.width = width
	ui.height = height

	// Define the root container
	root := widget.NewContainer(
		// Define the plain color to be used as a default background
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(ui.colors["dark_black"])),
	)

	// Set the heading size and font
	heading_face, _ := makeFace(18, fonts.IosevkaBold_ttf)

	tabProfile := widget.NewTabBookTab("Profile",
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(ui.colors["gray"])),
		widget.ContainerOpts.Layout(widget.NewGridLayout(
			//Define number of columns in the grid
			widget.GridLayoutOpts.Columns(2),
			widget.GridLayoutOpts.Spacing(10, 10),
			//specify the Stretch for each row and column.
			widget.GridLayoutOpts.Stretch([]bool{false, true}, nil),
		)),
	)
	// TODO(devraza): Show all of the player's stats
	// health := widget.NewButton(
	// 	widget.ButtonOpts.Text(strconv.Itoa(activePlayer.health), face, buttonTextColor),
	// )
	// tabProfile.AddChild(health)

	tabInventory := widget.NewTabBookTab("Inventory",
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(color.NRGBA{0, 255, 0, 0xff})),
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
	)
	inventoryButton := widget.NewText(
		widget.TextOpts.Text("Inventory content", heading_face, color.Black),
		widget.TextOpts.Position(widget.TextPositionCenter, widget.TextPositionCenter),
		widget.TextOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
			HorizontalPosition: widget.AnchorLayoutPositionCenter,
			VerticalPosition:   widget.AnchorLayoutPositionCenter,
		})),
	)
	tabInventory.AddChild(inventoryButton)

	tabOther := widget.NewTabBookTab("Other",
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(color.NRGBA{0, 255, 0, 0xff})),
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
	)
	otherButton := widget.NewText(
		widget.TextOpts.Text("Other content", heading_face, color.Black),
		widget.TextOpts.Position(widget.TextPositionCenter, widget.TextPositionCenter),
		widget.TextOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
			HorizontalPosition: widget.AnchorLayoutPositionCenter,
			VerticalPosition:   widget.AnchorLayoutPositionCenter,
		})),
	)
	tabOther.AddChild(otherButton)

	// Create the tabbook widget
	leftTabs := widget.NewTabBook(
		widget.TabBookOpts.TabButtonImage(buttonImage),
		widget.TabBookOpts.TabButtonText(heading_face, &widget.ButtonTextColor{Idle: color.White}),
		widget.TabBookOpts.TabButtonSpacing(0),
		widget.TabBookOpts.ContainerOpts(
			widget.ContainerOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				StretchHorizontal:  true,
				StretchVertical:    true,
				HorizontalPosition: widget.AnchorLayoutPositionCenter,
				VerticalPosition:   widget.AnchorLayoutPositionCenter,
			}),
			),
		),
		widget.TabBookOpts.TabButtonOpts(
			widget.ButtonOpts.TextPadding(widget.NewInsetsSimple(5)),
			widget.ButtonOpts.WidgetOpts(widget.WidgetOpts.MinSize(int(float32(width)/(3.5*3)), 0)),
		),
		widget.TabBookOpts.Tabs(tabProfile, tabInventory, tabOther),
	)
	// Define the left bar container
	leftBar := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(ui.colors["gray"])),
	)
	// Add the tabbook to the left bar
	leftBar.AddChild(leftTabs)

	// Set the position and size of the left bar
	leftBar.SetLocation(img.Rect(0, 0, int(float32(width)/3.5), height))
	// Add the left bar to the root container
	root.AddChild(leftBar)

	ui.base.Container = root

	return ui
}

// Load a button image
func loadButtonImage() (*widget.ButtonImage, error) {
	idle := image.NewNineSliceColor(hazakura["black"])
	hover := image.NewNineSliceColor(hazakura["light_gray"])
	pressed := image.NewNineSliceColor(hazakura["gray"])
	pressedHover := image.NewNineSliceColor(hazakura["gray"])
	disabled := image.NewNineSliceColor(hazakura["overlay"])

	return &widget.ButtonImage{
		Idle:         idle,
		Hover:        hover,
		Pressed:      pressed,
		PressedHover: pressedHover,
		Disabled:     disabled,
	}, nil
}

// Function to create a face providing font size and file (from assets)
func makeFace(size float64, fontfile []byte) (font.Face, error) {
	ttfFont, err := truetype.Parse(fontfile)
	if err != nil {
		return nil, err
	}

	return truetype.NewFace(ttfFont, &truetype.Options{
		Size:    size,
		DPI:     72,
		Hinting: font.HintingFull,
	}), nil
}
