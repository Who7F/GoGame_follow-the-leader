package maps

import (
	"fmt"
	"image"
	"path/filepath"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type TileProvider interface {
	GetImage(tileIndex int) *ebiten.Image
	GetObjectGroup(tileIndex int) *ObjectGroup
}

type TileSlice struct {
	image       *ebiten.Image
	FirstGID    int
	ImageWidth  int
	ImageHeight int
	ObjectGroup *ObjectGroup
}

type TileObj struct {
	images      map[int]*ebiten.Image
	FirstGID    int
	ObjectGroup map[int]*ObjectGroup
}

func (t *TileSlice) GetImage(tileIndex int) *ebiten.Image {
	tileIndex -= t.FirstGID
	tileWidth := t.image.Bounds().Dx() / t.ImageWidth
	sx := (tileIndex % tileWidth) * t.ImageWidth
	sy := (tileIndex / tileWidth) * t.ImageHeight
	return t.image.SubImage(image.Rect(sx, sy, sx+t.ImageWidth, sy+t.ImageHeight)).(*ebiten.Image)
}

func (t *TileSlice) GetObjectGroup(tileIndex int) *ObjectGroup {
	if t.ObjectGroup != nil {
		return t.ObjectGroup
	}
	return nil
}

func (t *TileObj) GetImage(tileIndex int) *ebiten.Image {
	img, ok := t.images[tileIndex-t.FirstGID]
	if !ok {
		return nil
	}
	return img
}

func (t *TileObj) GetObjectGroup(tileIndex int) *ObjectGroup {
	return t.ObjectGroup[tileIndex-t.FirstGID]
}

type Tile struct {
	ID          int
	Image       string
	ImageWidth  int
	ImageHeight int
	ObjectGroup *ObjectGroup
}

type ObjectGroup struct {
	Objects []CollisionObject
}

type Point struct {
	X float64
	Y float64
}

type CollisionObject struct {
	ID       int
	Name     string
	Type     string
	X        float64
	Y        float64
	Width    float64
	Height   float64
	Rotation float64

	Ellipe  bool
	Polygon []Point
	//Meta    map[string]string `json:"properties,omitempty"`
}

func LoadTilesets(tilemap *TilemapTiled) ([]TileProvider, error) {
	tileProvider := []TileProvider{}

	for _, tilesetInfo := range tilemap.Tilesets {

		if tilesetInfo.Parsed == nil {
			return nil, fmt.Errorf("tikeset %s was not parsed", tilesetInfo.Source)
		}

		tilesetData := tilesetInfo.Parsed
		GID := tilesetInfo.FirstGID

		provider, err := SetTilesetFromData(tilesetData, &tilesetInfo, GID)
		if err != nil {
			return nil, err
		}

		tileProvider = append(tileProvider, provider)

	}
	return tileProvider, nil
}

func SetTilesetFromData(tilesetData *TilesetSourceTiled, tilesetInfo *TilesetTiled, GID int) (TileProvider, error) {
	if tilesetData.Image != "" {
		return SetTilesSize(tilesetData, tilesetInfo, GID)
	}
	if len(tilesetData.Tiles) > 0 {
		return SetImageCollection(tilesetData, tilesetInfo, GID)
	}
	return nil, fmt.Errorf("failed to load tileset: no image or tile data found in %s", tilesetInfo.Source)
}

func SetTilesSize(tilesetData *TilesetSourceTiled, tilesetInfo *TilesetTiled, GID int) (*TileSlice, error) {
	path := filepath.Base(tilesetData.Image)

	img, _, err := ebitenutil.NewImageFromFile(filepath.Join("assets/images/maps", path))
	if err != nil {
		return nil, fmt.Errorf("failed to load tileset image: %v", err)
	}

	return &TileSlice{
		image:       img,
		FirstGID:    GID,
		ImageWidth:  tilesetData.TileWidth,
		ImageHeight: tilesetData.TileHeight,
	}, nil
}

func SetImageCollection(tilesetData *TilesetSourceTiled, tilesetInfo *TilesetTiled, GID int) (*TileObj, error) {
	images := make(map[int]*ebiten.Image)
	groups := make(map[int]*ObjectGroup)

	for _, collection := range tilesetData.Tiles {
		path := filepath.Base(collection.Image)

		img, _, err := ebitenutil.NewImageFromFile(filepath.Join("assets/images/maps/buildings", path))
		if err != nil {
			return nil, fmt.Errorf("failed to load tileset image: %v", err)
		}
		//obj := SetColliders(&collection)
		images[collection.ID] = img
		//if obj != nil {
		//	groups[collection.ID] = obj
		//}

	}
	return &TileObj{
		images:      images,
		FirstGID:    GID,
		ObjectGroup: groups,
	}, nil
}

func SetColliders(tilemap *TilemapTiled) (*ObjectGroup, error) {
	allObjects := []CollisionObject{}

	for _, layer := range tilemap.Tiles {
		if layer.Type == "objectgroup" && len(layer.Objects) > 0 {
			for _, obj := range layer.Objects {
				allObjects = append(allObjects, ConverObject(obj))
			}
		}
	}

	for _, tileset := range tilemap.Tilesets {
		if tileset.Parsed == nil {
			continue
		}
		for _, tile := range tileset.Parsed.Tiles {
			if tile.ObjectGroup != nil && len(tile.ObjectGroup.Objects) > 0 {
				for _, obj := range tile.ObjectGroup.Objects {
					allObjects = append(allObjects, ConverObject(obj))
				}
			}
		}
	}
	if len(allObjects) == 0 {
		fmt.Println("No colliders found")
		return nil, nil
	}

	return &ObjectGroup{Objects: allObjects}, nil
}

func ConverObject(obj ObjectTiled) CollisionObject {
	return CollisionObject{
		ID:       obj.ID,
		Name:     obj.Name,
		X:        obj.X,
		Y:        obj.Y,
		Type:     obj.Type,
		Width:    obj.Width,
		Height:   obj.Height,
		Rotation: obj.Rotation,
		Ellipe:   obj.Ellipe,
		Polygon:  converPoints(obj.Polygon),
	}
}

func converPoints(in []PointTiled) []Point {
	points := make([]Point, len(in))
	for i, pt := range in {
		points[i] = Point{X: pt.X, Y: pt.Y}
	}
	return points
}
