package ui

import (
	// Image-related packages
	img "image"
	"image/color"

	// Ambition
	p "github.com/devraza/ambition/frontend/player"

	// EbitenUI
	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"

	// Fonts
	"github.com/devraza/ambition/frontend/assets/fonts"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

// The UI struct
type UI struct {
	Base          ebitenui.UI
	player        *p.Player
	colors        map[string]color.RGBA
	width, height int
	textInput     *widget.TextInput
}

// The `hazakura` colorscheme (default)
var hazakura = map[string]color.RGBA{
	// The monotone colors
	"dark_black": color.RGBA{0x0d, 0x0d, 0x0f, 0xff},
	"black":      color.RGBA{0x15, 0x15, 0x17, 0xff},
	"dark_gray":  color.RGBA{0x24, 0x24, 0x26, 0xff},
	"gray":       color.RGBA{0x27, 0x27, 0x2b, 0xff},
	"light_gray": color.RGBA{0x45, 0x44, 0x49, 0xff},
	"overlay":    color.RGBA{0x5c, 0x5c, 0x61, 0xff},
	"subwhite":   color.RGBA{0xd9, 0x0, 0xd7, 0xff},
	"white":      color.RGBA{0xec, 0xe5, 0xea, 0xff},

	// Actual* colors
	"red":     color.RGBA{0xf0, 0x69, 0x69, 0xff},
	"magenta": color.RGBA{0xe8, 0x87, 0xbb, 0xff},
	"purple":  color.RGBA{0xa2, 0x92, 0xe8, 0xff},
	"blue":    color.RGBA{0x78, 0xb9, 0xc4, 0xff},
	"cyan":    color.RGBA{0x7e, 0xe6, 0xae, 0xff},
	"green":   color.RGBA{0x91, 0xd6, 0x5c, 0xff},
	"yellow":  color.RGBA{0xd9, 0xd5, 0x64, 0xff},
}

var ui UI

// Make the different faces
var headingFace, _ = makeFace(18, fonts.FiraBold_ttf)
var defaultFace, _ = makeFace(14, fonts.FiraRegular_ttf)

// Load the images for the button states
var buttonImage, _ = loadButtonImage()

// Function for UI initialization
func UiInit(width, height int, player *p.Player) UI {
	// Define the UI colors
	ui.colors = hazakura
	ui.player = player

	// Get the window width/height
	ui.width = width
	ui.height = height

	// Define the root container
	root := widget.NewContainer(
		// Define the plain color to be used as a default background
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(ui.colors["dark_black"])),
	)

	// Create the 'Profile' tab
	tabProfile := widget.NewTabBookTab("Profile",
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(ui.colors["gray"])),
		widget.ContainerOpts.Layout(widget.NewGridLayout(
			//Define number of columns in the grid
			widget.GridLayoutOpts.Columns(2),
			// Specify the Stretch for each row and column.
			widget.GridLayoutOpts.Stretch([]bool{false, true}, nil),
			// Define the spacing between items in the grid
			widget.GridLayoutOpts.Spacing(20, 15),
			// Define the padding in the grid
			widget.GridLayoutOpts.Padding(widget.NewInsetsSimple(20)),
		)),
	)
	// Add the player stats as content for the 'profile' tab
	makeStatsBars(tabProfile, ui, defaultFace)

	// Create the 'Inventory' tab
	tabInventory := widget.NewTabBookTab("Inventory",
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(ui.colors["gray"])),
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
	)
	inventoryButton := widget.NewText(
		widget.TextOpts.Text("Placeholder", headingFace, ui.colors["white"]),
		widget.TextOpts.Position(widget.TextPositionCenter, widget.TextPositionCenter),
		widget.TextOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
			HorizontalPosition: widget.AnchorLayoutPositionCenter,
			VerticalPosition:   widget.AnchorLayoutPositionCenter,
		})),
	)
	tabInventory.AddChild(inventoryButton)

	// Create the 'Other' tab
	tabOther := widget.NewTabBookTab("Other",
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(ui.colors["gray"])),
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
	)
	otherButton := widget.NewText(
		widget.TextOpts.Text("Placeholder", headingFace, ui.colors["white"]),
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
		widget.TabBookOpts.TabButtonText(headingFace, &widget.ButtonTextColor{Idle: color.White}),
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
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
	)
	// Add the tabbook to the left bar
	leftBar.AddChild(leftTabs)

	// Define the contents of the chat window
	chatContainer := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(ui.colors["black"])),
		widget.ContainerOpts.Layout(widget.NewAnchorLayout(
			widget.AnchorLayoutOpts.Padding(widget.NewInsetsSimple(20)),
		)),
	)
	chatContainer.AddChild(widget.NewText(
		widget.TextOpts.Text("Chat", headingFace, ui.colors["blue"]),
		widget.TextOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
			HorizontalPosition: widget.AnchorLayoutPositionStart,
			VerticalPosition:   widget.AnchorLayoutPositionStart,
		})),
	))
	// Place and show the chat container
	placeContainer(chatContainer, int(float32(ui.width)/3.5), int(float32(ui.height)/3.3), img.Point{0, ui.height - int(float32(ui.height)/3.3)})
	addContainer(leftBar, chatContainer)

	// Create the login window
	makeLoginWindow(ui.width, ui.height, root, leftBar, player)

	// Set the position and size of the left bar
	leftBar.SetLocation(img.Rect(0, 0, int(float32(width)/3.5), height))
	// Add the left bar to the root container
	//root.AddChild(leftBar)

	// Set the base container to be the root container
	ui.Base.Container = root

	return ui
}

