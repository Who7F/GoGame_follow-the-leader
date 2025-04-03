package maps

import (
	"fmt"
	"follow-the-leader/cmd/camera"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Tileset struct {
	Image    *ebiten.Image
	FirstGID int
}

func (t *Tileset) singleTile(tileIndex, size int) *ebiten.Image {
	tileIndex -= t.FirstGID

	tilesetWidthInTiles := t.Image.Bounds().Dx() / size
	sx := (tileIndex % tilesetWidthInTiles) * size
	sy := (tileIndex / tilesetWidthInTiles) * size

	return t.Image.SubImage(image.Rect(sx, sy, sx+size, sy+size)).(*ebiten.Image)
}

func LoadTilesets(imagePaths []string, firstGIDs []int) ([]*Tileset, error) {
	tilesets := []*Tileset{}

	for i, path := range imagePaths {
		img, _, err := ebitenutil.NewImageFromFile(path)
		if err != nil {
			return nil, fmt.Errorf("failed to load tileset image: %v", err)
		}
		tileset := &Tileset{
			Image:    img,
			FirstGID: firstGIDs[i],
		}
		tilesets = append(tilesets, tileset)

	}
	return tilesets, nil
}

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

				tileRect := tileset.singleTile(tileIndex, 16)

				// Create draw options
				opts := &ebiten.DrawImageOptions{}
				opts.GeoM.Translate(float64(x*16)+camcam.X, float64(y*16)+camcam.Y)

				// Draw the tile
				screen.DrawImage(tileRect, opts)
			}
		}
	}
}
