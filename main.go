package main

import (
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	player  *Player
	sprites []*Sprite
	enemies []*Enemy
}

func (g *Game) Update() error {
	CheckAndProcessWalk(g.player.Sprite)

	EnemyFollowByPlayer(g.player.Sprite, g.enemies)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{120, 180, 255, 255})

	playerImgOpts := &ebiten.DrawImageOptions{}
	playerImgOpts.GeoM.Translate(g.player.X, g.player.Y)

	// draw player image
	screen.DrawImage(
		g.player.Img.SubImage(
			image.Rect(0, 0, 0+TileSize, 0+TileSize),
		).(*ebiten.Image),
		playerImgOpts,
	)
	playerImgOpts.GeoM.Reset()

	DrawSprites(screen, g.sprites)

	enemiesSprites := make([]*Sprite, 0)
	for _, enemy := range g.enemies {
		enemiesSprites = append(enemiesSprites, enemy.Sprite)
	}

	DrawSprites(screen, enemiesSprites)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	playerImg, _, err := ebitenutil.NewImageFromFile(ImagesPath + "ninja_sprite.png")
	if err != nil {
		log.Fatal(err)
	}

	skeletonImg, _, err := ebitenutil.NewImageFromFile(ImagesPath + "skeleton_sprite.png")
	if err != nil {
		log.Fatal(err)
	}

	if err := ebiten.RunGame(&Game{
		player: NewPlayer(playerImg, 50, 50, 100),
		enemies: []*Enemy{
			NewEnemy(skeletonImg, true, 150, 150),
			NewEnemy(skeletonImg, false, 200, 200),
		},
	}); err != nil {
		log.Fatal(err)
	}
}
