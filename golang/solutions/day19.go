package solutions

import (
	"ochronus/aoc2021/utils"
	"strings"
)

type point struct {
	x, y, z int
}

type scanner map[point]bool

type absoluteMap map[point]bool

func InputToSlice(input string) ([]string, error) {
	return strings.Split(strings.TrimSuffix(input, "\n"), "\n"), nil
}

func parseDay19Input() []scanner {
	input := utils.ReadFileToString("../inputs/19.txt")
	s := strings.Split(input, "\n\n")
	scannerList := make([]scanner, 0)
	for i := range s {
		scan := make(scanner)
		lines, _ := InputToSlice(s[i])
		for l := range lines {
			if strings.HasPrefix(lines[l], "---") {
				continue
			}
			nums := strings.Split(lines[l], ",")

			x := utils.StrToInt(nums[0])
			y := utils.StrToInt(nums[1])
			z := utils.StrToInt(nums[2])
			scan[point{x: x, y: y, z: z}] = true
		}
		scannerList = append(scannerList, scan)
	}
	return scannerList
}

func (a absoluteMap) solve(s []scanner) ([]point, absoluteMap) {
	var scannerPoints []point
	abs := absoluteMap(s[0])
	for {
		var newScanner []scanner
		for scan := range s {
			x, y, z, sc := abs.findOffset(s[scan])
			for p := range sc {
				abs[point{
					x: p.x + x,
					y: p.y + y,
					z: p.z + z,
				}] = true
			}
			if len(sc) == 0 {
				newScanner = append(newScanner, s[scan])
				continue
			}
			scannerPoints = append(scannerPoints, point{x: x, y: y, z: z})
		}
		s = newScanner
		if len(s) == 0 {
			break
		}
	}
	return scannerPoints, a
}

func (s scanner) getOrientations() []scanner {
	orientations := make([]scanner, 24)
	for i := 0; i < 24; i++ {
		orientations[i] = make(scanner)
	}

	for p := range s {
		x, y, z := p.x, p.y, p.z
		orientations[0][point{x, y, z}] = true
		orientations[1][point{x, -z, y}] = true
		orientations[2][point{x, -y, -z}] = true
		orientations[3][point{x, z, -y}] = true
		orientations[4][point{-x, -y, z}] = true
		orientations[5][point{-x, -z, -y}] = true
		orientations[6][point{-x, y, -z}] = true
		orientations[7][point{-x, z, y}] = true
		orientations[8][point{-z, x, -y}] = true
		orientations[9][point{y, x, -z}] = true
		orientations[10][point{z, x, y}] = true
		orientations[11][point{-y, x, z}] = true
		orientations[12][point{z, -x, -y}] = true
		orientations[13][point{y, -x, z}] = true
		orientations[14][point{-z, -x, y}] = true
		orientations[15][point{-y, -x, -z}] = true
		orientations[16][point{-y, -z, x}] = true
		orientations[17][point{z, -y, x}] = true
		orientations[18][point{y, z, x}] = true
		orientations[19][point{-z, y, x}] = true
		orientations[20][point{z, y, -x}] = true
		orientations[21][point{-y, z, -x}] = true
		orientations[22][point{-z, -y, -x}] = true
		orientations[23][point{y, -z, -x}] = true
	}
	return orientations
}

func (a absoluteMap) findOffset(s scanner) (int, int, int, scanner) {
	orientations := s.getOrientations()
	points := make(map[point]int)
	for p := range orientations {
		for potentialPoint := range orientations[p] {
			for mapPoint := range a {
				points[point{
					x: mapPoint.x - potentialPoint.x,
					y: mapPoint.y - potentialPoint.y,
					z: mapPoint.z - potentialPoint.z,
				}]++
				if points[point{
					x: mapPoint.x - potentialPoint.x,
					y: mapPoint.y - potentialPoint.y,
					z: mapPoint.z - potentialPoint.z,
				}] > 11 {
					return mapPoint.x - potentialPoint.x, mapPoint.y - potentialPoint.y, mapPoint.z - potentialPoint.z, orientations[p]
				}
			}
		}
	}
	return 0, 0, 0, nil
}

func Day19P01() int {
	s := parseDay19Input()
	abs := absoluteMap(s[0])
	_, abs = abs.solve(s)

	return len(abs)
}

func Day19P02() int {
	s := parseDay19Input()
	abs := absoluteMap(s[0])
	scannerPoints, _ := abs.solve(s)
	var max int
	for i1 := range scannerPoints {
		for i2 := range scannerPoints {
			sum := utils.Abs(scannerPoints[i1].x-scannerPoints[i2].x) +
				utils.Abs(scannerPoints[i1].y-scannerPoints[i2].y) +
				utils.Abs(scannerPoints[i1].z-scannerPoints[i2].z)
			if max < int(sum) {
				max = int(sum)
			}
		}
	}
	return max
}
