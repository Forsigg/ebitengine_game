package main

import "math"

type Camera struct {
	X, Y float64
}

// NewCamera is a constructor for Camera struct.
func NewCamera(x, y float64) *Camera {
	return &Camera{x, y}
}

// FollowTarget change camera X and Y coords by the target X and Y.
func (c *Camera) FollowTarget(targetX, targetY float64) {
	c.X = -targetX + float64(ScreenWidth)/2
	c.Y = -targetY + float64(ScreenHeight)/2
}

// Constrain set camera X and Y that not allow camera out of map borders.
// By the way, if camera already in borders, it stop follow by target until it near map borders.
func (c *Camera) Constrain(tilemapWidthPixels, tilemapHeightPixels float64) {
	c.X = math.Min(c.X, 0.0)
	c.Y = math.Min(c.Y, 0.0)

	c.X = math.Max(c.X, tilemapWidthPixels-float64(ScreenWidth))
	c.Y = math.Max(c.Y, tilemapHeightPixels-float64(ScreenHeight))
}