// Set the size/location of a container
func placeContainer(container *widget.Container, x int, y int, vector img.Point) {
	DimensionX, DimensionY := x, y
	placement := img.Rect(0, 0, DimensionX, DimensionY)
	placement = placement.Add(vector)
	// Set the position and size of the container
	container.SetLocation(placement)
}

// Add a container to a parent container
func addContainer(parent *widget.Container, child *widget.Container) {
	parent.AddChild(child)
}

// Hide a container
func removeContainer(parent *widget.Container, child *widget.Container) {
	parent.RemoveChild(child)
}

// Set a window's location and open the window
func showWindow(window *widget.Window, v float32, h float32) {
	// Get the preferred size of the content
	x, y := window.Contents.PreferredSize()
	// Create a rect with the preferred size of the content
	r := img.Rect(0, 0, x, y)
	// Use the Add method to move the window to the specified point
	r = r.Add(img.Point{int(v), int(h)})
	// Set the windows location to the rect
	window.SetLocation(r)
}

// Create progressbars for all the player stats
func makeStatsBars(parent *widget.TabBookTab, ui UI, face font.Face) {
	// Health
	health := widget.NewText(
		widget.TextOpts.Text("Health", face, ui.colors["white"]),
		widget.TextOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.GridLayoutData{
			HorizontalPosition: widget.GridLayoutPositionStart,
		})),
	)
	health_progressbar := widget.NewProgressBar(
		widget.ProgressBarOpts.WidgetOpts(
			// Set the required anchor layout data to determine where in the container to place the progressbar
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionCenter,
				VerticalPosition:   widget.AnchorLayoutPositionCenter,
			}),
			// Set the minimum size for the progressbar.
			widget.WidgetOpts.MinSize(200, 20),
		),
		widget.ProgressBarOpts.Images(
			// Set the track colors
			&widget.ProgressBarImage{
				Idle:  image.NewNineSliceColor(ui.colors["black"]),
				Hover: image.NewNineSliceColor(ui.colors["black"]),
			},
			// Set the progress colors
			&widget.ProgressBarImage{
				Idle:  image.NewNineSliceColor(ui.colors["red"]),
				Hover: image.NewNineSliceColor(ui.colors["red"]),
			},
		),
		// Set the min, max, and current values for each progressbar
		widget.ProgressBarOpts.Values(0, ui.player.Health, ui.player.MaxHealth),
	)
	parent.AddChild(health)
	parent.AddChild(health_progressbar)

	// XP/Level
	level := widget.NewText(
		widget.TextOpts.Text("Level", face, ui.colors["white"]),
		widget.TextOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.GridLayoutData{
			HorizontalPosition: widget.GridLayoutPositionStart,
		})),
	)
	level_progressbar := widget.NewProgressBar(
		widget.ProgressBarOpts.WidgetOpts(
			// Set the required anchor layout data to determine where in the container to place the progressbar
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionCenter,
				VerticalPosition:   widget.AnchorLayoutPositionCenter,
			}),
			// Set the minimum size for the progressbar.
			widget.WidgetOpts.MinSize(200, 20),
		),
		widget.ProgressBarOpts.Images(
			// Set the track colors
			&widget.ProgressBarImage{
				Idle:  image.NewNineSliceColor(ui.colors["black"]),
				Hover: image.NewNineSliceColor(ui.colors["black"]),
			},
			// Set the progress colors
			&widget.ProgressBarImage{
				Idle:  image.NewNineSliceColor(ui.colors["cyan"]),
				Hover: image.NewNineSliceColor(ui.colors["cyan"]),
			},
		),
		// Set the min, max, and current values for each progressbar
		widget.ProgressBarOpts.Values(0, int(ui.player.Exp), int(ui.player.NextExp)),
	)
	parent.AddChild(level)
	parent.AddChild(level_progressbar)

	// Ambition
	ambition := widget.NewText(
		widget.TextOpts.Text("Ambition", face, ui.colors["white"]),
		widget.TextOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.GridLayoutData{
			HorizontalPosition: widget.GridLayoutPositionStart,
		})),
	)
	ambition_progressbar := widget.NewProgressBar(
		widget.ProgressBarOpts.WidgetOpts(
			// Set the required anchor layout data to determine where in the container to place the progressbar
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionCenter,
				VerticalPosition:   widget.AnchorLayoutPositionCenter,
			}),
			// Set the minimum size for the progressbar.
			widget.WidgetOpts.MinSize(200, 20),
		),
		widget.ProgressBarOpts.Images(
			// Set the track colors
			&widget.ProgressBarImage{
				Idle:  image.NewNineSliceColor(ui.colors["black"]),
				Hover: image.NewNineSliceColor(ui.colors["black"]),
			},
			// Set the progress colors
			&widget.ProgressBarImage{
				Idle:  image.NewNineSliceColor(ui.colors["purple"]),
				Hover: image.NewNineSliceColor(ui.colors["purple"]),
			},
		),
		// Set the min, max, and current values for each progressbar
		widget.ProgressBarOpts.Values(0, int(ui.player.Ambition), int(ui.player.MaxAmbition)),
	)
	parent.AddChild(ambition)
	parent.AddChild(ambition_progressbar)
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
