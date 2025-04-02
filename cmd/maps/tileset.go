package maps

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Tileset struct {
	Image *ebiten.Image
}

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

func LoadTilesets(imagePaths []string) ([]*Tileset, error) {
	tilesets := []*Tileset{}

	for _, path := range imagePaths {
		img, _, err := ebitenutil.NewImageFromFile(path)
		if err != nil {
			return nil, fmt.Errorf("failed to load tileset image: %v", err)
		}
		tileset := &Tileset{
			Image: img,
		}
		tilesets = append(tilesets, tileset)

	}
	return tilesets, nil
}
