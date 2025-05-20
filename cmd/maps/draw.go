package maps

import (
	"follow-the-leader/cmd/camera"

	"github.com/hajimehoshi/ebiten/v2"
)

func (t *TilemapTiled) Draw(screen *ebiten.Image, providers []TileProvider, cam *camera.Camera) {
	t.ForEachTile(func(layerIndex, tileIndex, x, y int) {
		provider := providers[layerIndex%len(providers)]
		img := provider.GetImage(tileIndex)
		if img == nil {
			return
		}
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(float64(x*t.TileWidth)+cam.X, float64((y*t.TileHeight)-img.Bounds().Dy()+t.TileHeight)+cam.Y)
		screen.DrawImage(img, opts)
	})
}

// Look here Python. Here is how you do decorators without the garbage @ syntax
func (t *TilemapTiled) ForEachTile(fn func(layerIndex, tileIndex, x, y int)) {
	for layerIndex, layer := range t.Layers {
		if layer.Type != "tilelayer" || len(layer.Data) == 0 {
			continue
		}

		for y := 0; y < layer.Height; y++ {
			for x := 0; x < layer.Width; x++ {
				tileIndex := layer.Data[y*layer.Width+x]

				if tileIndex == 0 {
					continue
				}

				fn(layerIndex, tileIndex, x, y)
			}
		}
	}
}

func (t *TilemapTiled) ForEachTileOld(providers []TileProvider, fn func(image *ebiten.Image, tileIndex, x, y int)) {
	for layerIndex, layer := range t.Layers {
		if layer.Type != "tilelayer" || len(layer.Data) == 0 {
			continue
		}

		provider := providers[layerIndex%len(providers)]
		for y := 0; y < layer.Height; y++ {
			for x := 0; x < layer.Width; x++ {
				tileIndex := layer.Data[y*layer.Width+x]

				if tileIndex == 0 {
					continue
				}

				img := provider.GetImage(tileIndex)
				if img == nil {
					continue
				}
				//group := provider.GetObjectGroup(tileIndex)
				fn(img, tileIndex, x, y)
			}
		}
	}
}
