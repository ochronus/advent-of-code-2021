package solutions

import (
	"ochronus/aoc2021/datastructures"
	"ochronus/aoc2021/utils"
	"sort"
)

type heightmap map[datastructures.Coordinate]int

func parseDay9Input() heightmap {
	lines := utils.ReadFileLines("../inputs/09.txt")
	grid := make(heightmap)
	for x, line := range lines {
		for y, height := range utils.StrToIntList(line, "") {
			grid[datastructures.Coordinate{X: x, Y: y}] = height
		}
	}
	return grid
}

func findLowPoints(grid heightmap) (sum int, lowpoints []datastructures.Coordinate) {
	for pos, height := range grid {
		lowest := 9
		for _, neighbor := range datastructures.Neighbors2D {
			if nHeight, ok := grid[datastructures.Coordinate{
				X: pos.X + neighbor.X,
				Y: pos.Y + neighbor.Y,
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
	seen := make(datastructures.BoolMap)
	for _, lowpoint := range lowpoints {
		basinSizes = append(basinSizes, getBasinSize(grid, lowpoint, lowpoint, seen))
	}
	sort.Ints(basinSizes)
	lastindex := len(basinSizes) - 1
	return basinSizes[lastindex] * basinSizes[lastindex-1] * basinSizes[lastindex-2]
}

func getBasinSize(grid map[datastructures.Coordinate]int, pos datastructures.Coordinate, start datastructures.Coordinate, seen datastructures.BoolMap) (sum int) {
	for _, neighbor := range datastructures.Neighbors2D {
		basinCandidate := datastructures.Coordinate{
			X: pos.X + neighbor.X,
			Y: pos.Y + neighbor.Y,
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
