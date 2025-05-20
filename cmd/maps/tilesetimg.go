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
}

type TileSlice struct {
	image       *ebiten.Image
	FirstGID    int
	ImageWidth  int
	ImageHeight int
}

type TileObj struct {
	images   map[int]*ebiten.Image
	FirstGID int
}

func (t *TileSlice) GetImage(tileIndex int) *ebiten.Image {
	tileIndex -= t.FirstGID
	tilePerRow := t.image.Bounds().Dx() / t.ImageWidth
	sx := (tileIndex % tilePerRow) * t.ImageWidth
	sy := (tileIndex / tilePerRow) * t.ImageHeight
	return t.image.SubImage(image.Rect(sx, sy, sx+t.ImageWidth, sy+t.ImageHeight)).(*ebiten.Image)
}

func (t *TileObj) GetImage(tileIndex int) *ebiten.Image {
	img, ok := t.images[tileIndex-t.FirstGID]
	if !ok {
		return nil
	}
	return img
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

	for _, collection := range tilesetData.Tiles {
		path := filepath.Base(collection.Image)

		img, _, err := ebitenutil.NewImageFromFile(filepath.Join("assets/images/maps/buildings", path))
		if err != nil {
			return nil, fmt.Errorf("failed to load tileset image: %v", err)
		}

		images[collection.ID] = img

	}
	return &TileObj{
		images:   images,
		FirstGID: GID,
	}, nil
}
