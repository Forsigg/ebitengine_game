package main

import "github.com/hajimehoshi/ebiten/v2"

type Player struct {
	*Sprite
	HP float64
}

func NewPlayer(img *ebiten.Image, x, y, hp float64) *Player {
	return &Player{
		Sprite: NewSprite(img, x, y),
		HP:     hp,
	}
}
