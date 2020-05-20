package game

import (
	"fmt"
)

// Point interface
type Point interface {
	Set(x, y float32)
	String() string
	Get() *point
}

type point struct {
	x, y, z float32
}

// NewPoint creates a new point
func NewPoint(x, y, z float32) Point {
	return &point{
		x: x,
		y: y,
		z: z,
	}
}

func (p *point) Set(x, y float32) {
	p.x = x
	p.y = y
}

func (p *point) String() string {
	return fmt.Sprintf("(%v, %v)", p.x, p.y)
}

func (p *point) Get() *point {
	return p
}
