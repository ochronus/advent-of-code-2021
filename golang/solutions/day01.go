package solutions

import (
	"bufio"
	"ochronus/aoc2021/utils"
	"os"
)

func getInput() []int {
	file, _ := os.Open("../inputs/01.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var lines []int
	for scanner.Scan() {
		lines = append(lines, utils.StrToInt(scanner.Text()))
	}
	return lines
}

func Day01P01() int {
	lines := getInput()
	counter := 0
	for index, i := range lines {
		if index > 0 && i > lines[index-1] {
			counter = counter + 1
		}
	}
	return counter
}

func Day01P02() int {
	lines := getInput()
	counter := 0
	for index, i := range lines {
		if index > 2 && i > lines[index-3] {
			counter = counter + 1
		}
	}
	return counter
}
