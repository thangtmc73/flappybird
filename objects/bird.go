package objects

import (
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/thangtmc73/flappybird/animation"
	"github.com/thangtmc73/flappybird/config"
	"github.com/thangtmc73/flappybird/resources"
	"github.com/thangtmc73/flappybird/spritesheet"
)

const (
	yellow     = 0
	blue       = 1
	red        = 2
	totalTypes = 3
)

// Bird describes bird in game
type Bird struct {
	positionX, positionY float64
	displayType          int
	anim                 *animation.Animation
	spriteSheet          *spritesheet.SpriteSheet
}

func NewBird() *Bird {
	bird := &Bird{}
	bird.displayType = rand.Intn(totalTypes)
	bird.anim = animation.New(bird.displayType*3, bird.displayType*3+2, bird.displayType*3, 200)
	bird.spriteSheet =
		spritesheet.New().
			SetNumCols(3).
			SetNumRows(3).
			SetTotalFrames(9).
			SetFrameHeight(config.BirdFrameHeight).
			SetFrameWidth(config.BirdFrameWidth)
	return bird
}

func (b *Bird) Update(deltaTime int64) {
	b.anim.Update(deltaTime)
}

func (b *Bird) Draw(screen *ebiten.Image, camera *Camera) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(config.BirdFrameWidth)/2, -float64(config.BirdFrameHeight)/2)
	op.GeoM.Translate(camera.Transform(b.positionX, b.positionY))
	screen.DrawImage(resources.GetSpriteByKey("bird").SubImage(b.spriteSheet.GetRect(*b.anim)).(*ebiten.Image), op)
}
