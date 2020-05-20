package game

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/arielril/basic-go-gl/util"

	"github.com/go-gl/glfw/v3.3/glfw"
)

const maxLines int = 300
const carStep = 0.2
const qtdFrames = 500.0

var rnd *rand.Rand
var lines []Line
var car Line

var shouldTest bool
var shouldShowLines bool
var shouldShowFps bool

var tx, ty, alpha float32

// Init the game
func Init() {
	shouldShowLines = true
	shouldTest = true

	fps = util.NewFps()

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

	// show menu
	printMenu()
}

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

// RotateCarLeft turn the car to the left
func RotateCarLeft() {
	alpha -= 3
}

// RotateCarRight turn the car to the right
func RotateCarRight() {
	alpha += 3
}

// MoveCarRight move the car to the right
func MoveCarRight() {
	tx += carStep
}

// MoveCarLeft move the car to the left
func MoveCarLeft() {
	tx -= carStep
}

// MoveCarUp move the car up
func MoveCarUp() {
	ty += carStep
}

// MoveCarDown move the car down
func MoveCarDown() {
	ty -= carStep
}

// Display the game
func Display(w *glfw.Window) {
	// displayFps()

	// gl.Clear(gl.COLOR_BUFFER_BIT)
	// gl.MatrixMode(gl.MODELVIEW)
	// gl.LoadIdentity()

	if shouldShowFps {
		baseTime := glfw.GetTime()

		for i := 0; i < qtdFrames; i++ {
			displayScenario()
		}
		newTime := glfw.GetTime()

		// fmt.Printf("nT %v, bT %v = %v\n", newTime, baseTime, (newTime-baseTime)/1000)

		fpsVal := qtdFrames / (newTime - baseTime)
		// fmt.Printf("qtd %v / sub %v = %v\n", qtdFrames, (newTime - baseTime), fpsVal)

		fmt.Printf("%v FPS.\n", fpsVal)
		shouldShowFps = false
	} else {
		displayScenario()
	}
}
