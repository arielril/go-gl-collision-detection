package game

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/arielril/go-gl-collision-detection/collision"
	"github.com/arielril/go-gl-collision-detection/objects"
)

// TurnShowLines turn on or off the lines in the display
func TurnShowLines() {
	shouldShowLines = !shouldShowLines
}

// ShowFps turn on or off the FPS count
func ShowFps() {
	shouldShowFps = true
}

// TurnShouldTest turn on or off the intersection test
func TurnShouldTest() {
	shouldTest = !shouldTest
	if shouldTest {
		fmt.Println("Interceccao LIGADA.")
	} else {
		fmt.Println("Interceccao DESLIGADA.")
	}
}

func resetLinesCollision() {
	for _, l := range lines {
		l.SetCollision(false)
	}
}

// ToggleRunBenchmark run the benchmark
func ToggleRunBenchmark() {
	shouldRunBenchmark = true
}

// ShowMenu x
func ShowMenu() {
	shouldShowMenu = true
}

// createLines of the environment
func createLines() {
	rnd = rand.New(
		rand.NewSource(time.Now().UnixNano()),
	)

	lines = make([]objects.Line, maxLines)

	for i := range lines {
		lines[i] = objects.GenerateLine(rnd, objects.LineSize.Small)
	}
}

func createCar() {
	car = objects.NewLineFromPoints(
		objects.NewPoint2D(0, 1),
		objects.NewPoint(0, -1, 0),
	)

	// set initial position of the car
	tx = 5
	ty = 5
	alpha = 0
}

func createCollisionStructure(hPieces, vPieces uint8) {
	collisionConfig := collision.NewConfig()

	collisionConfig.Lines = lines
	collisionConfig.Split.Horizontal = hPieces
	collisionConfig.Split.Vertical = vPieces
	collisionConfig.WindowSize.Width = horizontalSize
	collisionConfig.WindowSize.Height = verticalSize

	collisionStructure = collision.New(
		collisionConfig,
		collision.Provider.Professor,
		// collision.Provider.Me,
	)
}

func printMenu() {
	fmt.Println("f - imprime FPS.")
	fmt.Println("ESPACO - liga/desliga teste de colisao.")
}
