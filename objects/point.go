package objects

import (
	"fmt"

	"github.com/go-gl/gl/v2.1/gl"
)

// Point interface
type Point interface {
	Set2D(x, y float32)
	Set2DPoint(p Point)
	String() string
	Get() *point
	Clone() Point
}

type point struct {
	X, Y, Z float32
}

// NewPoint creates a new 3D point
func NewPoint(x, y, z float32) Point {
	return &point{
		X: x,
		Y: y,
		Z: z,
	}
}

// NewPoint2D creates a 2D point
func NewPoint2D(x, y float32) Point {
	return NewPoint(x, y, 0)
}

func (p *point) Set2D(x, y float32) {
	p.X = x
	p.Y = y
}

func (p *point) Set2DPoint(np Point) {
	npv := np.Get()
	p.Set2D(npv.X, npv.Y)
}

func (p *point) String() string {
	return fmt.Sprintf("(%v, %v)", p.X, p.Y)
}

func (p *point) Get() *point {
	return p
}

func (p *point) Clone() Point {
	pV := p.Get()

	// []gl.FLOAT => GLfloat[]
	newPoint := make([]float32, 4)
	var matrixGL [4][4]float32

	gl.GetFloatv(gl.MODELVIEW_MATRIX, &matrixGL[0][0])

	for i := range newPoint {
		newPoint[i] = matrixGL[0][i]*pV.X +
			matrixGL[1][i]*pV.Y +
			matrixGL[2][i]*pV.Z +
			matrixGL[3][i]
	}

	return NewPoint(newPoint[0], newPoint[1], newPoint[2])
}
