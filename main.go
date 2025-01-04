package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	PlayerImage *ebiten.Image
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{120, 180, 255, 255})

	// draw player image
	screen.DrawImage(
		g.PlayerImage,
		&ebiten.DrawImageOptions{},
	)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	playerImg, _, err := ebitenutil.NewImageFromFile(IMAGES_PATH + "ninja_sprite.png")
	if err != nil {
		log.Fatal(err)
	}

	if err := ebiten.RunGame(&Game{
		PlayerImage: playerImg,
	}); err != nil {
		log.Fatal(err)
	}
}
