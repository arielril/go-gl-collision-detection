package main

import (
	"fmt"
	"runtime"

	"github.com/arielril/basic-go-gl/game"
	"github.com/arielril/basic-go-gl/opengl"

	glfw "github.com/go-gl/glfw/v3.3/glfw"
)

const (
	wWidth  = 800
	wHeight = 600
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
	win, _ := opengl.NewWindow(wWidth, wHeight, "Space Invaders")
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
