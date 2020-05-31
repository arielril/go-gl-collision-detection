package game

import (
	"fmt"

	"github.com/arielril/go-gl-collision-detection/objects"
)

var showQtyCellsTested = false

// Collision interface
type Collision interface {
	Collide(car objects.Line)
}

type profCollision struct {
	p1, p2 objects.Point
}

type cellPos struct {
	i, j int
}

type windowCell struct {
	Lines []int
}

type myCollision struct {
	Cells                        [][]*windowCell
	horizontalSize, verticalSize int
}

// CreateProfessorCollision is a test
// Professor collision
func CreateProfessorCollision() Collision {
	some := &profCollision{}

	return some
}

func (pc *profCollision) Collide(car objects.Line) {
	var p1, p2 objects.Point
	temp := objects.NewPoint(0, 0, 0)

	for i := 0; i < len(lines); i++ {
		lineV := lines[i].Get()

		temp.Set2DPoint(lineV.Pa)
		p1 = temp.Clone()

		temp.Set2DPoint(lineV.Pb)
		p2 = temp.Clone()

		lineT := objects.NewLineFromPoints(p1, p2)

		if car.Intersect(lineT) {
			lines[i].SetCollision(true)
		} else {
			lines[i].SetCollision(false)
		}
	}
}

func existCell(c *cellPos, l []*cellPos) bool {
	for _, li := range l {
		if c.i == li.i && c.j == li.j {
			return true
		}
	}
	return false
}

func getCellsBetween(a, b *cellPos) []*cellPos {
	var minCell *cellPos
	var maxCell *cellPos

	if a.i < b.i ||
		a.j < b.j {
		minCell = a
		maxCell = b
	} else {
		maxCell = a
		minCell = b
	}

	res := make([]*cellPos, 0)

	if minCell.j == maxCell.j {
		for i := minCell.i; i < maxCell.i; i++ {
			ncell := &cellPos{i, minCell.j}

			if existCell(ncell, res) {
				continue
			}

			res = append(res, ncell)
		}
	} else if minCell.i == maxCell.i {
		for j := minCell.j; j < maxCell.j; j++ {
			ncell := &cellPos{minCell.i, j}

			if existCell(ncell, res) {
				continue
			}

			res = append(res, ncell)
		}
	} else {
		for i := minCell.i; i < maxCell.i; i++ {
			for j := minCell.j; j < maxCell.j; j++ {
				ncell := &cellPos{i, j}

				if existCell(ncell, res) {
					continue
				}

				res = append(res, ncell)
			}
		}
	}
	return res
}

func getCellsForLine(l objects.Line, wStep, hStep float32) []*cellPos {

	// fmt.Printf("Steps H: %v | V: %v\n", wStep, hStep)
	lineData := l.Get()
	linePa := lineData.Pa.Get()
	linePb := lineData.Pb.Get()

	// fmt.Printf("line pa (%v, %v)\n", linePa.X, linePa.Y)
	x := int(linePa.X / wStep)
	y := int(linePa.Y / hStep)
	cellPosA := &cellPos{y, x}

	// fmt.Printf("line pb (%v, %v)\n", linePb.X, linePb.Y)
	x = int(linePb.X / wStep)
	y = int(linePb.Y / hStep)

	cellPosB := &cellPos{y, x}

	res := []*cellPos{
		cellPosA,
	}

	if !existCell(cellPosB, res) {
		res = append(
			res,
			cellPosB,
		)
		res = append(res, getCellsBetween(cellPosA, cellPosB)...)
	}

	return res
}

func (cp *cellPos) String() string {
	return fmt.Sprintf("(%v, %v)", cp.i, cp.j)
}

// CreateMyCollision a
func CreateMyCollision(hSize, vSize int, lineList []objects.Line, wSplit, hSplit uint8) Collision {
	width, height := float32(hSize), float32(vSize)

	wStepSize := width / float32(hSplit)
	hStepSize := height / float32(wSplit)

	splittedWindow := make([][]*windowCell, 0)

	// create cells for window
	for i := float32(0); i < height; i += hStepSize {
		currentLine := make([]*windowCell, 0)

		for j := float32(0); j < width; j += wStepSize {
			currentLine = append(
				currentLine,
				&windowCell{
					Lines: make([]int, 0),
				},
			)
		}

		splittedWindow = append(splittedWindow, currentLine)
	}

	// fill the cells with lines
	for lIndex, l := range lineList {
		lineCells := getCellsForLine(l, wStepSize, hStepSize)

		for _, v := range lineCells {
			splittedWindow[v.i][v.j].Lines = append(
				splittedWindow[v.i][v.j].Lines,
				lIndex,
			)
		}
	}

	return &myCollision{
		Cells:          splittedWindow,
		verticalSize:   vSize,
		horizontalSize: hSize,
	}
}

func (mc *myCollision) Collide(car objects.Line) {
	width, height := float32(mc.horizontalSize), float32(mc.verticalSize)

	wStepSize := width / float32(VERTICAL_SPLIT)
	hStepSize := height / float32(HORIZONTAL_SPLIT)

	carCells := getCellsForLine(car, wStepSize, hStepSize)

	if showQtyCellsTested {
		fmt.Printf("Tested Cells: %v\n", len(carCells))
	}

	for _, carCell := range carCells {
		testCell := mc.Cells[carCell.i][carCell.j]

		for _, lineIndex := range testCell.Lines {
			if car.Intersect(lines[lineIndex]) {
				lines[lineIndex].SetCollision(true)
			} else {
				lines[lineIndex].SetCollision(false)
			}
		}
	}
}

func (mc *myCollision) String() string {
	str := "\n"

	for j := len(mc.Cells); j > 0; j-- {
		row := mc.Cells[j-1]

		str += fmt.Sprintf("%v: [", j-1)
		for _, col := range row {
			str += fmt.Sprintf("\t%v", col)
		}
		str += "\t]\n"
	}

	return str
}

/*
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
*/
