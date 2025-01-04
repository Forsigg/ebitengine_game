package main

import "github.com/hajimehoshi/ebiten/v2"

type Sprite struct {
	Img  *ebiten.Image
	X, Y float64
}

func NewSprite(img *ebiten.Image, x, y float64) *Sprite {
	return &Sprite{
		Img: img,
		X:   x,
		Y:   y,
	}
}
