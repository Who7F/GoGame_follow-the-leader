package maps

import (
	"follow-the-leader/cmd/camera"

	"github.com/hajimehoshi/ebiten/v2"
)

type TileRenderer interface {
	Draw(screen *ebiten.Image, x, y float64, cam *camera.Camera)
}

type AtlasTile struct {
	Tileset   *Tileset
	TileIndex int
	TileSize  int
}

func (a *AtlasTile) Draw(screen *ebiten.Image, x, y float64, cam *camera.Camera) {
	a.Tileset.Draw(screen, a.TileIndex, a.TileSize, x, y, cam)
}

// whole map.
func (t *TilemapJSON) Draw(screen *ebiten.Image, tilesets []*Tileset, cam *camera.Camera) {
	t.ForEachTile(tilesets, func(tile *ImageTile, tileIndex, x, y int) {
		tile.Draw(screen,
			float64(x*t.TileWidth),
			float64((y*t.TileHeight)-tile.ImageHeight+t.TileHeight),
			tilesets[0],
			cam)
	})
}

// image
func (i *ImageTile) Draw(screen *ebiten.Image, x, y float64, tileset *Tileset, cam *camera.Camera) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(x+cam.X, y+cam.Y)
	screen.DrawImage(i.Image, opts)
}

func (t *Tileset) Draw(screen *ebiten.Image, tileIndex, tileSize int, x, y float64, cam *camera.Camera) {
	tileRect := t.singleTile(tileIndex, tileSize)
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(x+cam.X, y+cam.Y)
	screen.DrawImage(tileRect, opts)
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

	t.ForEachTile(tilesets, func(tile *ImageTile, tileIndex, x, y int) {
		if tile.ObjectGroup == nil {
			return
		}
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

	})
	return colliders, nil
}

// Look here Python. Here is how you do decorators without the garbage @ syntax
func (t *TilemapJSON) ForEachTile(tilesets []*Tileset, fn func(tile *ImageTile, tikeIndex, x, y int)) {
	for layerIndex, layer := range t.Tiles {

		tileset := tilesets[layerIndex%len(tilesets)]
		for y := 0; y < layer.Height; y++ {
			for x := 0; x < layer.Width; x++ {
				tileIndex := layer.Data[y*layer.Width+x]

				if tileIndex == 0 {
					continue
				}

				tile := FindTileByID(tileIndex, tileset)
				if tile != nil {
					fn(tile, tileIndex, x, y)
				}

			}
		}
	}
}

func FindTileByID(tileIndex int, tileset *Tileset) *ImageTile {
	for _, tile := range tileset.Tiles {
		if tile.ID+tileset.FirstGID == tileIndex {
			return &tile
		}
	}

	return nil
}
