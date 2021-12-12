package solutions

import (
	"ochronus/aoc2021/utils"
	"strings"
	"unicode"
)

type adjacencyMap map[string][]string
type visitMap map[string]bool

func parseDay12Input() (adj adjacencyMap) {
	lines := utils.ReadFileLines("../inputs/12.txt")
	adj = make(adjacencyMap)
	for _, line := range lines {
		gr := strings.Split(line, "-")
		adj[gr[0]] = append(adj[gr[0]], gr[1])
		adj[gr[1]] = append(adj[gr[1]], gr[0])
	}
	return
}

func countPaths(part2 bool, adj adjacencyMap, currentCave string, alreadyVisited visitMap, visitedTwice bool) int {
	if currentCave == "end" {
		return 1
	}
	if part2 && currentCave == "start" && len(alreadyVisited) > 0 {
		return 0
	}
	_, isInSeen := alreadyVisited[currentCave]
	if unicode.IsLower(rune(currentCave[0])) && isInSeen {
		if part2 && !visitedTwice {
			visitedTwice = true // in part 2 we can visit small caves exactly twice
		} else { // in part 1 we can only visit small caves once
			return 0
		}
	}

	alreadyVisitedCopy := map[string]bool{} // prevent sub-paths from modifying the "root" visit log
	for k := range alreadyVisited {
		alreadyVisitedCopy[k] = true
	}
	alreadyVisitedCopy[currentCave] = true
	subPathCount := 0
	for _, cave := range adj[currentCave] {
		subPathCount += countPaths(part2, adj, cave, alreadyVisitedCopy, visitedTwice)
	}
	return subPathCount
}

func Day12P01() int {
	adj := parseDay12Input()
	return countPaths(false, adj, "start", visitMap{}, false)
}

func Day12P02() int {
	adj := parseDay12Input()
	return countPaths(true, adj, "start", visitMap{}, false)
}
