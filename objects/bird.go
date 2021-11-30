package objects

import (
	"math"
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

const RewindRotateDegree = -30
const FallDownRotateDegree = 90
const RewindDuration = 850 // ms

// Bird describes bird in game
type Bird struct {
	positionX, positionY float64
	displayType          int
	anim                 *animation.Animation
	spriteSheet          *spritesheet.SpriteSheet
	rotateDegree float64
	autoFall bool
	rewindTime int64
}

func NewBird(positionX, positionY float64) *Bird {
	bird := &Bird{}
	bird.positionX = positionX
	bird.positionY = positionY
	bird.displayType = rand.Intn(totalTypes)
	bird.anim = animation.New(bird.displayType*3, bird.displayType*3+2, bird.displayType*3, 200)
	bird.spriteSheet =
		spritesheet.New().
			SetNumCols(3).
			SetNumRows(3).
			SetTotalFrames(9).
			SetFrameHeight(config.BirdFrameHeight).
			SetFrameWidth(config.BirdFrameWidth)
	bird.rotateDegree = 0
	bird.autoFall = false
	return bird
}

func (b *Bird) Update(deltaTime int64) {
	if b.autoFall && b.rotateDegree < FallDownRotateDegree {
		b.rewindTime += deltaTime
		b.rotateDegree = float64(RewindRotateDegree) + float64(FallDownRotateDegree - RewindRotateDegree) * float64(b.rewindTime) / RewindDuration
		if b.rotateDegree > FallDownRotateDegree {
			b.rotateDegree = FallDownRotateDegree
			b.rewindTime = 0
		}
	}
	b.anim.Update(deltaTime)
}

func (b* Bird) GetPosition() (float64, float64) {
	return b.positionX, b.positionY
}

func (b * Bird) SetPositionX(positionX float64) *Bird {
	b.positionX = positionX
	return b
}

func (b * Bird) SetPositionY(positionY float64) *Bird {
	b.positionY = positionY
	return b
}

func (b *Bird) Fly() {
	b.rotateDegree = RewindRotateDegree
	b.rewindTime = 0
	b.autoFall = true
	b.positionY += 5
}

func (b *Bird) Draw(screen *ebiten.Image, camera *Camera) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(config.BirdFrameWidth)/2, -float64(config.BirdFrameHeight)/2)
	op.GeoM.Rotate(float64(int(b.rotateDegree)%360) * 2 * math.Pi / 360)
	op.GeoM.Translate(camera.Transform(b.positionX, b.positionY))
	screen.DrawImage(resources.GetSpriteByKey("bird").SubImage(b.spriteSheet.GetRect(*b.anim)).(*ebiten.Image), op)
}
