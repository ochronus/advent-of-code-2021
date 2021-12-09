package solutions

import (
	"ochronus/aoc2021/utils"
	"sort"
)

type coordinate struct {
	x int
	y int
}
type heightmap map[coordinate]int
type boolmap map[coordinate]bool

var neighbors = [4]coordinate{
	{0, -1},
	{0, 1},
	{-1, 0},
	{1, 0},
}

func parseDay9Input() heightmap {
	lines := utils.ReadFileLines("../inputs/09.txt")
	grid := make(heightmap)
	for x, line := range lines {
		for y, height := range utils.StrToIntList(line) {
			grid[coordinate{x, y}] = height
		}
	}
	return grid
}

func findLowPoints(grid heightmap) (sum int, lowpoints []coordinate) {
	for pos, height := range grid {
		lowest := 9
		for _, neighbor := range neighbors {
			if nHeight, ok := grid[coordinate{
				x: pos.x + neighbor.x,
				y: pos.y + neighbor.y,
			}]; ok {
				lowest = utils.Min(lowest, nHeight)
			}
		}

		if height < lowest {
			sum += height + 1
			lowpoints = append(lowpoints, pos)
		}
	}
	return
}

func Day09P01() int {
	grid := parseDay9Input()
	sum, _ := findLowPoints(grid)

	return sum
}

func Day09P02() int {
	grid := parseDay9Input()

	_, lowpoints := findLowPoints(grid)

	var basinSizes []int
	seen := make(boolmap)
	for _, lowpoint := range lowpoints {
		basinSizes = append(basinSizes, getBasinSize(grid, lowpoint, lowpoint, seen))
	}
	sort.Ints(basinSizes)
	lastindex := len(basinSizes) - 1
	return basinSizes[lastindex] * basinSizes[lastindex-1] * basinSizes[lastindex-2]
}

func getBasinSize(grid map[coordinate]int, pos coordinate, start coordinate, seen boolmap) (sum int) {
	for _, neighbor := range neighbors {
		basinCandidate := coordinate{
			x: pos.x + neighbor.x,
			y: pos.y + neighbor.y,
		}
		if _, ok := seen[basinCandidate]; !ok {
			seen[basinCandidate] = true
			height := grid[basinCandidate]
			if height < 9 && height > grid[start] {
				sum += getBasinSize(grid, basinCandidate, start, seen)
			}
		}
	}

	sum += 1
	return
}
