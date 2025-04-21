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

type Colliders struct {
	X, Y     float64
	Width    float64
	Height   float64
	Rotation float64
	Type     string
	TileID   int
	Meta     map[string]string
}

func (t *TilemapJSON) setColliders(tilesets []*Tileset) ([]*Colliders, error) {
	colliders := []*Colliders{}

	for layerIndex, layer := range t.Tiles {

		tileset := tilesets[layerIndex%len(tilesets)]
		for y := 0; y < layer.Height; y++ {
			for x := 0; x < layer.Width; x++ {
				tileIndex := layer.Data[y*layer.Width+x]

				if tileIndex == 0 {
					continue
				}
				for _, tile := range tileset.Tiles {
					if tile.ID+tileset.FirstGID == tileIndex && tile.ObjectGroup != nil {
						for _, obj := range tile.ObjectGroup.Objects {
							worldX := float64(x*16) + obj.X
							worldY := float64(y*16) + obj.Y

							colliders = append(colliders, &Colliders{
								X:        worldX,
								Y:        worldY,
								Width:    obj.Width,
								Height:   obj.Height,
								Rotation: obj.Rotation,
								Type:     obj.Type,
								Meta:     map[string]string{"name": obj.Name},
							})
						}
					}
				}

			}
		}
	}
	return colliders, nil
}
