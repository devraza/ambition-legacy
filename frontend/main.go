package main

import (
	// Misc.
	"log"
	"math/rand"

	// Ambition
	p "github.com/devraza/ambition/frontend/player"
	u "github.com/devraza/ambition/frontend/ui"

	// Ebitengine
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Create the `Game` struct
type Game struct {
	ui           u.UI
	activePlayer p.Player
}

// Define the window width/height
const (
	window_width  = 1440
	window_height = 960
)

// Update implements Game
func (g *Game) Update() error {
	g.ui.Base.Update()
	g.activePlayer.Update()
	return nil
}

// Draw implements Game
func (g *Game) Draw(screen *ebiten.Image) {
	// Draw the UI onto the screen
	g.ui.Base.Draw(screen)
	ebitenutil.DebugPrint(screen, g.activePlayer.JwtToken)
}

// Layout implements Game
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

// Main function
func main() {
	// Randomize window titles!
	window_titles := []string{
		"coding into the online cyberframe",
		"seriously prideful",
		"[REDACTED]",
		"just another day of shooting down bevies",
		"mud is delicious",
	}
	window_title := "Ambition: " + window_titles[rand.Intn(len(window_titles))]

	// Engine setup
	ebiten.SetWindowSize(window_width, window_height)
	ebiten.SetWindowTitle(window_title)

	// Initialise the game
	game := Game{}
	game.activePlayer = p.NewPlayer()
	game.ui = u.UiInit(window_width, window_height, &game.activePlayer)

	// Log and exit on error
	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
