package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/thangtmc73/flappybird/config"
	"github.com/thangtmc73/flappybird/managers"
	"time"
)

type Mode int
const (
	Start Mode	= 0
	Play Mode	= 1
)

// TickPerFrame represents number of ticks per frame
const TickPerFrame = 1000 / 30

// Game contains everything in game
type Game struct {
	frameStart int64
	running    bool
	screenManager *managers.ScreenManager
	mode Mode
}

// Init initialized everything in game
func (g *Game) Init() {
	g.running = false
	g.screenManager = managers.NewScreenManager()
	g.mode = Start
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
	currentScreen, _ := g.screenManager.Back()
	if currentScreen == nil {
		return
	}
	(*currentScreen).Update(deltaTime)
}

// Draw draws everything in game
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	currentScreen, _ := g.screenManager.Back()
	if currentScreen == nil {
		return
	}
	(*currentScreen).Draw(screen)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return config.ScreenWidth, config.ScreenHeight
}
