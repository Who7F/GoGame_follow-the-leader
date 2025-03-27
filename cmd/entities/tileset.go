package entities

import(
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Tileset struct{
	Image *ebiten.Image
	TileWidth, TileHeight int
}

func LoadTileset(imagePath string, tileWidth, tileHeight int)(*Tileset, error){
	img, _, err := ebitenutil.NewImageFromFile(imagePath)
	if err != nil {
		return nil, fmt.Errorf("failed to load tileset image: %v", err)
	}
	
	tileset := &Tileset{
		Image: img,
		TileWidth: tileWidth,
		TileHeight: tileHeight,
	}
	return tileset, nil
}
