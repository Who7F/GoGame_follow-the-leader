package maps

import (
	"encoding/json"
	"fmt"
	"image"
	"os"
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

type TilesetJSON struct {
	ImagePath  string `json:"image"`
	TileWidth  int    `json:"tilewidth"`
	TileHeight int    `json:"tileheight"`
	TileCount  int    `json:"tilecount"`
	Columns    int    `json:"columns"`
	Tiles      []Tile `json:"tiles"`
}

type Tile struct {
	ID          int    `json:"id"`
	Image       string `json:"image"`
	ImageWidth  int    `json:"imagewidth"`
	ImageHeight int    `json:"imageheight"`
	ObjectGroup *ObjectGroup
}

type ObjectGroup struct {
	Objects []CollisionObject `json:"objects"`
}

type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type CollisionObject struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Type     string  `json:"type"`
	X        float64 `json:"x"`
	Y        float64 `json:"y"`
	Width    float64 `json:"width"`
	Height   float64 `json:"height"`
	Rotation float64 `json:"rotation"`

	Ellipe  bool    `json:"ellipse,omitempty"`
	Polygon []Point `json:"polygon,omitempty"`
	//Meta    map[string]string `json:"properties,omitempty"`
}

func LoadTilesets(tilemap *TilemapJSON) ([]TileProvider, error) {
	tileProvider := []TileProvider{}

	for _, tilesetInfo := range tilemap.Tilesets {

		contents, err := os.ReadFile(filepath.Join("assets/maps/tilesset/", tilesetInfo.Source))
		if err != nil {
			return nil, err
		}

		tilesetData := TilesetJSON{}
		err = json.Unmarshal(contents, &tilesetData)
		if err != nil {
			return nil, err
		}

		tiles, err := SetTilesetFromData(&tilesetData, &tilesetInfo)
		if err != nil {
			return nil, err
		}

		tileProvider = append(tileProvider, tiles)

	}
	return tileProvider, nil
}

func SetTilesetFromData(tilesetData *TilesetJSON, tilesetInfo *TilesetInfo) (TileProvider, error) {
	if tilesetData.ImagePath != "" {
		return SetTilesSize(tilesetData, tilesetInfo)
	}
	if len(tilesetData.Tiles) > 0 {
		return SetImageCollection(tilesetData, tilesetInfo)
	}
	return nil, fmt.Errorf("failed to load tileset: no image or tile data found in %s", tilesetInfo.Source)
}

func SetTilesSize(tilesetData *TilesetJSON, tilesetInfo *TilesetInfo) (*TileSlice, error) {
	path := filepath.Base(tilesetData.ImagePath)

	img, _, err := ebitenutil.NewImageFromFile(filepath.Join("assets/images/maps", path))
	if err != nil {
		return nil, fmt.Errorf("failed to load tileset image: %v", err)
	}

	return &TileSlice{
		image:       img,
		FirstGID:    tilesetInfo.FirstGID,
		ImageWidth:  tilesetData.TileWidth,
		ImageHeight: tilesetData.TileHeight,
	}, nil
}

func SetImageCollection(tilesetData *TilesetJSON, tilesetInfo *TilesetInfo) (*TileObj, error) {
	images := make(map[int]*ebiten.Image)
	groups := make(map[int]*ObjectGroup)

	for _, collection := range tilesetData.Tiles {
		path := filepath.Base(collection.Image)

		img, _, err := ebitenutil.NewImageFromFile(filepath.Join("assets/images/maps/buildings", path))
		if err != nil {
			return nil, fmt.Errorf("failed to load tileset image: %v", err)
		}
		obj := SetColliders(&collection)
		images[collection.ID] = img
		if obj != nil {
			groups[collection.ID] = obj
		}

	}
	return &TileObj{
		images:      images,
		FirstGID:    tilesetInfo.FirstGID,
		ObjectGroup: groups,
	}, nil
}

func SetColliders(tile *Tile) *ObjectGroup {
	if tile.ObjectGroup == nil || len(tile.ObjectGroup.Objects) == 0 {
		return nil
	}

	collisionObjects := []CollisionObject{}
	for _, obj := range tile.ObjectGroup.Objects {
		collisionObjects = append(collisionObjects, CollisionObject{
			ID:       obj.ID,
			Name:     obj.Name,
			X:        obj.X,
			Y:        obj.Y,
			Type:     obj.Type,
			Width:    obj.Width,
			Height:   obj.Height,
			Rotation: obj.Rotation,
			Ellipe:   obj.Ellipe,
			Polygon:  obj.Polygon,
		})
	}
	return &ObjectGroup{Objects: collisionObjects}
}
