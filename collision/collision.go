package collision

import (
	"fmt"

	"github.com/arielril/go-gl-collision-detection/objects"
)

var showQtyCellsTested = false

// Type is the type of the collision structure that will be created
type Type string

type collisionProvider struct {
	Me        Type
	Professor Type
}

// Provider enum
var Provider = &collisionProvider{
	Me:        "me",
	Professor: "prof",
}

var collisionDict = map[Type]func(Config) Collision{
	Provider.Me:        CreateMyCollision,
	Provider.Professor: CreateProfessorCollision,
}

// Collision interface
type Collision interface {
	Collide(car objects.Line)
}

// Config is the config for the collision structure
type Config struct {
	WindowSize struct {
		Width, Height float32
	}
	Lines []objects.Line
	Split struct {
		Horizontal uint8
		Vertical   uint8
	}
}

// NewConfig return a new collision config
func NewConfig() Config {
	return Config{}
}

// New creates a new collision
func New(cfg Config, t Type) Collision {
	var createFunc func(Config) Collision

	switch t {
	case Provider.Me:
		createFunc = CreateMyCollision
		break
	case Provider.Professor:
		createFunc = CreateProfessorCollision
		break
	default:
		panic(fmt.Sprintf("invalid collision type: %v", t))
	}

	return createFunc(cfg)
}
