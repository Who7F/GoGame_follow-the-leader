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

type Tileset struct {
	Image    *ebiten.Image
	Tiles    []ImageTile
	FirstGID int
}

type ImageTile struct {
	ID          int
	Image       *ebiten.Image
	ImageWidth  int
	ImageHeight int
	ObjectGroup *ObjectGroup
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

type CollisionObject struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Type     string  `json:"type"`
	X        float64 `json:"x"`
	Y        float64 `json:"y"`
	Width    float64 `json:"width"`
	Height   float64 `json:"height"`
	Rotation float64 `json:"rotation"`
}

func (t *Tileset) singleTile(tileIndex, size int) *ebiten.Image {

	tileIndex -= t.FirstGID

	tilesetWidthInTiles := t.Image.Bounds().Dx() / size
	sx := (tileIndex % tilesetWidthInTiles) * size
	sy := (tileIndex / tilesetWidthInTiles) * size

	return t.Image.SubImage(image.Rect(sx, sy, sx+size, sy+size)).(*ebiten.Image)
}

func LoadTilesets(tilemap *TilemapJSON) ([]*Tileset, error) {
	tilesets := []*Tileset{}

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

		tileset, err := SetTilesetFromData(&tilesetData, &tilesetInfo)
		if err != nil {
			return nil, err
		}

		tilesets = append(tilesets, tileset)

	}
	return tilesets, nil
}

func SetTilesetFromData(tilesetData *TilesetJSON, tilesetInfo *TilesetInfo) (*Tileset, error) {
	if tilesetData.ImagePath != "" {
		return SetTilesSize(tilesetData, tilesetInfo)
	}
	if len(tilesetData.Tiles) > 0 {
		return SetImageCollection(tilesetData, tilesetInfo)
	}
	return nil, fmt.Errorf("failed to load tileset: no image or tile data found in %s", tilesetInfo.Source)
}

func SetTilesSize(tilesetData *TilesetJSON, tilesetInfo *TilesetInfo) (*Tileset, error) {
	path := filepath.Base(tilesetData.ImagePath)

	img, _, err := ebitenutil.NewImageFromFile(filepath.Join("assets/images/maps", path))
	if err != nil {
		return nil, fmt.Errorf("failed to load tileset image: %v", err)
	}

	return &Tileset{
		Image:    img,
		FirstGID: tilesetInfo.FirstGID,
	}, nil
}

func SetImageCollection(tilesetData *TilesetJSON, tilesetInfo *TilesetInfo) (*Tileset, error) {
	tiles := []ImageTile{}
	for _, collection := range tilesetData.Tiles {
		path := filepath.Base(collection.Image)

		img, _, err := ebitenutil.NewImageFromFile(filepath.Join("assets/images/maps/buildings", path))
		if err != nil {
			return nil, fmt.Errorf("failed to load tileset image: %v", err)
		}
		obj := SetColliders(&collection)

		tiles = append(tiles, ImageTile{
			ID:          collection.ID,
			Image:       img,
			ImageWidth:  collection.ImageWidth,
			ImageHeight: collection.ImageHeight,
			ObjectGroup: obj,
		})
	}
	return &Tileset{
		Tiles:    tiles,
		FirstGID: tilesetInfo.FirstGID,
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
		})
	}
	return &ObjectGroup{Objects: collisionObjects}
}
