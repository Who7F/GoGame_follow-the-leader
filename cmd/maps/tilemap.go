package maps

import (
	"encoding/json"
	"follow-the-leader/cmd/camera"
	"image"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

type TilemapLayerJSON struct {
	Data   []int  `json:"data"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Name   string `json:"name"`
}

type TilemapJSON struct {
	Tiles []TilemapLayerJSON `json:"layers"`
}

func NewTilemapJSON(filepath string) (*TilemapJSON, error) {
	contents, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	tilemapJSON := TilemapJSON{}
	err = json.Unmarshal(contents, &tilemapJSON)
	if err != nil {
		return nil, err
	}

	return &tilemapJSON, nil
}

// Draw the tilemap using the tileset
func (t *TilemapJSON) Draw(screen *ebiten.Image, tileset *Tileset, camcam *camera.Camera) {
	for _, layer := range t.Tiles {
		for y := 0; y < layer.Height; y++ {
			for x := 0; x < layer.Width; x++ {
				//tileIndex := layer.Data[y*layer.Width + x]
				tileIndex := layer.Data[y*layer.Width+x] - 1 // Adjust for 1-based indexing

				// Skip empty tiles
				if tileIndex < 0 {
					continue
				}

				// Get the tile's position in the tileset
				tilesetWidthInTiles := tileset.Image.Bounds().Dx() / tileset.TileWidth
				sx := (tileIndex % tilesetWidthInTiles) * tileset.TileWidth
				sy := (tileIndex / tilesetWidthInTiles) * tileset.TileHeight

				// Define the tile rectangle
				tileRect := image.Rect(sx, sy, sx+tileset.TileWidth, sy+tileset.TileHeight)

				// Create draw options
				opts := &ebiten.DrawImageOptions{}
				opts.GeoM.Translate(float64(x*tileset.TileWidth)+camcam.X, float64(y*tileset.TileHeight)+camcam.Y)

				// Draw the tile
				screen.DrawImage(tileset.Image.SubImage(tileRect).(*ebiten.Image), opts)
			}
		}
	}
}
