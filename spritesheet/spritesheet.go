package spritesheet

import (
	"image"

	"github.com/thangtmc73/flappybird/animation"
)

type SpriteSheet struct {
	numCols, numRows, totalFrames int
	frameHeight, frameWidth       int
}

func New() *SpriteSheet {
	return &SpriteSheet{}
}

func (s *SpriteSheet) SetNumCols(cols int) *SpriteSheet {
	s.numCols = cols
	return s
}

func (s *SpriteSheet) SetNumRows(rows int) *SpriteSheet {
	s.numRows = rows
	return s
}

func (s *SpriteSheet) SetTotalFrames(total int) *SpriteSheet {
	s.totalFrames = total
	return s
}

func (s *SpriteSheet) SetFrameHeight(height int) *SpriteSheet {
	s.frameHeight = height
	return s
}

func (s *SpriteSheet) SetFrameWidth(width int) *SpriteSheet {
	s.frameWidth = width
	return s
}

func (s SpriteSheet) GetRect(anim animation.Animation) image.Rectangle {
	if s.totalFrames == 0 || s.numCols == 0 || s.numRows == 0 {
		return image.Rect(0, 0, 0, 0)
	}
	index := anim.GetIndex()
	sx := (index % s.numCols) * s.frameWidth
	row := (index / s.numCols)
	sy := row * s.frameHeight
	return image.Rect(sx, sy, sx+s.frameWidth, sy+s.frameHeight)
}
