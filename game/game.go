package game

import "github.com/arielril/go-gl-collision-detection/collision"

var defaultConfig *gameConfig = &gameConfig{
	hSplit:        horizontalSplit,
	vSplit:        verticalSplit,
	nLines:        maxLines,
	showCar:       true,
	showLines:     shouldShowLines,
	showMenu:      shouldShowMenu,
	collisionType: collision.Provider.Professor,
}

// Init the game
func Init() {
	shouldShowLines = true
	shouldTest = true
	shouldRunBenchmark = false
	shouldShowMenu = true

	// fps = util.NewFps()

	createCar()
	createLines()
	createCollisionStructure(horizontalSplit, verticalSplit)

	if shouldShowMenu {
		// show menu
		printMenu()
		shouldShowMenu = false
	}
}

type gameConfig struct {
	hSplit, vSplit uint8
	nLines         int
	showCar        bool
	showLines      bool
	showMenu       bool
	collisionType  collision.Type
}

func newGameConfig() *gameConfig {
	return new(gameConfig)
}

// InitCustom init the game with customized configuration
func InitCustom(config *gameConfig) {
	if config.hSplit >= 2 {
		horizontalSplit = config.hSplit
	}

	if config.vSplit >= 2 {
		verticalSplit = config.vSplit
	}

	maxLines = config.nLines
	shouldShowLines = config.showLines
	shouldShowMenu = config.showMenu

	createLines()

	collisionConfig := collision.NewConfig()
	collisionConfig.Lines = lines
	collisionConfig.Split.Horizontal = config.hSplit
	collisionConfig.Split.Vertical = config.vSplit
	collisionConfig.WindowSize.Width = horizontalSize
	collisionConfig.WindowSize.Height = verticalSize

	collisionStructure = collision.New(
		collisionConfig,
		config.collisionType,
	)
}
