package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/thangtmc73/flappybird/config"
	"github.com/thangtmc73/flappybird/game"
	"github.com/thangtmc73/flappybird/resources"
)

// NewGame init new game
func NewGame() *game.Game {
	resources.Init()
	g := &game.Game{}
	g.Init()
	return g
}

func main() {
	ebiten.SetWindowSize(config.ScreenWidth*2, config.ScreenHeight*2)
	ebiten.SetWindowTitle("Flappy Bird")
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
