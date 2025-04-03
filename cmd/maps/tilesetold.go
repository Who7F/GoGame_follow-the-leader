package maps

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func LoadTileset(imagePath string) (*Tileset, error) {
	img, _, err := ebitenutil.NewImageFromFile(imagePath)
	if err != nil {
		return nil, fmt.Errorf("failed to load tileset image: %v", err)
	}

	tileset := &Tileset{
		Image: img,
	}
	return tileset, nil
}
