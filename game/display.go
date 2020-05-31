package game

import (
	"fmt"

	"github.com/go-gl/gl/v2.1/gl"

	"github.com/arielril/go-gl-collision-detection/objects"
	"github.com/arielril/go-gl-collision-detection/util"
)

var fps util.FPS

func displayFps() {
	acc := fps.SetFPS().GetAccumulated()
	if acc >= 1 { // print every second
		fmt.Printf("FPS: %v\n", fps.GetFPS())
		fps.Reset()
	}
}

func displayLines() {
	gl.LineWidth(2)
	for _, l := range lines {
		gl.PushMatrix()
		{
			if l.Get().Collision {
				gl.Color3f(1, 0, 0)
			} else {
				gl.Color3f(0, 1, 0)
			}
			l.Draw()
		}
		gl.PopMatrix()
	}
}

func displayScenario() {
	var p1 objects.Point
	var p2 objects.Point
	temp := objects.NewPoint2D(0, 0)

	carS := car.Get()

	gl.PushMatrix()
	{
		gl.Translatef(tx, ty, 0)
		gl.Rotatef(alpha, 0, 0, 1)

		temp.Set2DPoint(carS.Pa)
		p1 = temp.Clone()

		temp.Set2DPoint(carS.Pb)
		p2 = temp.Clone()
	}
	gl.PopMatrix()

	if shouldTest {
		collisionStructure.Collide(
			objects.NewLineFromPoints(p1, p2),
		)
	}

	gl.Color3f(1, 0, 1)
	gl.LineWidth(5)
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
