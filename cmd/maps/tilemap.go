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
	//Tileheight int `json: tilewidth`
	//Tilewidth int `json: tilewidth`
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
func (t *TilemapJSON) Draw(screen *ebiten.Image, tilesets []*Tileset, camcam *camera.Camera) {
	for layerIndex, layer := range t.Tiles {

		tileset := tilesets[layerIndex%len(tilesets)]
		for y := 0; y < layer.Height; y++ {
			for x := 0; x < layer.Width; x++ {
				//tileIndex := layer.Data[y*layer.Width + x]
				tileIndex := layer.Data[y*layer.Width+x] - 1 // Adjust for 1-based indexing

				// Skip empty tiles
				if tileIndex < 0 {
					continue
				}

				// Get the tile's position in the tileset
				tilesetWidthInTiles := tileset.Image.Bounds().Dx() / 16
				sx := (tileIndex % tilesetWidthInTiles) * 16
				sy := (tileIndex / tilesetWidthInTiles) * 16

				// Define the tile rectangle
				tileRect := image.Rect(sx, sy, sx+16, sy+16)

				// Create draw options
				opts := &ebiten.DrawImageOptions{}
				opts.GeoM.Translate(float64(x*16)+camcam.X, float64(y*16)+camcam.Y)

				// Draw the tile
				screen.DrawImage(tileset.Image.SubImage(tileRect).(*ebiten.Image), opts)
			}
		}
	}
}
