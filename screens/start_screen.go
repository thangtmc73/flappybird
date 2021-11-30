package screens

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/thangtmc73/flappybird/config"
	"github.com/thangtmc73/flappybird/objects"
)

type StartScreen struct {
	initialized bool
	camera     *objects.Camera
	bird       *objects.Bird
	background *objects.Background
	ground     *objects.Ground
}

func NewStartScreen() *StartScreen {
	screen := &StartScreen{}
	return screen
}

func (startScreen *StartScreen) Init() {
	startScreen.camera = objects.NewCamera(0, config.ScreenHeight)
	startScreen.bird = objects.NewBird(72, 120)
	startScreen.background = objects.NewBackground()
	startScreen.ground = objects.NewGround()
	startScreen.initialized = true
}

func (startScreen * StartScreen) Initialized() bool {
	return startScreen.initialized
}

func (startScreen *StartScreen) Update(deltaTime int64) {
	startScreen.ground.Update(deltaTime)
	startScreen.bird.Update(deltaTime)
}

func (startScreen *StartScreen) Draw(screen *ebiten.Image) {
	startScreen.background.Draw(screen, startScreen.camera)
	startScreen.ground.Draw(screen, startScreen.camera)
	startScreen.bird.Draw(screen, startScreen.camera)
}

func (startScreen * StartScreen) Destroy() {

}
