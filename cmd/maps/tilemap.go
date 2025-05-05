package maps

import (
	"encoding/json"
	"fmt"
	"os"
)

const TilesetDir = "assets/maps/tilesset/"

func NewTilemapJSON(filepath string) (*TilemapTiled, error) {
	contents, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to find file %w", err)
	}

	tilemapTiled := TilemapTiled{}
	err = json.Unmarshal(contents, &tilemapTiled)
	if err != nil {
		return nil, fmt.Errorf("failed to load tilemap %w", err)
	}

	for i, ts := range tilemapTiled.Tilesets {
		tilesetPath := TilesetDir + ts.Source
		data, err := os.ReadFile(tilesetPath)
		if err != nil {
			return nil, fmt.Errorf("failed to read tileset %s: %w", tilesetPath, err)
		}
		parsed := TilesetSourceTiled{}
		err = json.Unmarshal(data, &parsed)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal tileset %s: %w", tilesetPath, err)
		}
		tilemapTiled.Tilesets[i].Parsed = &parsed
	}

	return &tilemapTiled, nil
}
