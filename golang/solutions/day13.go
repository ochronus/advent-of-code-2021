package solutions

import (
	"fmt"
	"ochronus/aoc2021/datastructures"
	"ochronus/aoc2021/utils"
	"strings"
)

type foldInstruction struct {
	axis     rune
	foldLine int
}

func parseDay13Input() (grid boolGrid, instructions []foldInstruction) {
	input := strings.Split(utils.ReadFileToString("../inputs/13.txt"), "\n\n")
	grid = make(boolGrid)

	for _, line := range strings.Split(input[0], "\n") {
		dots := utils.StrToIntList(line, ",")
		grid[datastructures.Coordinate{X: dots[0], Y: dots[1]}] = true
	}
	for _, line := range strings.Split(input[1], "\n") {
		p := strings.Split(line, "=")
		instructions = append(instructions, foldInstruction{
			axis:     rune(p[0][len(p[0])-1]),
			foldLine: utils.StrToInt(p[1]),
		})
	}
	return
}

func fold(grid boolGrid, instruction foldInstruction) (foldedGrid boolGrid) {
	foldedGrid = make(boolGrid)
	if instruction.axis == 'x' {
		for leftHalfX := 0; leftHalfX < instruction.foldLine; leftHalfX++ {
			for dot := range grid {
				if dot.X == leftHalfX || dot.X == 2*instruction.foldLine-leftHalfX {
					foldedGrid[datastructures.Coordinate{X: leftHalfX, Y: dot.Y}] = true
				}
			}
		}
	}
	if instruction.axis == 'y' {
		for topHalfY := 0; topHalfY < instruction.foldLine; topHalfY++ {
			for dot := range grid {
				if dot.Y == topHalfY || dot.Y == 2*instruction.foldLine-topHalfY {
					foldedGrid[datastructures.Coordinate{X: dot.X, Y: topHalfY}] = true
				}
			}
		}
	}
	return
}

func displayGrid(grid boolGrid) {
	topLeft := datastructures.Coordinate{X: 9999, Y: 9999}
	bottomRight := datastructures.Coordinate{X: 0, Y: 0}
	for c := range grid {
		topLeft = datastructures.Coordinate{X: utils.Min(topLeft.X, c.X), Y: utils.Min(topLeft.Y, c.Y)}
		bottomRight = datastructures.Coordinate{X: utils.Max(bottomRight.X, c.X), Y: utils.Max(bottomRight.Y, c.Y)}
	}
	for y := topLeft.Y; y <= bottomRight.Y; y++ {
		for x := topLeft.X; x <= bottomRight.X; x++ {
			if grid[datastructures.Coordinate{X: x, Y: y}] {
				fmt.Print("##")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
}

func Day13P01() int {
	grid, instr := parseDay13Input()
	grid = fold(grid, instr[0])
	return len(grid)
}

func Day13P02() int {
	grid, instr := parseDay13Input()
	for _, i := range instr {
		grid = fold(grid, i)
	}
	displayGrid(grid)
	return 0
}
