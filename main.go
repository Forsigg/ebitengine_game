package main

import (
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	playerSprite *Sprite
	sprites      []*Sprite
}

func (g *Game) Update() error {
	CheckAndProcessWalk(g.playerSprite)

	EnemyFollowByPlayer(g.playerSprite, g.sprites)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{120, 180, 255, 255})

	playerImgOpts := &ebiten.DrawImageOptions{}
	playerImgOpts.GeoM.Translate(g.playerSprite.X, g.playerSprite.Y)

	// draw player image
	screen.DrawImage(
		g.playerSprite.Img.SubImage(
			image.Rect(0, 0, 0+TileSize, 0+TileSize),
		).(*ebiten.Image),
		playerImgOpts,
	)
	playerImgOpts.GeoM.Reset()

	for _, sprite := range g.sprites {
		spriteOpts := &ebiten.DrawImageOptions{}
		spriteOpts.GeoM.Translate(sprite.X, sprite.Y)

		screen.DrawImage(
			sprite.Img.SubImage(
				image.Rect(0, 0, 0+TileSize, 0+TileSize),
			).(*ebiten.Image),
			spriteOpts,
		)
		spriteOpts.GeoM.Reset()
	}
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
		playerSprite: NewSprite(playerImg, 50, 50),
		sprites: []*Sprite{
			NewSprite(skeletonImg, 150, 150),
			NewSprite(skeletonImg, 200, 200),
		},
	}); err != nil {
		log.Fatal(err)
	}
}
