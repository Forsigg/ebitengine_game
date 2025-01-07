package main

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"

	"ebit_game/entities"
)

func CheckAndProcessWalk(p *entities.Sprite) {
	// TODO: replace hardcoded WASD keys to configurable
	if ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.IsKeyPressed(ebiten.KeyD) {
		p.X += PlayerMoveSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.IsKeyPressed(ebiten.KeyA) {
		p.X -= PlayerMoveSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) || ebiten.IsKeyPressed(ebiten.KeyW) {
		p.Y -= PlayerMoveSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) || ebiten.IsKeyPressed(ebiten.KeyS) {
		p.Y += PlayerMoveSpeed
	}
}

func EnemyFollowByPlayer(player *entities.Sprite, enemies []*entities.Enemy) {
	for _, enemy := range enemies {
		if !enemy.FollowPlayer {
			continue
		}

		// process X direction
		if enemy.X < player.X {
			enemy.X += PlayerMoveSpeed / 2.0
		} else if enemy.X > player.X {
			enemy.X -= PlayerMoveSpeed / 2.0
		}

		// process Y direction
		if enemy.Y > player.Y {
			enemy.Y -= PlayerMoveSpeed / 2.0
		} else if enemy.Y < player.Y {
			enemy.Y += PlayerMoveSpeed / 2.0
		}
	}
}

func DrawSprites(screen *ebiten.Image, camera *Camera, sprites []*entities.Sprite) {
	for _, sprite := range sprites {
		spriteOpts := &ebiten.DrawImageOptions{}
		spriteOpts.GeoM.Translate(sprite.X, sprite.Y)
		spriteOpts.GeoM.Translate(camera.X, camera.Y)

		screen.DrawImage(
			sprite.Img.SubImage(
				image.Rect(0, 0, 0+TileSize, 0+TileSize),
			).(*ebiten.Image),
			spriteOpts,
		)
		spriteOpts.GeoM.Reset()
	}
}
