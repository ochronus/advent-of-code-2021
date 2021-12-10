package solutions

import (
	"ochronus/aoc2021/utils"
	"sort"
)

var openClosePairs = map[rune]rune{
	')': '(',
	']': '[',
	'}': '{',
	'>': '<',
}

var part1ScoreMap = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

var part2ScoreMap = map[rune]int{
	'(': 1,
	'[': 2,
	'{': 3,
	'<': 4,
}

func processLine(line string) (stack []rune, score int, corrupted bool) {
	for _, char := range line {
		if char == '(' || char == '[' || char == '{' || char == '<' {
			stack = append(stack, char)

		} else {
			if openClosePairs[char] != stack[len(stack)-1] { // the line is corrupted
				corrupted = true
				score += part1ScoreMap[char]
				return
			}

			stack = stack[:len(stack)-1] // stack pop
		}
	}
	return
}

func Day10P01() int {
	lines := utils.ReadFileLines("../inputs/10.txt")
	score := 0
	for _, line := range lines {
		_, lScore, _ := processLine(line)
		score += lScore
	}
	return score
}

func Day10P02() int {
	lines := utils.ReadFileLines("../inputs/10.txt")
	var scores []int
	for _, line := range lines {
		stack, _, corrupted := processLine(line)
		if !corrupted {
			total := 0
			for i := len(stack) - 1; i >= 0; i-- {
				total = total*5 + part2ScoreMap[stack[i]]
			}
			scores = append(scores, total)
		}
	}
	sort.Ints(scores)
	return scores[len(scores)/2]
}
