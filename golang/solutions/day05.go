package solutions

import (
	"ochronus/aoc2021/utils"
	"strings"
)

type Coordinate struct {
	X int
	Y int
}

func parseCoords(line string) (Coordinate, Coordinate) {
	s := strings.Split(line, " -> ")
	c1 := strings.Split(s[0], ",")
	c2 := strings.Split(s[1], ",")
	x1 := utils.StrToInt(c1[0])
	y1 := utils.StrToInt(c1[1])
	x2 := utils.StrToInt(c2[0])
	y2 := utils.StrToInt(c2[1])
	return Coordinate{X: x1, Y: y1}, Coordinate{X: x2, Y: y2}
}

func getUnitDirection(a int, b int) int {
	if b-a > 0 {
		return 1
	}
	if b-a < 0 {
		return -1
	}
	return 0
}

func solve(part2 bool) int {
	linemap := map[Coordinate]int{}
	for _, line := range utils.ReadFileLines("../inputs/05.txt") {
		start, end := parseCoords(line)
		if start.X == end.X || start.Y == end.Y || part2 {
			xIncrement := getUnitDirection(start.X, end.X)
			yIncrement := getUnitDirection(start.Y, end.Y)
			curPoint := start
			for ok := true; ok; ok = (curPoint.X != end.X+xIncrement || curPoint.Y != end.Y+yIncrement) {
				_, ok := linemap[curPoint]
				if ok {
					linemap[curPoint]++
				} else {
					linemap[curPoint] = 1
				}
				curPoint.X += xIncrement
				curPoint.Y += yIncrement
			}
		}
	}
	counter := 0
	for _, val := range linemap {
		if val > 1 {
			counter++
		}
	}
	return counter
}

func Day05P01() int {
	return solve(false)
}

func Day05P02() int {
	return solve(true)
}
