package game

type node struct {
	wmin, wmax float32
	hmin, hmax float32
	lists      []int
}

type collisionAlgorithm struct{}

// CollisionNode interface
type CollisionNode interface{}

// CollisionAlgorithm interface
type CollisionAlgorithm interface {
	CreateCollisionStructure(
		lines []Line,
		wWidth, wHeight float32,
		wSplit, hSplit uint8,
	) interface{}
}

// NewCollisionNode creates a new collision node
func NewCollisionNode(wmin, wmax, hmin, hmax float32) CollisionNode {
	return &node{
		wmin:  wmin,
		wmax:  wmax,
		hmin:  hmin,
		hmax:  hmax,
		lists: make([]int, 0),
	}
}

// NewCollisionAlgorithm creat a new collision algorithm
func NewCollisionAlgorithm() CollisionAlgorithm {
	return &collisionAlgorithm{}
}

func (ca *collisionAlgorithm) CreateCollisionStructure(
	lines []Line,
	wWidth, wHeight float32,
	wSplit, hSplit uint8,
) interface{} {
	return ca
}

var collisionList []CollisionNode

func CreateQuadrantStructure(wWidth, wHeight float32, wSplit, hSplit uint8) {
	collisionList = make([]CollisionNode, (wSplit+1)*(hSplit+1))

	wStepSize := wWidth / float32(wSplit)
	hStepSize := wHeight / float32(hSplit)

	// start in the top left of the window
	// and goes down right
	for i := wHeight; i > 0; i -= hStepSize {
		for j := float32(0); j < wWidth; j += wStepSize {

			quad := NewCollisionNode(
				j, j+wStepSize,
				i-hStepSize, i,
			)

			collisionList = append(collisionList, quad)
		}
	}
}

func getQuadFromLine(l Line) interface{} {
	for _, quad := range collisionList {
		// ld := l.Get()

		lineInsideQuadrant := false
		if lineInsideQuadrant {
			return quad
		}
	}

	return nil
}

func setLinesToQuadrants(lines []Line) {
	for _, l := range lines {
		lineQuad := getQuadFromLine(l)

		if lineQuad != nil {
			// lineQuad.AddLine(l)
		}
	}
}
