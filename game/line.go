package game

import (
	"github.com/go-gl/gl/v2.1/gl"
)

// Line interface
type Line interface {
	Draw()
	Get() *line
}

type line struct {
	x1, y1, x2, y2 float32
}

// NewLine creates a new Line
func NewLine(x1, y1, x2, y2 float32) Line {
	return &line{
		x1: x1,
		y1: y1,
		x2: x2,
		y2: y2,
	}
}

func (l *line) Get() *line {
	return l
}

// GenerateLine create a normal sized line
func GenerateLine() Line {
	x1 := float32(rnd.Int31()%100) / 10
	y1 := float32(rnd.Int31()%100) / 10
	x2 := float32(rnd.Int31()%100) / 10
	y2 := float32(rnd.Int31()%100) / 10

	// fmt.Printf("Line: (%v, %v) (%v, %v)\n", x1, y1, x2, y2)

	return NewLine(x1, y1, x2, y2)
}

// GenerateSmallLine create a small line
func GenerateSmallLine() Line {
	x1 := float32(rnd.Int31()%90) / 10
	y1 := float32(rnd.Int31()%90) / 10

	deltaX := float32(rnd.Int31()%100)/50 - 1
	deltaY := float32(rnd.Int31()%100)/50 - 1

	// fmt.Printf("Line: (%v, %v) (%v, %v)\n", x1, y1, x1+deltaX, y1+deltaY)

	return NewLine(x1, y1, x1+deltaX, y1+deltaY)
}

func (l *line) Draw() {
	gl.Begin(gl.LINES)
	{
		gl.Vertex2f(l.x1, l.y1)
		gl.Vertex2f(l.x2, l.y2)
	}
	gl.End()
}
