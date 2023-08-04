package ui

import (
	// Image
	img "image"

	// Ambition
	p "github.com/devraza/ambition/frontend/player"

	// EbitenUI
	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
)

func makeLoginWindow(width, height int, root *widget.Container, shown *widget.Container, player *p.Player) {
	// Define the contents of the login window
	loginContainer := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(ui.colors["gray"])),
		// the container will use a row layout to layout the textinput widgets
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(widget.DirectionVertical),
			widget.RowLayoutOpts.Spacing(10),
			widget.RowLayoutOpts.Padding(widget.NewInsetsSimple(20)))),
	)
	loginContainer.AddChild(widget.NewText(
		widget.TextOpts.Text("Login", headingFace, ui.colors["purple"]),
		widget.TextOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.RowLayoutData{
			Position: widget.RowLayoutPositionCenter,
			Stretch:  false,
		})),
	))
	loginContainer.AddChild(widget.NewGraphic(
		widget.GraphicOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.RowLayoutData{
			Stretch:   true,
			MaxHeight: 2,
		})),
		widget.GraphicOpts.ImageNineSlice(image.NewNineSliceColor(ui.colors["gray"])),
	))
	// Create the text inputs
	usernameInput := widget.NewTextInput(
		widget.TextInputOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Position: widget.RowLayoutPositionCenter,
				Stretch:  true,
			}),
		),
		widget.TextInputOpts.Image(&widget.TextInputImage{
			Idle:     image.NewNineSliceColor(ui.colors["black"]),
			Disabled: image.NewNineSliceColor(ui.colors["cyan"]),
		}),
		widget.TextInputOpts.Face(defaultFace),
		widget.TextInputOpts.Color(&widget.TextInputColor{
			Idle:     ui.colors["white"],
			Disabled: ui.colors["light_gray"],
			Caret:    ui.colors["green"],
		}),
		widget.TextInputOpts.Padding(widget.NewInsetsSimple(10)),
		widget.TextInputOpts.CaretOpts(
			widget.CaretOpts.Size(defaultFace, 2),
		),
		widget.TextInputOpts.Placeholder("Username"),
	)
	passwordInput := widget.NewTextInput(
		widget.TextInputOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Position: widget.RowLayoutPositionCenter,
				Stretch:  true,
			}),
		),
		widget.TextInputOpts.Image(&widget.TextInputImage{
			Idle:     image.NewNineSliceColor(ui.colors["black"]),
			Disabled: image.NewNineSliceColor(ui.colors["cyan"]),
		}),
		widget.TextInputOpts.Face(defaultFace),
		widget.TextInputOpts.Color(&widget.TextInputColor{
			Idle:     ui.colors["white"],
			Disabled: ui.colors["light_gray"],
			Caret:    ui.colors["green"],
		}),
		widget.TextInputOpts.Padding(widget.NewInsetsSimple(10)),
		widget.TextInputOpts.CaretOpts(
			widget.CaretOpts.Size(defaultFace, 2),
		),
		// Hide the text inputted
		widget.TextInputOpts.Secure(true),

		widget.TextInputOpts.Placeholder("Password"),
	)
	// Add the text inputs to the login window
	loginContainer.AddChild(usernameInput)
	loginContainer.AddChild(passwordInput)
	loginContainer.AddChild(widget.NewGraphic(
		widget.GraphicOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.RowLayoutData{
			Stretch:   true,
			MaxHeight: 20,
		})),
		widget.GraphicOpts.ImageNineSlice(image.NewNineSliceColor(ui.colors["gray"])),
	))

	// Define the 'confirm' button
	confirmButton := widget.NewButton(
		// Button options
		widget.ButtonOpts.WidgetOpts(
			// Center the button both horizontally and vertically
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Position: widget.RowLayoutPositionCenter,
				Stretch:  false,
			}),
		),
		// More options
		widget.ButtonOpts.Image(buttonImage),
		widget.ButtonOpts.Text("Confirm", defaultFace, &widget.ButtonTextColor{
			Idle: ui.colors["white"],
		}),
		// Define the padding on the button
		widget.ButtonOpts.TextPadding(widget.Insets{
			Left:   20,
			Right:  20,
			Top:    10,
			Bottom: 10,
		}),
		// Button on-click handler
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			player.Init(usernameInput.InputText, passwordInput.InputText)
			removeContainer(root, loginContainer)
			addContainer(root, shown)
		}),
	)
	loginContainer.AddChild(confirmButton)

	// Place and show the login container
	placeContainer(loginContainer, ui.width/3, ui.height/4, img.Point{ui.width / 3, ui.height / 3})
	addContainer(root, loginContainer)
}
