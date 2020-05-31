package collision

import (
	"fmt"

	"github.com/arielril/go-gl-collision-detection/objects"
)

type cellPos struct {
	i, j int
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
