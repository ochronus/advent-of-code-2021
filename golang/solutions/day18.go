package solutions

import (
	"ochronus/aoc2021/datastructures"
	"ochronus/aoc2021/utils"
)

func sum(numbers []*datastructures.SnailFishNumber) int {
	sum := &datastructures.SnailFishNumber{}
	for _, num := range numbers {
		sum = sum.Add(num)
	}
	return sum.Magnitude()
}

func maxSumMagnitude(numbers []*datastructures.SnailFishNumber) int {
	maxSum := 0
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers); j++ {
			if i != j {
				sum := numbers[i].MakeCopy().Add(numbers[j].MakeCopy()).Magnitude()
				if sum > maxSum {
					maxSum = sum
				}
			}
		}
	}
	return maxSum
}

func parseDay18Input() []*datastructures.SnailFishNumber {
	lines := utils.ReadFileLines("../inputs/18.txt")
	numbers := make([]*datastructures.SnailFishNumber, 0, 1)
	for _, line := range lines {
		number := &datastructures.SnailFishNumber{}
		for _, token := range line {
			switch token {
			case '[':
				number.Left = &datastructures.SnailFishNumber{Parent: number}
				number.Right = &datastructures.SnailFishNumber{Parent: number}
				number = number.Left
			case ',':
				number = number.Parent.Right
			case ']':
				number = number.Parent
			default:
				number.Value = utils.StrToInt(string(token))
			}
		}
		numbers = append(numbers, number)
	}

	return numbers
}

func Day18P01() int {
	numbers := parseDay18Input()
	return sum(numbers)
}

func Day18P02() int {
	numbers := parseDay18Input()
	return maxSumMagnitude(numbers)
}
