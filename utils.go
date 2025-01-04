package main

import "github.com/hajimehoshi/ebiten/v2"

func CheckAndProcessWalk(g *Game) {
	// TODO: replace hardcoded WASD keys to configurable
	if ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.IsKeyPressed(ebiten.KeyD) {
		g.X += PlayerMoveSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.IsKeyPressed(ebiten.KeyA) {
		g.X -= PlayerMoveSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) || ebiten.IsKeyPressed(ebiten.KeyW) {
		g.Y -= PlayerMoveSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) || ebiten.IsKeyPressed(ebiten.KeyS) {
		g.Y += PlayerMoveSpeed
	}
}
