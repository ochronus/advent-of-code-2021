package solutions

import (
	"fmt"
	"ochronus/aoc2021/datastructures"
	"ochronus/aoc2021/utils"
)

func parseDay17Input() (areabounds datastructures.AreaBounds) {
	contents := utils.ReadFileToString("../inputs/17.txt")
	p1 := datastructures.Coordinate{}
	p2 := datastructures.Coordinate{}
	fmt.Sscanf(contents, "target area: x=%d..%d, y=%d..%d", &p1.X, &p2.X, &p1.Y, &p2.Y)
	areabounds.TopLeft = p1
	areabounds.BottomRight = p2
	return
}

func Day17P01() int {
	areabounds := parseDay17Input()
	maxy, _ := calc(areabounds)
	return maxy
}

func Day17P02() int {
	areabounds := parseDay17Input()
	_, cnt := calc(areabounds)
	return cnt
}

func landsInArea(velocity datastructures.Coordinate, areabounds datastructures.AreaBounds) (int, bool) {
	highest := 0
	pos := datastructures.Coordinate{X: 0, Y: 0}
	for {
		pos.X += velocity.X
		pos.Y += velocity.Y
		if pos.Y > highest {
			highest = pos.Y
		}
		if pos.InArea(areabounds) {
			return highest, true
		}

		if pos.Y < areabounds.TopLeft.Y || pos.X > areabounds.BottomRight.X {
			break
		}

		if velocity.X > 0 {
			velocity.X--
		}
		if velocity.X < 0 {
			velocity.X++
		}
		velocity.Y--
	}
	return highest, false
}

func calc(areabounds datastructures.AreaBounds) (int, int) {

	var highhest int
	successfulVelocities := make(map[datastructures.Coordinate]bool)
	bound := 300 // guesstimate, looking at the max pos of the area. Too lazy to accurately calculate
	for velX := bound * -1; velX < bound; velX++ {
		for velY := bound * -1; velY < bound; velY++ {
			velocity := datastructures.Coordinate{X: velX, Y: velY}
			if height, ok := landsInArea(velocity, areabounds); ok {
				if height > highhest {
					highhest = height
				}
				successfulVelocities[velocity] = true
			}
		}
	}
	return highhest, len(successfulVelocities)
}
