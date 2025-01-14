package main

import (
	"encoding/json"
	"os"
	"path"
)

type TilemapLayerJSON struct {
	Data   []int  `json:"data"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Name   string `json:"name"`
}

type TilemapJSON struct {
	Layers   []TilemapLayerJSON `json:"layers"`
	Tilesets []map[string]any   `json:"tilesets"`
}

func (t *TilemapJSON) GenTilesets() ([]Tileset, error) {
	tilesets := make([]Tileset, 0)

	for _, tilesetData := range t.Tilesets {
		tilesetPath := path.Join(MapsPath, tilesetData["source"].(string))
		tileset, err := NewTileset(tilesetPath, int(tilesetData["firstgid"].(float64)))
		if err != nil {
			return nil, err
		}

		tilesets = append(tilesets, tileset)
	}

	return tilesets, nil
}

func NewTilemapJSON(filepath string) (*TilemapJSON, error) {
	contents, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	var tilemap TilemapJSON
	err = json.Unmarshal(contents, &tilemap)
	if err != nil {
		return nil, err
	}
	return &tilemap, nil
}
