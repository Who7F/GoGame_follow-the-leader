package maps

import (
	"fmt"
	"path/filepath"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func SetColliders(tilemap *TilemapTiled) ([]ColliderProvider, error) {
	fmt.Printf("start")
	allObjects := []ColliderProvider{}

	//layer Object
	for _, layer := range tilemap.Layers {
		if layer.Type == "objectgroup" && len(layer.Objects) > 0 {
			for _, obj := range layer.Objects {
				col := ConvertToCollider(obj, 0, 0, nil)
				allObjects = append(allObjects, col)
			}
		}
	}

	//Tile Object

	tilemap.ForEachTile(func(layerIndex, tileIndex, x, y int) {
		layer := tilemap.Layers[layerIndex]
		if len(layer.Data) > 0 {
			if tileIndex == 111 {
				fmt.Printf("one")
			}
		}

		//ovider := providers[layerIndex%len(providers)]
		//oup := provider.GetObjectGroup(tileIndex)
		// group != nil {
		//or _, obj := range group.Objects {
		//col := ConvertToCollider(obj, float64(x*tilemap.TileWidth), float64(y*tilemap.TileHeight), nil)
		//allObjects = append(allObjects, col)
		//
		//
	})

	if len(allObjects) == 0 {
		fmt.Println("No colliders found")
		return nil, nil
	}

	return allObjects, nil
}

func SetImageCollectiona(tilesetData *TilesetSourceTiled, tilesetInfo *TilesetTiled, GID int) (*TileObj, error) {
	images := make(map[int]*ebiten.Image)
	//groups := make(map[int]*ObjectGroup)

	for _, collection := range tilesetData.Tiles {
		path := filepath.Base(collection.Image)

		img, _, err := ebitenutil.NewImageFromFile(filepath.Join("assets/images/maps/buildings", path))
		if err != nil {
			return nil, fmt.Errorf("failed to load tileset image: %v", err)
		}

		images[collection.ID] = img
		//if obj != nil {
		//	groups[collection.ID] = obj
		//}

	}
	return &TileObj{
		images:   images,
		FirstGID: GID,
		//ObjectGroup: groups,
	}, nil
}

/*
func (t *TileSlice) GetObjectGroup(tileIndex int) ColliderProvider {
	if t.ObjectGroup != nil {
		return t.ObjectGroup
	}
	return nil
}

func (t *TileObj) GetObjectGroup(tileIndex int) ColliderProvider {
	return t.ObjectGroup[tileIndex-t.FirstGID]
}*/

func findTile(tilemap *TilemapTiled, tileIndex int) (*TilesetTiled, *TileTiled) {
	for _, tileset := range tilemap.Tilesets {
		if tileset.Parsed == nil {
			continue
		}
		firstGID := tileset.FirstGID
		lastGID := firstGID + tileset.Parsed.TileCount

		if tileIndex >= firstGID && tileIndex < lastGID {
			tileID := tileIndex - firstGID
			for _, tile := range tileset.Parsed.Tiles {
				if tile.ID == tileID {
					return tileset, &tile
				}
			}
			return tileset, nil // Implicit tile
		}
	}
	return nil, nil
}
