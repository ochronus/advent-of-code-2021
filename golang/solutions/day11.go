package solutions

import (
	"fmt"
	"ochronus/aoc2021/datastructures"
	"ochronus/aoc2021/utils"
	"strings"
	"time"

	"github.com/gookit/color"
)

type intGrid map[datastructures.Coordinate]int
type boolGrid map[datastructures.Coordinate]bool

var neighborOctopuses = []datastructures.Coordinate{
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
			grid[datastructures.Coordinate{x, y}] = level
		}
	}
	return
}

func propagateFlash(grid intGrid, flashMap boolGrid, octopus datastructures.Coordinate) {
	flashMap[octopus] = true

	for _, neighborDiff := range neighborOctopuses {
		neighbor := datastructures.Coordinate{octopus.X + neighborDiff.X, octopus.Y + neighborDiff.Y}
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

func drawGrid(grid intGrid) {

	time.Sleep(100 * time.Millisecond)
	fmt.Print("\033[H\033[2J")
	lines := [10][10]int{}
	for octopus, level := range grid {
		lines[octopus.X][octopus.Y] = level
	}
	for _, line := range lines {
		for _, level := range line {
			c := 25 + 255/10*level
			color.Printf("<bg=%d,%d,%d>🐙</>", c, c, c)
		}
		fmt.Println("")
	}
	fmt.Println("\n")
}

func simulate(grid intGrid, part2 bool) (int, int) {
	flashCount := 0
	for step := 1; ; step++ {
		if !part2 {
			drawGrid(grid)
		}
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
