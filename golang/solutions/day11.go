package solutions

import (
	"ochronus/aoc2021/utils"
	"strings"
)

type intGrid map[Coordinate]int
type boolGrid map[Coordinate]bool

var neighborOctopuses = []Coordinate{
	{1, 0},
	{1, 1},
	{1, -1},
	{0, 1},
	{-1, 0},
	{-1, 1},
	{-1, -1},
	{0, -1},
}

func fileToGrid(fname string) (grid intGrid) {
	lines := utils.ReadFileLines(fname)
	grid = make(intGrid)
	for y, line := range lines {
		var nums []int
		for _, nStr := range strings.Split(line, "") {
			nums = append(nums, utils.StrToInt(nStr))
		}
		for x, level := range nums {
			grid[Coordinate{x, y}] = level
		}
	}
	return
}

func propagateFlash(grid intGrid, flashMap boolGrid, octopus Coordinate) {
	flashMap[octopus] = true

	for _, neighborDiff := range neighborOctopuses {
		neighbor := Coordinate{octopus.X + neighborDiff.X, octopus.Y + neighborDiff.Y}
		if _, ok := grid[neighbor]; ok {
			if flashMap[neighbor] {
				continue
			}
			grid[neighbor]++
			if grid[neighbor] > 9 {
				propagateFlash(grid, flashMap, neighbor)
			}
		}

	}
}

func simulate(grid intGrid, part2 bool) (int, int) {
	flashCount := 0
	for step := 1; ; step++ {
		flashMap := make(boolGrid)

		for octopus := range grid {
			grid[octopus]++
			if grid[octopus] > 9 && !flashMap[octopus] {
				propagateFlash(grid, flashMap, octopus)
			}
		}

		for octopus := range flashMap {
			flashCount++
			grid[octopus] = 0
		}

		if !part2 && step == 100 {
			return flashCount, 0
		}
		if part2 && len(flashMap) == len(grid) {
			return 0, step + 1
		}
	}
}

func Day11P01() int {
	grid := fileToGrid("../inputs/11.txt")
	res, _ := simulate(grid, false)
	return res
}

func Day11P02() int {
	grid := fileToGrid("../inputs/11.txt")
	_, res := simulate(grid, true)
	return res
}