package game

import (
	"fmt"

	"github.com/arielril/basic-go-gl/util"

	"github.com/go-gl/glfw/v3.3/glfw"
)

// Init the game
func Init() {
	fmt.Println("Init game")
	fps = util.NewFps()
}

// Display the game
func Display(w *glfw.Window) {
	displayFps()
}
