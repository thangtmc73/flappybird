package objects

import (
	"github.com/thangtmc73/flappybird/math"
)

type Camera struct {
	viewportX, viewportY float64
}

func NewCamera(viewport_x, viewport_y float64) *Camera {
	cam := &Camera{
		viewportX: viewport_x,
		viewportY: viewport_y,
	}
	return cam
}

func (c Camera) Transform(x, y float64) (float64, float64) {
	mat := math.Mat4Identity()
	mat.C22 = -1
	mat.C41 = -c.viewportX
	mat.C42 = c.viewportY

	pos := math.Vec4{X: x, Y: y, Z: 1, W: 1}
	newPos := pos.Transform(*mat)
	return newPos.X, newPos.Y
}
