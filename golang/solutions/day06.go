package solutions

import (
	"ochronus/aoc2021/utils"
)

func iter(fish []int, maxIter int) []int {
	for cnt := 0; cnt < maxIter; cnt++ {
		newGenCnt := 0
		for i, timer := range fish {
			if timer > 0 {
				fish[i]--
			} else {
				fish[i] = 6
				newGenCnt++
			}
		}
		for i := 0; i < newGenCnt; i++ {
			fish = append(fish, 8)
		}
	}
	return fish
}

func fastIter(fish []int, maxIter int) int {
	timerMap := map[int]int{}
	for _, timer := range fish {
		_, ok := timerMap[timer]
		if ok {
			timerMap[timer]++
		} else {
			timerMap[timer] = 1
		}
	}
	for cnt := 0; cnt < maxIter; cnt++ {
		newTimerMap := map[int]int{6: 0, 8: 0}
		for timer, count := range timerMap {
			if timer > 0 {
				_, ok := newTimerMap[timer-1]
				if ok {
					newTimerMap[timer-1] += count
				} else {
					newTimerMap[timer-1] = count
				}
			} else {
				newTimerMap[6] += count
				newTimerMap[8] += count
			}
		}
		timerMap = newTimerMap
	}
	totalFish := 0
	for _, count := range timerMap {
		totalFish += count
	}
	return totalFish
}

func Day06P01() int {
	fish := utils.ReadFileToIntList("../inputs/06.txt")
	return len(iter(fish, 80))
}

func Day06P02() int {
	fish := utils.ReadFileToIntList("../inputs/06.txt")
	return fastIter(fish, 256)
}
