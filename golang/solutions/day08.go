package solutions

import (
	"math"
	"ochronus/aoc2021/utils"
	"strings"
)

type segmentSet struct {
	probes  []string
	outputs []string
}

func parseDay8Input() []segmentSet {
	lines := utils.ReadFileLines("../inputs/08.txt")
	var s []segmentSet
	for _, line := range lines {
		p := strings.Split(line, " | ")
		s = append(s, segmentSet{
			probes:  strings.Split(p[0], " "),
			outputs: strings.Split(p[1], " "),
		})
	}
	return s
}

func Day08P01() int {
	input := parseDay8Input()
	cnt := 0
	for _, s := range input {
		for _, output := range s.outputs {
			l := len(output)
			if l == 2 || l == 3 || l == 4 || l == 7 {
				cnt++
			}
		}
	}
	return cnt
}

func Day08P02() int {
	input := parseDay8Input()
	cnt := 0.0
	for _, s := range input {
		segments := [10]string{}
		for _, digit := range s.probes {
			l := len(digit)
			if l == 2 {
				segments[1] = digit
			} else if l == 3 {
				segments[7] = digit
			} else if l == 4 {
				segments[4] = digit
			} else if l == 7 {
				segments[8] = digit
			}
			for i := 0; i < 10; i++ {
				for _, digit := range s.probes {
					switch len(digit) {
					case 5:
						if utils.StringIntersection(digit, segments[1]) == 2 {
							segments[3] = digit
						}
						if utils.StringIntersection(digit, segments[7]) == 2 &&
							utils.StringIntersection(digit, segments[1]) == 1 &&
							utils.StringIntersection(digit, segments[3]) == 4 &&
							utils.StringIntersection(digit, segments[4]) == 2 {
							segments[2] = digit
						}
						if utils.StringIntersection(digit, segments[2]) == 3 &&
							utils.StringIntersection(digit, segments[4]) == 3 {
							segments[5] = digit
						}

					case 6:
						if utils.StringIntersection(digit, segments[1]) == 2 &&
							utils.StringIntersection(digit, segments[2]) == 4 &&
							utils.StringIntersection(digit, segments[3]) == 4 {
							segments[0] = digit
						}
						if utils.StringIntersection(digit, segments[7]) == 3 &&
							utils.StringIntersection(digit, segments[0]) == 5 {
							segments[9] = digit
						}
						if utils.StringIntersection(digit, segments[1]) == 1 &&
							utils.StringIntersection(digit, segments[2]) == 4 {
							segments[6] = digit
						}

					}
				}
			}

		}

		for k, v := range segments {
			segments[k] = utils.SortString(v)
		}

		tenExp := 3.0
		for _, output := range s.outputs {
			output = utils.SortString(output)

			for numericValue, digit := range segments {
				if output == digit {
					cnt += math.Pow(10, tenExp) * float64(numericValue)
					tenExp -= 1.0
				}
			}
		}

	}

	return int(cnt)
}
