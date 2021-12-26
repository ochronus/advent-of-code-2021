package solutions

import (
	"ochronus/aoc2021/utils"
	"strings"
	"sync"
)

func Day25() int {
	parseDay25Input()
	step := 0
	run := true
	for run {
		run = processEastBunchInParallel()
		run = processSouthBunchInParallel() || run
		step++
	}
	return step
}

var (
	cucumberEast   map[int]map[int]bool
	cucumberSouth  map[int]map[int]bool
	maxCol, maxRow int
)

func parseDay25Input() {
	lines := utils.ReadFileLines("../inputs/25.txt")
	cucumberSouth = map[int]map[int]bool{}
	cucumberEast = map[int]map[int]bool{}
	row := 0
	for _, line := range lines {
		arr := strings.Split(line, "")
		maxCol = len(arr)
		for col, v := range arr {
			if v == "." {
				continue
			}

			if v == ">" {
				if cucumberEast[row] == nil {
					cucumberEast[row] = map[int]bool{}
				}
				cucumberEast[row][col] = true
				continue
			}

			if cucumberSouth[col] == nil {
				cucumberSouth[col] = map[int]bool{}
			}
			cucumberSouth[col][row] = true
		}
		row++
	}
	maxRow = row
}

func processSouthBunchInParallel() bool {
	var hasMove bool
	var wg sync.WaitGroup
	for i := 0; i < maxCol; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			doCycle := !cucumberSouth[i][0] && !cucumberEast[0][i] && cucumberSouth[i][maxRow-1]
			for j := 1; j < maxRow; j++ {
				if !cucumberSouth[i][j] && !cucumberEast[j][i] && cucumberSouth[i][j-1] {
					hasMove = true
					cucumberSouth[i][j] = true
					cucumberSouth[i][j-1] = false
					j++

				}
			}
			if doCycle {
				hasMove = true
				cucumberSouth[i][0] = true
				cucumberSouth[i][maxRow-1] = false
			}
		}(i)
	}
	wg.Wait()
	return hasMove
}

func processEastBunchInParallel() bool {
	var hasMove bool
	var wg sync.WaitGroup
	for i := 0; i < maxRow; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			doCycle := !cucumberEast[i][0] && !cucumberSouth[0][i] && cucumberEast[i][maxCol-1]
			for j := 1; j < maxCol; j++ {
				if !cucumberEast[i][j] && !cucumberSouth[j][i] && cucumberEast[i][j-1] {
					hasMove = true
					cucumberEast[i][j] = true
					cucumberEast[i][j-1] = false
					j++
				}
			}
			if doCycle {
				hasMove = true
				cucumberEast[i][0] = true
				cucumberEast[i][maxCol-1] = false
			}
		}(i)
	}
	wg.Wait()
	return hasMove
}
