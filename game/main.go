package game

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/arielril/basic-go-gl/tree"

	"github.com/arielril/basic-go-gl/util"

	"github.com/go-gl/glfw/v3.3/glfw"
)

const maxLines int = 300
const carStep = 0.2
const qtdFrames = 500.0

var rnd *rand.Rand
var lines []Line
var car Line
var collisionTree interface{}

var shouldTest bool
var shouldShowLines bool
var shouldShowFps bool

var tx, ty, alpha float32

func createLines() {
	rnd = rand.New(
		rand.NewSource(time.Now().UnixNano()),
	)

	car = NewLine(0, 1, 0, -1)
	lines = make([]Line, maxLines)

	for i := 0; i < maxLines; i++ {
		// lines[i] = GenerateLine()
		lines[i] = GenerateSmallLine()
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

	collisionTree = tree.NewTree()

	// show menu
	printMenu()
}

// Display the game
func Display(w *glfw.Window) {
	if shouldShowFps {
		baseTime := glfw.GetTime()

		for i := 0; i < qtdFrames; i++ {
			displayScenario()
		}

		newTime := glfw.GetTime()
		fpsVal := qtdFrames / (newTime - baseTime)

		fmt.Printf("%v FPS.\n", fpsVal)
		shouldShowFps = false
	} else {
		displayScenario()
	}
}
