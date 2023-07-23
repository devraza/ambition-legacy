package main

import (
	// Random
	"math/rand"

	// Logs
	"log"
	
	"strconv"

	// Ebitengine
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Create the `Game` struct
type Game struct {
	ui UI
	activePlayer Player
}

const (
	window_width = 640
	window_height = 480
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
	
	ebitenutil.DebugPrintAt(screen, "health: " + strconv.Itoa(g.activePlayer.health), 0, 0)
	ebitenutil.DebugPrintAt(screen, "level: " + strconv.Itoa(g.activePlayer.level), 0, 10)
	ebitenutil.DebugPrintAt(screen, "exp: " + strconv.Itoa(int(g.activePlayer.exp)), 0, 20)
	ebitenutil.DebugPrintAt(screen, "ambittion: " + strconv.Itoa(int(g.activePlayer.ambition)), 0, 30)
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
		ui: uiInit(window_width, window_height),
		activePlayer: testPlayer,
	}

	// Log and exit on error
	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
