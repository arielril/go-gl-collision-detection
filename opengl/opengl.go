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

// Reshape the GL window
func Reshape(w *glfw.Window) {
	width, height := w.GetFramebufferSize()

	gl.Clear(gl.COLOR_BUFFER_BIT)

	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()

	gl.Ortho(0, 10, 0, 10, 1, 0)

	gl.Viewport(0, 0, int32(width), int32(height))

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
