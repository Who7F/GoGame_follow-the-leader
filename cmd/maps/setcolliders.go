package maps

import (
	"fmt"
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
		tileset := tilemap.Tilesets[layerIndex]
		gid := tileset.FirstGID
		parsed, ok := tileset.GetParsed()
		if !ok {
			return
		}

		tile := tileIndex - gid

		if tile == 266 {
			fmt.Printf("one")
		}

		if len(parsed.Tiles) > 0 {
			object, ok := parsed.Tiles[tile].GetObjectGroup()
			if !ok {
				return
			}
			if len(object.Objects) > 0 {

				for _, obj := range object.Objects {
					fmt.Printf("id is %v. ", obj.ID)
					//allObjects = append(allObjects, (obj, float64(x*tilemap.TileWidth), float64(y*tilemap.TileHeight), nil))
				}
			}
		}

	})

	if len(allObjects) == 0 {
		fmt.Println("No colliders found")
		return nil, nil
	}

	return allObjects, nil
}

func extertLayerColliders(layers []LayerTiled) []ColliderProvider {
	colliders := []ColliderProvider{}
	for _, layer := range layers {
		if layer.Type == "objectgroup" && len(layer.Objects) > 0 {
			for _, obj := range layer.Objects {
				col := ConvertToCollider(obj, 0, 0, nil)
				colliders = append(colliders, col)
			}
		}
	}
	return colliders
}

func getTileObjetGroup(tilemap *TilemapTiled, layerIndex, tileIndex int) (*ObjectGroupTiled, bool) {
	if layerIndex >= len(tilemap.Layers) {
		return nil, false
	}
	tileset := tilemap.Tilesets[layerIndex]
	gid := tileset.FirstGID

	parsed, ok := tileset.GetParsed()
	if !ok || tileIndex < gid {
		return nil, false
	}

	tileID := tileIndex - gid
	if tileID >= len(parsed.Tiles) {
		return nil, false
	}

	return parsed.Tiles[tileID].ObjectGroup, parsed.Tiles[tileID].ObjectGroup != nil
}

/*
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
*/
