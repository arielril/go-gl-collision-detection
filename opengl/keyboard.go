package opengl

import (
	"fmt"

	"github.com/go-gl/glfw/v3.3/glfw"

	"github.com/arielril/basic-go-gl/game"
)

// KeyCallback is a callback function for key strikes
func KeyCallback(
	w *glfw.Window,
	key glfw.Key,
	scancode int,
	action glfw.Action,
	mods glfw.ModifierKey,
) {
	if action == glfw.Press || action == glfw.Repeat {
		switch key {
		case glfw.KeyEscape:
			w.SetShouldClose(true) // close the window
			break
		case glfw.KeySpace:
			game.TurnShouldTest()
			break
		case glfw.KeyLeft:
			game.MoveCarLeft()
			break
		case glfw.KeyRight:
			game.MoveCarRight()
			break
		case glfw.KeyUp:
			game.MoveCarUp()
			break
		case glfw.KeyDown:
			game.MoveCarDown()
			break
		}
	}
}

// CharCallback is a callback function for char strikes
func CharCallback(w *glfw.Window, char rune) {
	switch char {
	case 'q':
		w.SetShouldClose(true)
		break
	case 'r':
		game.RotateCarRight()
		break
	case 'R':
		game.RotateCarLeft()
		break
	case 'e':
		game.TurnShowLines()
		break
	case 'f':
		game.ShowFps()
		fmt.Println("Comecou a contar...")
		break
	}
}
