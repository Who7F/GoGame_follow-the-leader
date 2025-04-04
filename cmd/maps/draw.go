package maps

import (
	"follow-the-leader/cmd/camera"

	"github.com/hajimehoshi/ebiten/v2"
)

// Draw the tilemap using the tileset
func (t *TilemapJSON) Draw(screen *ebiten.Image, tilesets []*Tileset, camcam *camera.Camera) {
	for layerIndex, layer := range t.Tiles {

		tileset := tilesets[layerIndex%len(tilesets)]
		for y := 0; y < layer.Height; y++ {
			for x := 0; x < layer.Width; x++ {

				tileIndex := layer.Data[y*layer.Width+x] // get indexing

				// Skip empty tiles
				if tileIndex == 0 {
					continue
				}

				if tileset.Image != nil {

					tileRect := tileset.singleTile(tileIndex, 16)

					// Create draw options
					opts := &ebiten.DrawImageOptions{}
					opts.GeoM.Translate(float64(x*16)+camcam.X, float64(y*16)+camcam.Y)

					// Draw the tile
					screen.DrawImage(tileRect, opts)
				} else {
					for _, tile := range tileset.Tiles {
						if tile.ID+tileset.FirstGID == tileIndex {
							// Create draw options
							opts := &ebiten.DrawImageOptions{}
							opts.GeoM.Translate(float64(x*16)+camcam.X, float64((y*16)-tile.ImageHeight+16)+camcam.Y)

							// Draw the tile
							screen.DrawImage(tile.Image, opts)
							break
						}
					}
				}

			}
		}
	}
}
