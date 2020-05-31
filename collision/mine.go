package collision

import (
	"fmt"

	"github.com/arielril/go-gl-collision-detection/objects"
)

type myCollision struct {
	Cells                [][]*windowCell
	hStepSize, vStepSize float32
	Lines                []objects.Line
}

type windowCell struct {
	// Lines is the list of indexes of the lines
	Lines []int
}

// CreateMyCollision make my collision
func CreateMyCollision(config Config) Collision {
	width, height := config.WindowSize.Width, config.WindowSize.Height

	wStepSize := width / float32(config.Split.Vertical)
	hStepSize := height / float32(config.Split.Horizontal)

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
	for lIndex, l := range config.Lines {
		lineCells := getCellsForLine(l, wStepSize, hStepSize)

		for _, v := range lineCells {
			splittedWindow[v.i][v.j].Lines = append(
				splittedWindow[v.i][v.j].Lines,
				lIndex,
			)
		}
	}

	return &myCollision{
		Cells:     splittedWindow,
		hStepSize: wStepSize,
		vStepSize: hStepSize,
		Lines:     config.Lines,
	}
}

func (mc *myCollision) Collide(car objects.Line) {
	carCells := getCellsForLine(car, mc.hStepSize, mc.vStepSize)

	if showQtyCellsTested {
		fmt.Printf("Tested Cells: %v\n", len(carCells))
	}

	lines := mc.Lines

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
