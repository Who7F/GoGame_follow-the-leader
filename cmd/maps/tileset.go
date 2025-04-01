package maps

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Tileset struct {
	Image                 *ebiten.Image
	TileWidth, TileHeight int
}

func LoadTileset(imagePath string, tileWidth, tileHeight int) (*Tileset, error) {
	img, _, err := ebitenutil.NewImageFromFile(imagePath)
	if err != nil {
		return nil, fmt.Errorf("failed to load tileset image: %v", err)
	}

	tileset := &Tileset{
		Image:      img,
		TileWidth:  tileWidth,
		TileHeight: tileHeight,
	}
	return tileset, nil
}

func LoadTilesets(imagePaths []string, tileWidth, tileHeight int) ([]*Tileset, error) {
	tilesets := []*Tileset{}

	for _, path := range imagePaths {
		img, _, err := ebitenutil.NewImageFromFile(path)
		if err != nil {
			return nil, fmt.Errorf("failed to load tileset image: %v", err)
		}
		tileset := &Tileset{
			Image:      img,
			TileWidth:  tileWidth,
			TileHeight: tileHeight,
		}
		tilesets = append(tilesets, tileset)

	}
	return tilesets, nil
}
