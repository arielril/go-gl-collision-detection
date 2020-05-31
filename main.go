package main

import (
	"fmt"
	"runtime"

	"github.com/arielril/go-gl-collision-detection/game"
	"github.com/arielril/go-gl-collision-detection/opengl"

	glfw "github.com/go-gl/glfw/v3.3/glfw"
)

const (
	wWidth  = 600
	wHeight = 500
)

func init() {
	runtime.LockOSThread()

	if err := glfw.Init(); err != nil {
		panic(fmt.Errorf("failed to init glfw: %v", err))
	}

	glfw.WindowHint(glfw.Resizable, glfw.True)
	glfw.WindowHint(glfw.ContextVersionMajor, 2)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
}

func main() {
	win, _ := opengl.NewWindow(wWidth, wHeight, "Cálculo de colisão.")
	defer glfw.Terminate()

	win.SetKeyCallback(opengl.KeyCallback)
	win.SetCharCallback(opengl.CharCallback)

	opengl.Setup()
	game.Init()
	for !win.ShouldClose() {
		opengl.Reshape(win)
		game.Display(win)

		win.SwapBuffers()
		glfw.PollEvents()
	}
}
