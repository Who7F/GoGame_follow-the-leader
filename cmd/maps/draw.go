package maps

import (
	"follow-the-leader/cmd/camera"

	"github.com/hajimehoshi/ebiten/v2"
)

// whole map.
func (t *TilemapJSON) Draw(screen *ebiten.Image, provider []TileProvider, cam *camera.Camera) {
	t.ForEachTile(provider, func(img *ebiten.Image, group *ObjectGroup, tileIndex, x, y int) {
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(float64(x*t.TileWidth)+cam.X, float64((y*t.TileHeight)-img.Bounds().Dy()+t.TileHeight)+cam.Y)
		screen.DrawImage(img, opts)
	})
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

func (t *TilemapJSON) SetColliders(providers []TileProvider) ([]*Colliders, error) {
	colliders := []*Colliders{}

	t.ForEachTile(providers, func(_ *ebiten.Image, group *ObjectGroup, tileIndex, x, y int) {
		if group == nil {
			return
		}
		for _, obj := range group.Objects {
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

	})
	return colliders, nil
}

// Look here Python. Here is how you do decorators without the garbage @ syntax
func (t *TilemapJSON) ForEachTile(providers []TileProvider, fn func(image *ebiten.Image, group *ObjectGroup, tileIndex, x, y int)) {
	for layerIndex, layer := range t.Tiles {

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
				group := provider.GetObjectGroup(tileIndex)
				fn(img, group, tileIndex, x, y)
			}
		}
	}
}

//func FindTileByID(tileIndex int, tileset *Tileset) *ImageTile {
//	for _, tile := range tileset.Tiles {
//		if tile.ID+tileset.FirstGID == tileIndex {
//			return &tile
//		}
//	}
//	return nil
//}
