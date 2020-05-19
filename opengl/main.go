package opengl

import (
	"fmt"

	"github.com/go-gl/gl/v2.1/gl"
	glfw "github.com/go-gl/glfw/v3.3/glfw"
)

// Setup the gl environment
func Setup() {
	gl.ClearColor(255, 255, 255, 1)
}

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
		}
	}
}

// CharCallback is a callback function for char strikes
func CharCallback(w *glfw.Window, char rune) {
	switch char {
	case 'q':
		w.SetShouldClose(true)
		break
	}
}

// Reshape the GL window
func Reshape(w *glfw.Window) {
	width, height := w.GetFramebufferSize()

	gl.Viewport(0, 0, int32(width), int32(height))
	gl.Clear(gl.COLOR_BUFFER_BIT)

	gl.MatrixMode(gl.PROJECTION)

	gl.LoadIdentity()
	gl.Ortho(0, 10, 0, 10, 1, 0)

	gl.MatrixMode(gl.MODELVIEW)

	gl.LoadIdentity()
}

// NewWindow creates a new GL window
func NewWindow(w, h int, title string) (*glfw.Window, error) {
	win, err := glfw.CreateWindow(w, h, title, nil, nil)
	if err != nil {
		panic(fmt.Errorf("failed to create the window: %v", err))
	}

	win.MakeContextCurrent()
	glfw.SwapInterval(1)

	if err := gl.Init(); err != nil {
		panic(fmt.Errorf("failed to init openGL: %v", err))
	}

	return win, err
}
