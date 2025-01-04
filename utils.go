package main

import "github.com/hajimehoshi/ebiten/v2"

func CheckAndProcessWalk(p *Sprite) {
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

func EnemyFollowByPlayer(player *Sprite, enemies []*Sprite) {
	for _, enemy := range enemies {
		if enemy.X < player.X {
			enemy.X += PlayerMoveSpeed / 2.0
		} else if enemy.X > player.X {
			enemy.X -= PlayerMoveSpeed / 2.0
		}

		if enemy.Y > player.Y {
			enemy.Y -= PlayerMoveSpeed / 2.0
		} else if enemy.Y < player.Y {
			enemy.Y += PlayerMoveSpeed / 2.0
		}
	}
}
