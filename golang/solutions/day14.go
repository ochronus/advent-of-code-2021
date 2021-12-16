package solutions

import (
	"ochronus/aoc2021/utils"
	"strings"
)

type polymerMap map[string]string
type pairCount map[string]int
type freqMap map[rune]int

func parseDay14Input() (template string, rules polymerMap) {
	input := strings.Split(utils.ReadFileToString("../inputs/14.txt"), "\n\n")

	rules = make(polymerMap)
	template = input[0]
	for _, line := range strings.Split(input[1], "\n") {
		rule := strings.Split(line, " -> ")
		rules[rule[0]] = rule[1]
	}

	return
}

func growPolymer(template string, rules polymerMap, iterations int) int {
	pairCounter := make(pairCount)

	for i := 0; i < len(template)-1; i++ {
		pairCounter[template[i:i+2]] += 1
	}
	pairCounter[template[len(template)-1:]] += 1 // damn you, off-by-1 error!

	for i := 0; i < iterations; i++ {
		tmpPairCounter := make(pairCount) // we need a copy so insertions are simultaneous

		for pair, cnt := range pairCounter {
			if insertion, ok := rules[pair]; ok {
				// AB --> AXB --> AX, XB
				newPair1 := pair[0:1] + insertion // AX
				newPair2 := insertion + pair[1:2] // XB
				tmpPairCounter[newPair1] += cnt
				tmpPairCounter[newPair2] += cnt
			} else {
				tmpPairCounter[pair] = cnt // preserve intact pairs' counters
			}
		}
		pairCounter = tmpPairCounter
	}

	polymerFrequencies := make(freqMap)
	minFreq := 0
	maxFreq := 0
	for pair, cnt := range pairCounter {
		// since pairs are formed using a sliding window:
		// ABACDCE --> AB, BA, AC, CD, DC, CE, E (remember the off-by-1 error above?)
		// we only need to look at the first polymer in each pair to avoid duplication
		polymerFrequencies[rune(pair[0])] += cnt
		minFreq = polymerFrequencies[rune(pair[0])]
		maxFreq = polymerFrequencies[rune(pair[0])]
	}

	for _, freq := range polymerFrequencies {
		if freq < minFreq {
			minFreq = freq
		}
		if freq > maxFreq {
			maxFreq = freq
		}
	}
	return maxFreq - minFreq
}

func Day14P01() int {
	template, rules := parseDay14Input()
	return growPolymer(template, rules, 10)
}

func Day14P02() int {
	template, rules := parseDay14Input()
	return growPolymer(template, rules, 40)
}
