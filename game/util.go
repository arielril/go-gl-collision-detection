package game

import (
	"fmt"
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
