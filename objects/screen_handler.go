package objects

import "github.com/hajimehoshi/ebiten/v2"

type ScreenHandler interface {
	Init()
	Initialized() bool
	Update(deltaTime int64)
	Draw(screen *ebiten.Image)
	Destroy()
}
