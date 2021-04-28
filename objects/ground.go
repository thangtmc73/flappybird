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
	positionX1, positionY1 float64
	positionX2, positionY2 float64
	displayType          int
	anim                 *animation.Animation
	spriteSheet          *spritesheet.SpriteSheet
}

func NewGround() *Ground {
	ground := &Ground{}
	initialX := float64(config.GroundFrameWidth) / 2
	initialY := float64(config.GroundFrameHeight) / 2
	ground.positionX = initialX
	ground.positionY = initialY
	ground.positionX1 = initialX
	ground.positionY1 = initialY
	ground.positionX2 = initialX + float64(config.GroundFrameWidth)
	ground.positionY2 = initialY
	ground.displayType = rand.Intn(totalTypes)
	ground.anim = animation.New(0, 0, 0, 0)
	ground.spriteSheet =
		spritesheet.New().
			SetNumCols(1).
			SetNumRows(1).
			SetTotalFrames(1).
			SetFrameHeight(config.GroundFrameHeight).
			SetFrameWidth(config.GroundFrameWidth)
	ground.vX = config.ScrollingVX
	return ground
}

func (g *Ground) Update(deltaTime int64) {
	length := float64(config.GroundFrameWidth)
	g.positionX1 += g.vX * float64(deltaTime)
	g.positionY1 += g.vY * float64(deltaTime)
	g.positionX2 += g.vX * float64(deltaTime)
	g.positionY2 += g.vY * float64(deltaTime)
	if g.positionX1 <= float64(-length) / 2 {
		g.positionX1 = g.positionX2 + config.GroundFrameWidth
		return
	}
	if g.positionX2 <= float64(-length) / 2 {
		g.positionX2 = g.positionX1 + config.GroundFrameWidth
	}
}

func (g *Ground) Draw(screen *ebiten.Image, camera *Camera) {
	op1 := &ebiten.DrawImageOptions{}
	op2 := &ebiten.DrawImageOptions{}
	op1.GeoM.Translate(-float64(config.GroundFrameWidth)/2, -float64(config.GroundFrameHeight)/2)
	op1.GeoM.Translate(camera.Transform(g.positionX1, g.positionY1))
	screen.DrawImage(resources.GetSpriteByKey("ground").SubImage(g.spriteSheet.GetRect(*g.anim)).(*ebiten.Image), op1)
	op2.GeoM.Translate(-float64(config.GroundFrameWidth)/2, -float64(config.GroundFrameHeight)/2)
	op2.GeoM.Translate(camera.Transform(g.positionX2, g.positionY2))
	screen.DrawImage(resources.GetSpriteByKey("ground").SubImage(g.spriteSheet.GetRect(*g.anim)).(*ebiten.Image), op2)
}
