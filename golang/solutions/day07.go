package solutions

import (
	"ochronus/aoc2021/utils"
	"sort"
)

func mean(i []int) int {
	total := 0
	for _, number := range i {
		total = total + number
	}
	return total / len(i)
}

func median(i []int) int {
	sort.Ints(i)

	l := len(i)
	median := 0
	if l == 0 {
		return 0
	} else if l%2 == 0 {
		median = mean(i[l/2-1 : l/2+1])
	} else {
		median = i[l/2]
	}

	return median
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func diffSums(i []int, position int, part2 bool) int {
	sum := 0
	for _, num := range i {
		d := abs(num - position)
		if part2 {
			sum += d * (d + 1) / 2 // the sum of natural numbers up to d
		} else {
			sum += abs(num - position)
		}
	}
	return sum
}

func Day07P01() int {
	positions := utils.ReadFileToIntList("../inputs/07.txt")
	median := median(positions)
	return diffSums(positions, median, false)
}

func Day07P02() int {
	positions := utils.ReadFileToIntList("../inputs/07.txt")
	median := mean(positions)
	return diffSums(positions, median, true)
}
