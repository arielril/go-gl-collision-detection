package game

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/go-gl/gl/v2.1/gl"

	"github.com/arielril/go-gl-collision-detection/objects"
	"github.com/arielril/go-gl-collision-detection/util"

	"github.com/go-gl/glfw/v3.3/glfw"
)

const maxLines int = 3000
const carStep = 0.2
const qtdFrames = 1000.0

const HORIZONTAL_SPLIT = 25
const VERTICAL_SPLIT = 2

const horizontalSize = 10
const verticalSize = 10

var rnd *rand.Rand
var lines []objects.Line
var car objects.Line
var collisionStructure Collision

var shouldTest bool
var shouldShowLines bool
var shouldShowFps bool

var tx, ty, alpha float32

func createLines() {
	rnd = rand.New(
		rand.NewSource(time.Now().UnixNano()),
	)

	car = objects.NewLineFromPoints(
		objects.NewPoint2D(0, 1),
		objects.NewPoint(0, -1, 0),
	)
	lines = make([]objects.Line, maxLines)

	for i := 0; i < maxLines; i++ {
		// lines[i] = objects.GenerateLine(rnd)
		lines[i] = objects.GenerateSmallLine(rnd)
	}

	// set initial position of the car
	tx = 5
	ty = 5
	alpha = 0
}

// Init the game
func Init() {
	shouldShowLines = true
	shouldTest = true

	fps = util.NewFps()

	createLines()

	// collisionStructure = CreateProfessorCollision()
	collisionStructure = CreateMyCollision(
		horizontalSize,
		verticalSize,
		lines,
		HORIZONTAL_SPLIT,
		VERTICAL_SPLIT,
	)

	// show menu
	printMenu()
}

// Display the game
func Display(w *glfw.Window) {
	gl.Clear(gl.COLOR_BUFFER_BIT)
	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()

	if shouldShowLines {
		displayLines()
	}

	resetLinesCollision()

	if shouldShowFps {
		baseTime := glfw.GetTime()

		for i := 0; i < qtdFrames; i++ {
			if i+1 >= qtdFrames {
				showQtyCellsTested = true
			}
			displayScenario()
		}
		showQtyCellsTested = false

		newTime := glfw.GetTime()
		fpsVal := qtdFrames / (newTime - baseTime)

		fmt.Printf("%v FPS.\n", fpsVal)
		shouldShowFps = false
	} else {
		displayScenario()
	}
}
