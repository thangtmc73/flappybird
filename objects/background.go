package objects

import (
	"github.com/thangtmc73/flappybird/config"
	"image"

	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/thangtmc73/flappybird/resources"
)

const (
	day   = 0
	night = 1
	total = 2
)

type Background struct {
	positionX, positionY float64
	displayType int
}

// NewBackground initializes background
func NewBackground() *Background {
	background := &Background{}
	background.positionX = config.BackgroundWidth / 2
	background.positionY = config.BackgroundHeight / 2
	background.displayType = rand.Intn(total)
	return background
}

// Update updates background
func (bg *Background) Update() error {
	return nil
}

// Draw draws background
func (bg *Background) Draw(screen *ebiten.Image, camera *Camera) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(config.BackgroundWidth)/2, -float64(config.BackgroundHeight)/2)
	op.GeoM.Translate(camera.Transform(bg.positionX, bg.positionY))
	screen.DrawImage(resources.GetSpriteByKey("background").SubImage(getBackgroundImageRect(bg.displayType)).(*ebiten.Image), op)
}

func getBackgroundImageRect(displayType int) image.Rectangle {
	switch displayType {
	case day:
		return image.Rect(0, 0, 144, 256)
	case night:
		return image.Rect(144, 0, 288, 256)
	default:
		return image.Rect(0, 0, 144, 256)
	}
}
