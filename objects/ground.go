package objects

import (
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/thangtmc73/flappybird/animation"
	"github.com/thangtmc73/flappybird/config"
	"github.com/thangtmc73/flappybird/resources"
	"github.com/thangtmc73/flappybird/spritesheet"
)

// Ground describes bird in game
type Ground struct {
	positionX, positionY float64
	vX                   float64
	vY                   float64
	displayType          int
	anim                 *animation.Animation
	spriteSheet          *spritesheet.SpriteSheet
}

func NewGround() *Ground {
	ground := &Ground{}
	ground.positionX = 84
	ground.positionY = 28
	ground.displayType = rand.Intn(totalTypes)
	ground.anim = animation.New(0, 0, 0, 0)
	ground.spriteSheet =
		spritesheet.New().
			SetNumCols(1).
			SetNumRows(1).
			SetTotalFrames(1).
			SetFrameHeight(config.GroundFrameHeight).
			SetFrameWidth(config.GroundFrameWidth)
	return ground
}

func (g *Ground) Update(deltaTime int64) {
}

func (g *Ground) Draw(screen *ebiten.Image, camera *Camera) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(config.GroundFrameWidth)/2, -float64(config.GroundFrameHeight)/2)
	op.GeoM.Translate(camera.Transform(g.positionX, g.positionY))
	screen.DrawImage(resources.GetSpriteByKey("ground").SubImage(g.spriteSheet.GetRect(*g.anim)).(*ebiten.Image), op)
}
