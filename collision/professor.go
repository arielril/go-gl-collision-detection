package collision

import "github.com/arielril/go-gl-collision-detection/objects"

type profCollision struct {
	lines []objects.Line
}

// CreateProfessorCollision is a test
// Professor collision
func CreateProfessorCollision(config Config) Collision {
	some := &profCollision{
		lines: config.Lines,
	}

	return some
}

func (pc *profCollision) Collide(car objects.Line) {
	var p1, p2 objects.Point
	temp := objects.NewPoint(0, 0, 0)

	gameLines := pc.lines

	for i := 0; i < len(gameLines); i++ {
		lineV := gameLines[i].Get()

		temp.Set2DPoint(lineV.Pa)
		p1 = temp.Clone()

		temp.Set2DPoint(lineV.Pb)
		p2 = temp.Clone()

		lineT := objects.NewLineFromPoints(p1, p2)

		if car.Intersect(lineT) {
			gameLines[i].SetCollision(true)
		} else {
			gameLines[i].SetCollision(false)
		}
	}
}
