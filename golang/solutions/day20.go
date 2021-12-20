package solutions

import (
	"math"
	"ochronus/aoc2021/datastructures"
	"ochronus/aoc2021/utils"
)

func parseDay20Input() (algo string, inputImage datastructures.BoolMap) {
	lines := utils.ReadFileLines("../inputs/20.txt")
	algo = lines[0]
	inputImage = make(datastructures.BoolMap)
	for x, line := range lines[2:] {
		for y, v := range line {
			inputImage[datastructures.Coordinate{x, y}] = v == '#'
		}
	}
	return
}

func Day20P01() int {
	algo, image := parseDay20Input()
	return enhance(algo, image, 2)
}

func Day20P02() int {
	algo, image := parseDay20Input()
	return enhance(algo, image, 50)
}

func getImageDimension(image datastructures.BoolMap) (topLeft, bottomRight datastructures.Coordinate) {
	var maxx, maxy int
	var minx, miny = math.MaxInt, math.MaxInt
	for c := range image {
		maxx = utils.Max(maxx, c.X)
		maxy = utils.Max(maxy, c.Y)
		minx = utils.Min(minx, c.X)
		miny = utils.Min(miny, c.Y)
	}
	topLeft = datastructures.Coordinate{minx, miny}
	bottomRight = datastructures.Coordinate{maxx, maxy}
	return
}

func enhance(algo string, image datastructures.BoolMap, iterCount int) int {
	algoFirst := rune(algo[0]) == '#'
	algoLast := rune(algo[len(algo)-1]) == '#'
	for i := 0; i < iterCount; i++ {
		outputImage := make(datastructures.BoolMap)
		topLeft, bottomRight := getImageDimension(image)

		for x := topLeft.X - 1; x <= bottomRight.X+1; x++ {
			for y := topLeft.Y - 1; y <= bottomRight.Y+1; y++ {
				var algoIndex int
				for _, neighborX := range []int{x - 1, x, x + 1} {
					for _, neighborY := range []int{y - 1, y, y + 1} {
						algoIndex = algoIndex << 1 // ads a 0 bit on the right
						if pixel, ok := image[datastructures.Coordinate{neighborX, neighborY}]; ok {
							if pixel {
								algoIndex |= 1 // sets the rightmost (just added) bit to 1
							}
						} else {
							if i%2 != 0 {
								if algoFirst {
									algoIndex |= 1
								}

							} else {
								if algoLast {
									algoIndex |= 1
								}
							}
						}
					}
				}
				outputImage[datastructures.Coordinate{x, y}] = algo[algoIndex] == '#'
			}
		}

		image = outputImage
	}
	return countLitPixels(image)
}

func countLitPixels(image datastructures.BoolMap) (res int) {
	for _, pixel := range image {
		if pixel {
			res++
		}
	}
	return
}
