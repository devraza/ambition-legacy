package main

import (
	// Random
	"math/rand"

	// Logs
	"log"
	
	"strconv"

	// Ebitengine
	"github.com/hajimehoshi/ebiten/v2"
	// "github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Initialise the test player
var testPlayer = initPlayer()

// Create the `Game` struct
type Game struct {
	ui           UI
	activePlayer Player
}

// Define the window width/height
const (
	window_width  = 1440
	window_height = 960
)

// Update implements Game
func (g *Game) Update() error {
	g.ui.base.Update()
	g.activePlayer.update()
	return nil
}

// Draw implements Game
func (g *Game) Draw(screen *ebiten.Image) {
	// Draw the UI onto the screen
	g.ui.base.Draw(screen)
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
	
	testPlayer := initPlayer()

	// Initialise the game
	game := Game{
		// Initialise the UI
		activePlayer: testPlayer,
		ui:           uiInit(window_width, window_height),
	}

	// Log and exit on error
	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
