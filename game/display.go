package game

import (
	"fmt"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.3/glfw"

	"github.com/arielril/go-gl-collision-detection/collision"
	"github.com/arielril/go-gl-collision-detection/objects"
)

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

func displayCar() {
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

	if shouldShowLines {
		displayLines()
	}

	displayCar()
}

// Display the game
func Display(w *glfw.Window) {
	gl.Clear(gl.COLOR_BUFFER_BIT)
	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()

	resetLinesCollision()

	if shouldShowFps {
		baseTime := glfw.GetTime()

		for i := 0; i < qtdFrames; i++ {
			if i+1 >= qtdFrames {
				collision.SetShowQtdCellsTested(true)
			}
			displayScenario()
		}
		collision.SetShowQtdCellsTested(false)

		newTime := glfw.GetTime()
		fpsVal := qtdFrames / (newTime - baseTime)

		fmt.Printf("%v FPS.\n", fpsVal)
		shouldShowFps = false
	} else if shouldRunBenchmark {
		RunBenchmark()
		shouldRunBenchmark = false
	} else {
		displayScenario()
	}
}
