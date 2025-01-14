package main

import (
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"ebit_game/entities"
)

type Game struct {
	player      *entities.Player
	sprites     []*entities.Sprite
	enemies     []*entities.Enemy
	tilemapJSON *TilemapJSON
	tilesets    []Tileset
	tilemapImg  *ebiten.Image
	camera      *Camera
}

// Update method for Game. It calls every tick of game (by default 60 times per second)
func (g *Game) Update() error {
	CheckAndProcessWalk(g.player.Sprite)

	EnemyFollowByPlayer(g.player.Sprite, g.enemies)
	g.camera.FollowTarget(
		g.player.X+float64(
			TileSize,
		)/2, // with half tilesize offset for center camera on player sprite
		g.player.Y+float64(TileSize)/2, // same
	)
	g.camera.Constrain(
		float64(g.tilemapJSON.Layers[0].Width),
		float64(g.tilemapJSON.Layers[0].Height),
	)
	return nil
}

// Draw method for drawing sprites.
// It calls every frame (depends on monitor's refresh rate, for 60Hz is called 60 times per second)
func (g *Game) Draw(screen *ebiten.Image) {
	// screen.Fill(color.RGBA{120, 180, 255, 255})

	opts := &ebiten.DrawImageOptions{}

	// loop over the layers
	for layerIndex, layer := range g.tilemapJSON.Layers {
		for index, id := range layer.Data {

			if id == 0 {
				continue
			}

			x := (index % layer.Width) * TileSize
			y := (index / layer.Width) * TileSize

			img := g.tilesets[layerIndex].Img(id)

			opts.GeoM.Translate(float64(x), float64(y))
			opts.GeoM.Translate(0.0, -(float64(img.Bounds().Dy()))+float64(TileSize))
			opts.GeoM.Translate(g.camera.X, g.camera.Y)

			screen.DrawImage(img, opts)

			opts.GeoM.Reset()
		}
	}

	playerImgOpts := &ebiten.DrawImageOptions{}
	playerImgOpts.GeoM.Translate(g.player.X, g.player.Y)
	playerImgOpts.GeoM.Translate(g.camera.X, g.camera.Y)

	// draw player image
	screen.DrawImage(
		g.player.Img.SubImage(
			image.Rect(0, 0, 0+TileSize, 0+TileSize),
		).(*ebiten.Image),
		playerImgOpts,
	)
	playerImgOpts.GeoM.Reset()

	DrawSprites(screen, g.camera, g.sprites)

	enemiesSprites := make([]*entities.Sprite, 0)
	for _, enemy := range g.enemies {
		enemiesSprites = append(enemiesSprites, enemy.Sprite)
	}

	DrawSprites(screen, g.camera, enemiesSprites)
}

// Layout set screen widht and height for game
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Tutorial Ninja RPG")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	playerImg, _, err := ebitenutil.NewImageFromFile(ImagesPath + "ninja_sprite.png")
	if err != nil {
		log.Fatal(err)
	}

	skeletonImg, _, err := ebitenutil.NewImageFromFile(ImagesPath + "skeleton_sprite.png")
	if err != nil {
		log.Fatal(err)
	}
	tilemap, err := NewTilemapJSON(MapsPath + "map.json")
	if err != nil {
		log.Fatal(err)
	}

	tilemapImg, _, err := ebitenutil.NewImageFromFile(ImagesPath + "TilesetFloor.png")
	if err != nil {
		log.Fatal(err)
	}

	tilesets, err := tilemap.GenTilesets()
	if err != nil {
		log.Fatal(err)
	}

	if err := ebiten.RunGame(&Game{
		player: entities.NewPlayer(playerImg, 50, 50, 100),
		enemies: []*entities.Enemy{
			entities.NewEnemy(skeletonImg, true, 150, 150),
			entities.NewEnemy(skeletonImg, false, 200, 200),
		},
		tilemapJSON: tilemap,
		tilemapImg:  tilemapImg,
		camera:      NewCamera(0.0, 0.0),
		tilesets:    tilesets,
	}); err != nil {
		log.Fatal(err)
	}
}
