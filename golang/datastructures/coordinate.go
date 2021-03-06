package datastructures

import "ochronus/aoc2021/utils"

type Coordinate struct {
	X int
	Y int
}

func (p1 Coordinate) ManhattanDistance(p2 Coordinate) int {
	return utils.Abs(p2.X-p1.X) + utils.Abs(p2.Y-p1.Y)
}

func (c Coordinate) InArea(a AreaBounds) bool {
	return c.X >= a.TopLeft.X && c.X <= a.BottomRight.X && c.Y >= a.TopLeft.Y && c.Y <= a.BottomRight.Y
}

var Neighbors2D = [4]Coordinate{
	{0, -1},
	{0, 1},
	{-1, 0},
	{1, 0},
}

type AreaBounds struct {
	TopLeft     Coordinate
	BottomRight Coordinate
}
