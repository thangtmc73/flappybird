package resources

import (
	"image/png"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	// mapSpriteByKey contains data loaded from image
	mapSpriteByKey map[string]*ebiten.Image
)

func loadImageToSprite(path string, key string) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	img, err := png.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	mapSpriteByKey[key] = ebiten.NewImageFromImage(img)
}

// Init to initialize resources
func Init() {
	mapSpriteByKey = make(map[string]*ebiten.Image)
	loadImageToSprite("resources/background.png", "background")
	loadImageToSprite("resources/bird.png", "bird")
}

// GetSpriteByKey allow getting resource with key
func GetSpriteByKey(key string) *ebiten.Image {
	sprite, ok := mapSpriteByKey[key]
	if ok {
		return sprite
	}
	return nil
}
