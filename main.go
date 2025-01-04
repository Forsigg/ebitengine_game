package main

import (
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	PlayerImage *ebiten.Image
	X, Y        float64
}

func (g *Game) Update() error {
	CheckAndProcessWalk(g)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{120, 180, 255, 255})

	playerImgOpts := &ebiten.DrawImageOptions{}
	playerImgOpts.GeoM.Translate(g.X, g.Y)

	// draw player image
	screen.DrawImage(
		g.PlayerImage.SubImage(
			image.Rect(0, 0, 0+TileSize, 0+TileSize),
		).(*ebiten.Image),
		playerImgOpts,
	)
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

	if err := ebiten.RunGame(&Game{
		PlayerImage: playerImg,
		X:           100,
		Y:           100,
	}); err != nil {
		log.Fatal(err)
	}
}
