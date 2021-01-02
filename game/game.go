package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/thangtmc73/flappybird/config"
)

// Game contains everything in game
type Game struct {
}

// Init initialized everything in game
func (g *Game) Init() {
}

// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	return nil
}

// Draw draws everything in game
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return config.SCREEN_WIDTH, config.SCREEN_HEIGHT
}
