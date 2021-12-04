package solutions

import (
	"bufio"
	"ochronus/aoc2021/utils"
	"os"
	"strings"
)

type instr struct {
	Direction string
	Amount    int
}

func getDay2Input() []instr {
	file, _ := os.Open("../inputs/02.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var lines []instr
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		lines = append(lines, instr{
			Direction: parts[0],
			Amount:    utils.StrToInt(parts[1]),
		})
	}
	return lines
}

func Day02P01() int {
	lines := getDay2Input()
	var depth, horizontal int
	for _, i := range lines {
		switch i.Direction {
		case "forward":
			horizontal += i.Amount
		case "down":
			depth += i.Amount
		case "up":
			depth -= i.Amount
		}
	}
	return depth * horizontal
}

func Day02P02() int {
	lines := getDay2Input()
	var depth, horizontal, aim int

	for _, i := range lines {
		switch i.Direction {
		case "forward":
			horizontal += i.Amount
			depth += i.Amount * aim
		case "down":
			aim += i.Amount
		case "up":
			aim -= i.Amount
		}
	}
	return depth * horizontal
}
