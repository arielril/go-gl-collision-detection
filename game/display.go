package game

import (
	"fmt"

	"github.com/go-gl/gl/v2.1/gl"

	"github.com/arielril/basic-go-gl/util"
)

var fps util.FPS

func displayFps() {
	acc := fps.SetFPS().GetAccumulated()
	if acc >= 1 { // print every second
		fmt.Printf("FPS: %v\n", fps.GetFPS())
		fps.Reset()
	}
}

func instantiatePoint(p Point) Point {
	pV := p.Get()

	// []gl.FLOAT => GLfloat[]
	newPoint := make([]float32, 4)
	var matrixGL [4][4]float32

	gl.GetFloatv(gl.MODELVIEW_MATRIX, &matrixGL[0][0])

	for i := range newPoint {
		newPoint[i] = matrixGL[0][i]*pV.x +
			matrixGL[1][i]*pV.y +
			matrixGL[2][i]*pV.z +
			matrixGL[3][i]
	}

	return NewPoint(newPoint[0], newPoint[1], newPoint[2])
}

func displayScenario() {
	var p1 Point
	var p2 Point

	var pa Point
	var pb Point

	temp := NewPoint(0, 0, 0)

	carS := car.Get()

	gl.PushMatrix()
	{
		gl.Translatef(tx, ty, 0)
		gl.Rotatef(alpha, 0, 0, 1)
		temp.Set(carS.x1, carS.y1)
		p1 = instantiatePoint(temp)

		temp.Set(carS.x2, carS.y2)
		p2 = instantiatePoint(temp)

	}
	gl.PopMatrix()

	gl.LineWidth(1)
	gl.Color3f(1, 1, 0)

	for i := 0; i < maxLines; i++ {
		if shouldTest {
			lineV := lines[i].Get()

			temp.Set(lineV.x1, lineV.y1)
			pa = instantiatePoint(temp)

			temp.Set(lineV.x2, lineV.y2)
			pb = instantiatePoint(temp)

			if HasIntersection(pa, pb, p1, p2) {
				gl.Color3f(1, 0, 0)
			} else {
				gl.Color3f(0, 1, 0)
			}
		} else {
			gl.Color3f(0, 1, 0)
		}

		if shouldShowLines {
			lines[i].Draw()
		}
	}

	gl.Color3f(1, 0, 1)
	gl.LineWidth(3)
	gl.PushMatrix()
	{
		gl.Translatef(tx, ty, 0)
		gl.Rotatef(alpha, 0, 0, 1)
		car.Draw()
	}
	gl.PopMatrix()
}

func printMenu() {
	fmt.Println("f - imprime FPS.")
	fmt.Println("ESPACO - liga/desliga teste de colisao.")
}
