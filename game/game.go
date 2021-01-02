package game

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/thangtmc73/flappybird/config"
	"github.com/thangtmc73/flappybird/objects"
)

// TickPerFrame represents number of ticks per frame
const TickPerFrame = 1000 / 30

// Game contains everything in game
type Game struct {
	frameStart int64
	running    bool
	camera     *objects.Camera
	bird       *objects.Bird
}

// Init initialized everything in game
func (g *Game) Init() {
	g.running = false
	g.camera = objects.NewCamera(0, config.ScreenHeight)
	g.bird = objects.NewBird()
}

// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	if !g.running {
		g.running = true
		g.frameStart = time.Now().UnixNano()
	}

	g.running = true
	newTime := time.Now().UnixNano()
	deltaTime := (newTime - g.frameStart) / 1000000
	if deltaTime >= TickPerFrame {
		g.frameStart = newTime
		g.updateInGame(deltaTime)
	}
	return nil
}

func (g *Game) updateInGame(deltaTime int64) {
	g.bird.Update(deltaTime)
}

// Draw draws everything in game
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	g.bird.Draw(screen, g.camera)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return config.ScreenWidth, config.ScreenHeight
}
