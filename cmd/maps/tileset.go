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
		tiles = append(tiles, ImageTile{
			ID:          collection.ID,
			Image:       img,
			ImageWidth:  collection.ImageWidth,
			ImageHeight: collection.ImageHeight,
		})
	}
	return &Tileset{
		Tiles:    tiles,
		FirstGID: tilesetInfo.FirstGID,
	}, nil
}
