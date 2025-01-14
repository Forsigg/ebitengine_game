package main

import (
	"encoding/json"
	"image"
	"os"
	"path/filepath"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Tileset interface {
	Img(id int) *ebiten.Image
}

type UniformTilesetJSON struct {
	Path string `json:"image"`
	Gid  int
}

type UniformTileset struct {
	img *ebiten.Image
	gid int
}

func (u *UniformTileset) Img(id int) *ebiten.Image {
	id -= u.gid

	srcX := id % 22
	srcY := id / 22

	srcX *= TileSize
	srcY *= TileSize

	return u.img.SubImage(
		image.Rect(
			srcX, srcY, srcX+TileSize, srcY+TileSize,
		),
	).(*ebiten.Image)
}

type TileJSON struct {
	Id     int    `json:"id"`
	Path   string `json:"image"`
	Width  int    `json:"imagewidth"`
	Height int    `json:"imageheight"`
}

type DynamicTilesetJSON struct {
	Tiles []*TileJSON `json:"tiles"`
}

type DynamicTileset struct {
	imgs []*ebiten.Image
	gid  int
}

func (d *DynamicTileset) Img(id int) *ebiten.Image {
	id -= d.gid

	return d.imgs[id]
}

func NewTileset(path string, gid int) (Tileset, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	if strings.Contains(path, "buildings") {
		// return DinamicTileset
		var dynTilesetJSON DynamicTilesetJSON
		err = json.Unmarshal(content, &dynTilesetJSON)
		if err != nil {
			return nil, err
		}
		dynTileset := DynamicTileset{}
		dynTileset.gid = gid
		dynTileset.imgs = make([]*ebiten.Image, 0)

		for _, tileJSON := range dynTilesetJSON.Tiles {
			tileJSONPath := filepath.Clean(tileJSON.Path)
			tileJSONPath = strings.ReplaceAll(tileJSONPath, "\\", "/")
			tileJSONPath = strings.TrimPrefix(tileJSONPath, "../")
			tileJSONPath = strings.TrimPrefix(tileJSONPath, "../")
			tileJSONPath = filepath.Join(AssetsPath, tileJSONPath)

			img, _, err := ebitenutil.NewImageFromFile(tileJSONPath)
			if err != nil {
				return nil, err
			}

			dynTileset.imgs = append(dynTileset.imgs, img)
		}

		return &dynTileset, nil
	}
	// return UniformTileset

	var unifiormTilesetJSON UniformTilesetJSON
	err = json.Unmarshal(content, &unifiormTilesetJSON)
	if err != nil {
		return nil, err
	}

	uniformTileset := UniformTileset{}
	tileJSONPath := filepath.Clean(unifiormTilesetJSON.Path)
	tileJSONPath = strings.ReplaceAll(tileJSONPath, "\\", "/")
	tileJSONPath = strings.TrimPrefix(tileJSONPath, "../")
	tileJSONPath = strings.TrimPrefix(tileJSONPath, "../")
	tileJSONPath = filepath.Join(AssetsPath, tileJSONPath)
	img, _, err := ebitenutil.NewImageFromFile(tileJSONPath)
	if err != nil {
		return nil, err
	}

	uniformTileset.img = img
	uniformTileset.gid = gid

	return &uniformTileset, nil
}
