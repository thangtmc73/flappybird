package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/thangtmc73/flappybird/config"
	"github.com/thangtmc73/flappybird/game"
)

// NewGame init new game
func NewGame() *game.Game {
	g := &game.Game{}
	g.Init()
	return g
}

func main() {
	ebiten.SetWindowSize(config.SCREEN_WIDTH*2, config.SCREEN_HEIGHT*2)
	ebiten.SetWindowTitle("Flappy Bird")
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
