package solutions

import (
	"fmt"
	"ochronus/aoc2021/utils"
	"regexp"
	"sort"
)

var stepRegexp = regexp.MustCompile(`(on|off) x=([\d-]+)..([\d-]+),y=([\d-]+)..([\d-]+),z=([\d-]+)..([\d-]+)`)

type Step struct {
	status                             bool
	xmin, xmax, ymin, ymax, zmin, zmax int
}

func parseStep(input string) Step {
	matches := stepRegexp.FindStringSubmatch(input)
	if len(matches) != 8 {
		panic(fmt.Errorf("bad input: %v", input))
	}
	return Step{
		status: matches[1] == "on",
		xmin:   utils.StrToInt(matches[2]),
		xmax:   utils.StrToInt(matches[3]),
		ymin:   utils.StrToInt(matches[4]),
		ymax:   utils.StrToInt(matches[5]),
		zmin:   utils.StrToInt(matches[6]),
		zmax:   utils.StrToInt(matches[7]),
	}
}

func parseDay22Input(input []string, filter func(*Step) bool) (steps []Step) {
	steps = []Step{}
	for _, line := range input {
		step := parseStep(line)
		if !filter(&step) {
			continue
		}
		steps = append(steps, step)
	}
	return
}

func unique(slice []int) (result []int) {
	set := make(map[int]bool)
	for _, n := range slice {
		set[n] = true
	}
	result = make([]int, 0)
	for k := range set {
		result = append(result, k)
	}
	return
}

func reactor(steps []Step) int {
	xaxis := make([]int, 0)
	yaxis := make([]int, 0)
	zaxis := make([]int, 0)

	for _, step := range steps {
		xaxis = append(xaxis, step.xmin, step.xmax+1)
		yaxis = append(yaxis, step.ymin, step.ymax+1)
		zaxis = append(zaxis, step.zmin, step.zmax+1)
	}

	xaxis = unique(xaxis)
	yaxis = unique(yaxis)
	zaxis = unique(zaxis)
	sort.Ints(xaxis)
	sort.Ints(yaxis)
	sort.Ints(zaxis)

	xmap := make(map[int]int)
	for i, x := range xaxis {
		xmap[x] = i
	}

	ymap := make(map[int]int)
	for j, y := range yaxis {
		ymap[y] = j
	}

	zmap := make(map[int]int)
	for k, x := range zaxis {
		zmap[x] = k
	}

	grid := make([][][]bool, len(xaxis))
	for x := 0; x < len(xaxis); x++ {
		grid[x] = make([][]bool, len(yaxis))
		for y := 0; y < len(yaxis); y++ {
			grid[x][y] = make([]bool, len(zaxis))
			for z := 0; z < len(zaxis); z++ {
				grid[x][y][z] = false
			}
		}
	}

	for _, step := range steps {
		xstart := xmap[step.xmin]
		xend := xmap[step.xmax+1] - 1
		ystart := ymap[step.ymin]
		yend := ymap[step.ymax+1] - 1
		zstart := zmap[step.zmin]
		zend := zmap[step.zmax+1] - 1

		for x := xstart; x <= xend; x++ {
			for y := ystart; y <= yend; y++ {
				for z := zstart; z <= zend; z++ {
					grid[x][y][z] = step.status
				}
			}
		}
	}

	enabled := 0
	for x := 0; x < len(xaxis)-1; x++ {
		for y := 0; y < len(yaxis)-1; y++ {
			for z := 0; z < len(zaxis)-1; z++ {
				if grid[x][y][z] {
					enabled += (xaxis[x+1] - xaxis[x]) * (yaxis[y+1] - yaxis[y]) * (zaxis[z+1] - zaxis[z])
				}
			}
		}
	}

	return enabled
}

func Day22P01() int {
	input := utils.ReadFileLines("../inputs/22.txt")
	return reactor(parseDay22Input(input, func(step *Step) bool {
		if 50 < step.xmax || step.xmin < -50 ||
			50 < step.ymax || step.ymin < -50 ||
			50 < step.zmax || step.zmin < -50 {
			return false
		}
		return true
	}))
}

func Day22P02() int {
	input := utils.ReadFileLines("../inputs/22.txt")
	return reactor(parseDay22Input(input, func(step *Step) bool {
		return true
	}))
}
