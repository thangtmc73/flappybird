package screens

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/thangtmc73/flappybird/config"
	"github.com/thangtmc73/flappybird/objects"
	"math"
)

type PlayScreen struct {
	initialized bool
	camera     *objects.Camera
	bird       *objects.Bird
	background *objects.Background
	ground     *objects.Ground
	birdReadyTime int64
	startPlay  bool
}

const InitialBirdPositionY = 120

func NewPlayScreen() *PlayScreen {
	screen := &PlayScreen{}
	return screen
}

func (playScreen *PlayScreen) Init() {
	playScreen.camera = objects.NewCamera(0, config.ScreenHeight)
	playScreen.bird = objects.NewBird(72, InitialBirdPositionY)
	playScreen.background = objects.NewBackground()
	playScreen.ground = objects.NewGround()
	playScreen.initialized = true
	playScreen.birdReadyTime = 0
	playScreen.startPlay = false
}

func (playScreen * PlayScreen) Initialized() bool {
	return playScreen.initialized
}

func (playScreen *PlayScreen) Update(deltaTime int64) {
	if !playScreen.startPlay {
		playScreen.birdReadyTime += deltaTime
		playScreen.bird.SetPositionY(InitialBirdPositionY + 3 * math.Sin(math.Pi * 2 / 1200 * float64(playScreen.birdReadyTime)))
	}
	playScreen.ground.Update(deltaTime)
	playScreen.bird.Update(deltaTime)
}

func (playScreen *PlayScreen) Draw(screen *ebiten.Image) {
	playScreen.background.Draw(screen, playScreen.camera)
	playScreen.ground.Draw(screen, playScreen.camera)
	playScreen.bird.Draw(screen, playScreen.camera)
}

func (playScreen * PlayScreen) Destroy() {

}