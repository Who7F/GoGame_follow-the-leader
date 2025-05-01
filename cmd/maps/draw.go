package maps

import (
	"fmt"
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
	Ellipe   bool
	Meta     map[string]string
	Polygon  []Point
}

func (t *TilemapJSON) SetColliders(providers []TileProvider) ([]ColliderProvider, error) {
	colliders := []ColliderProvider{}

	t.ForEachTile(providers, func(e *ebiten.Image, group *ObjectGroup, tileIndex, x, y int) {
		if group == nil {
			return
		}

		for _, obj := range group.Objects {

			worldX := float64(x*t.TileWidth) + obj.X
			worldY := float64((y * t.TileHeight) + t.TileHeight - int(obj.Height)) //+obj.Y

			meta := make(map[string]string)
			meta["name"] = obj.Name

			switch {
			case obj.Ellipe:
				colliders = append(colliders, &CircleColliders{
					X:      worldX + obj.Width/2,
					Y:      worldY + obj.Height/2,
					Layer:  0,
					Type:   obj.Type,
					Radius: obj.Width / 2,
					Meta:   meta,
				})
			case len(obj.Polygon) > 0:
				fmt.Printf("oh: %v wh:%v", obj.Height, worldY)
				points := make([]Point, len(obj.Polygon))
				for i, pt := range obj.Polygon {
					points[i] = Point{
						X: pt.X,
						Y: pt.Y - float64(e.Bounds().Dy()),
					}
				}
				colliders = append(colliders, &PolygonColliders{
					X:        worldX,
					Y:        worldY,
					Layer:    0,
					Width:    obj.Width,
					Height:   obj.Height,
					Rotation: obj.Rotation,
					//Polygon:  obj.Polygon,
					Polygon: points,
					Type:    obj.Type,
					Meta:    meta,
				})
			default:
				colliders = append(colliders, &RectColliders{
					X:        worldX,
					Y:        worldY,
					Layer:    0,
					Width:    obj.Width,
					Height:   obj.Height,
					Rotation: obj.Rotation,
					Type:     obj.Type,
					Meta:     meta,
				})
			}
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
