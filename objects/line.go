package objects

import (
	"fmt"
	"math/rand"

	"github.com/go-gl/gl/v2.1/gl"
)

// Line interface
type Line interface {
	Draw()
	Get() *line
	Intersect(l Line) bool
	SetCollision(v bool) Line
}

type line struct {
	Pa, Pb    Point
	Collision bool
}

func newLine(x1, y1, x2, y2 float32) Line {
	pa := NewPoint(x1, y1, 0)
	pb := NewPoint(x2, y2, 0)
	return &line{
		Pa:        pa,
		Pb:        pb,
		Collision: false,
	}
}

// NewLineFromPoints create a new line from 2 points in the space
func NewLineFromPoints(p1, p2 Point) Line {
	pv1 := p1.Get()
	pv2 := p2.Get()
	return newLine(pv1.X, pv1.Y, pv2.X, pv2.Y)
}

// _intersec2D Compute the intersection between 2 lines
/* ********************************************************************** */
/*                                                                        */
/*  Calcula a interseccao entre 2 retas (no plano "XY" Z = 0)             */
/*                                                                        */
/* k : ponto inicial da reta 1                                            */
/* l : ponto final da reta 1                                              */
/* m : ponto inicial da reta 2                                            */
/* n : ponto final da reta 2                                              */
/*                                                                        */
/* s: valor do parâmetro no ponto de interseção (sobre a reta KL)         */
/* t: valor do parâmetro no ponto de interseção (sobre a reta MN)         */
/*                                                                        */
/* ********************************************************************** */
func _intersec2D(k, l, m, n Point) (s, t float32, ok bool) {
	pk := k.Get()
	pl := l.Get()
	pm := m.Get()
	pn := n.Get()

	det := (pn.X-pm.X)*(pl.Y-pk.Y) - (pn.Y-pm.Y)*(pl.X-pk.X)

	if det == 0 {
		return 0, 0, false // no intersection
	}

	s = ((pn.X-pm.X)*(pm.Y-pk.Y) - (pn.Y-pm.Y)*(pm.X-pk.X)) / det
	t = ((pl.X-pk.X)*(pm.Y-pk.Y) - (pl.Y-pk.Y)*(pm.X-pk.X)) / det

	return s, t, true
}

func (l *line) Intersect(l2 Line) bool {
	l2v := l2.Get()

	s, t, ok := _intersec2D(
		l.Pa, l.Pb,
		l2v.Pa, l2v.Pb,
	)

	if !ok {
		return false
	}

	if s >= 0 && s <= 1 && t >= 0 && t <= 1 {
		return true
	}

	return false
}

func (l *line) Get() *line {
	return l
}

// GenerateLine create a normal sized line
func GenerateLine(rnd *rand.Rand) Line {
	x1 := float32(rnd.Int31()%100) / 10
	y1 := float32(rnd.Int31()%100) / 10
	x2 := float32(rnd.Int31()%100) / 10
	y2 := float32(rnd.Int31()%100) / 10

	// fmt.Printf("Line: (%v, %v) (%v, %v)\n", x1, y1, x2, y2)

	return newLine(x1, y1, x2, y2)
}

// GenerateSmallLine create a small line
func GenerateSmallLine(rnd *rand.Rand) Line {
	x1 := float32(rnd.Int31()%90) / 10
	y1 := float32(rnd.Int31()%90) / 10

	deltaX := float32(rnd.Int31()%100)/50 - 1
	deltaY := float32(rnd.Int31()%100)/50 - 1

	x2 := x1 + deltaX
	y2 := y1 + deltaY

	if x2 < 0 {
		x2 = 0
	}
	if y2 < 0 {
		y2 = 0
	}

	// fmt.Printf("Line: (%v, %v) (%v, %v)\n", x1, y1, x1+deltaX, y1+deltaY)

	return newLine(x1, y1, x2, y2)
}

func (l *line) Draw() {
	gl.Begin(gl.LINES)
	{
		lpav := l.Pa.Get()
		lpbv := l.Pb.Get()

		gl.Vertex2f(lpav.X, lpav.Y)
		gl.Vertex2f(lpbv.X, lpbv.Y)
	}
	gl.End()
}

func (l *line) SetCollision(v bool) Line {
	l.Collision = v
	return l
}

func (l *line) String() string {
	return fmt.Sprintf("Pa=%v Pb=%v", l.Pa, l.Pb)
}
