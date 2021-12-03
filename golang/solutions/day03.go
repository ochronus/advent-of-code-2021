package solutions

import (
	"bufio"
	"os"
	"strconv"
)

func getDay3Input() []string {
	file, _ := os.Open("../inputs/03.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func Day03P01() int64 {
	lines := getDay3Input()
	ones := make([]int, len(lines[0]))
	majorityLimit := len(lines) / 2
	gammaBits := ""
	epsilonBits := ""

	for _, line := range lines {
		for bitPos, bit := range line {
			if bit == '1' {
				ones[bitPos]++
			}
		}
	}

	for _, onesCount := range ones {
		if onesCount > majorityLimit {
			gammaBits += "1"
			epsilonBits += "0"
		} else {
			gammaBits += "0"
			epsilonBits += "1"
		}
	}
	gamma, _ := strconv.ParseInt(gammaBits, 2, 64)
	epsilon, _ := strconv.ParseInt(epsilonBits, 2, 64)
	return gamma * epsilon

}

func countBits(pos int, lines []string) (ones int, zeros int) {
	for _, line := range lines {
		if line[pos] == '1' {
			ones++
		} else {
			zeros++
		}
	}
	return
}

func filter(lines []string, oxy bool) int64 {
	filtered := lines

	for pos := 0; pos < len(lines[0]); pos++ {
		ones, zeros := countBits(pos, filtered)
		var tmp []string
		for _, line := range filtered {
			if oxy {
				if ones >= zeros {
					if line[pos] == '1' {
						tmp = append(tmp, line)
					}
				} else {
					if line[pos] == '0' {
						tmp = append(tmp, line)
					}
				}
			} else {
				if ones < zeros {
					if line[pos] == '1' {
						tmp = append(tmp, line)
					}
				} else {
					if line[pos] == '0' {
						tmp = append(tmp, line)
					}
				}
			}
		}
		filtered = tmp
		if len(filtered) == 1 {
			break
		}
	}
	val, _ := strconv.ParseInt(filtered[0], 2, 64)
	return val
}

func Day03P02() int64 {
	lines := getDay3Input()

	oxygen := filter(lines, true)
	co2 := filter(lines, false)

	return oxygen * co2
}
