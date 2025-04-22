package maps

import (
	"encoding/json"
	"os"
)

type TilemapLayerJSON struct {
	Data   []int  `json:"data"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Name   string `json:"name"`
}

type TilesetInfo struct {
	FirstGID int    `json:"firstgid"`
	Source   string `json:"source"`
}

type TilemapJSON struct {
	Tilesets   []TilesetInfo      `json:"tilesets"`
	Tiles      []TilemapLayerJSON `json:"layers"`
	TileWidth  int                `json:"tilewidth"`
	TileHeight int                `json:"tileheight"`
}

func NewTilemapJSON(filepath string) (*TilemapJSON, error) {
	contents, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	tilemapJSON := TilemapJSON{}
	err = json.Unmarshal(contents, &tilemapJSON)
	if err != nil {
		return nil, err
	}

	return &tilemapJSON, nil
}
