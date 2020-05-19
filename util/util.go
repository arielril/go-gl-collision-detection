package util

import (
	"os"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func setup() {
	gl.ClearColor(1, 1, 1, 1)
}

func keyCallback(
	w *glfw.Window,
	key glfw.Key,
	scancode int,
	action glfw.Action,
	mods glfw.ModifierKey,
) {
	switch key {
	case glfw.KeyEscape:
		os.Exit(0)
		break
	}
}

func charCallback(w *glfw.Window, char rune) {
	switch char {
	case 'q':
		w.SetShouldClose(true)
		break
	}
}

func reshape(w *glfw.Window) {
	width, height := w.GetFramebufferSize()

	gl.Viewport(0, 0, int32(width), int32(height))
	gl.Clear(gl.COLOR_BUFFER_BIT)

	gl.MatrixMode(gl.PROJECTION)

	gl.LoadIdentity()
	gl.Ortho(0, 10, 0, 10, 1, 0)

	gl.MatrixMode(gl.MODELVIEW)

	gl.LoadIdentity()
}
