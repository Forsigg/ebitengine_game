package entities

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

type Enemy struct {
	*Sprite
	FollowPlayer bool
}

func NewEnemy(img *ebiten.Image, isFollow bool, x, y float64) *Enemy {
	return &Enemy{
		Sprite:       NewSprite(img, x, y),
		FollowPlayer: isFollow,
	}
}
