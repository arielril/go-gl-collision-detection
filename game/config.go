package game

import (
	"math/rand"

	"github.com/arielril/go-gl-collision-detection/collision"
	"github.com/arielril/go-gl-collision-detection/objects"
	"github.com/arielril/go-gl-collision-detection/util"
)

const carStep = 0.2
const qtdFrames = 1000.0
const horizontalSize = 10
const verticalSize = 10

var maxLines = 250
var horizontalSplit uint8 = 25
var verticalSplit uint8 = 2

var rnd *rand.Rand
var lines []objects.Line
var car objects.Line
var collisionStructure collision.Collision

var shouldTest bool
var shouldShowLines bool
var shouldShowFps bool
var shouldRunBenchmark bool
var shouldShowMenu bool

var tx, ty, alpha float32

var fps util.FPS
