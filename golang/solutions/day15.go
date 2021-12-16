package solutions

import (
	"ochronus/aoc2021/datastructures"
	"ochronus/aoc2021/utils"

	"github.com/beefsack/go-astar"
)

type CavePosition struct {
	cave *datastructures.Cave
	loc  datastructures.Coordinate
}

func (s CavePosition) PathNeighbors() (results []astar.Pather) {
	for _, delta := range datastructures.Neighbors2D {
		newLoc := datastructures.Coordinate{X: s.loc.X + delta.X, Y: s.loc.Y + delta.Y}
		_, ok := s.cave.TileAt(newLoc)
		if !ok {
			continue
		}

		results = append(results, CavePosition{cave: s.cave, loc: newLoc})
	}

	return
}

func (s CavePosition) PathNeighborCost(to astar.Pather) float64 {
	toCavePosition := to.(CavePosition)
	v, _ := s.cave.TileValueAt(toCavePosition.loc)

	return float64(v)
}

func (s CavePosition) PathEstimatedCost(to astar.Pather) float64 {
	other := to.(CavePosition)
	return float64(s.loc.ManhattanDistance(other.loc))
}

var _ astar.Pather = CavePosition{}

func parseDay15Input() (cave *datastructures.Cave) {
	lines := utils.ReadFileLines("../inputs/15.txt")
	cave = datastructures.NewCave(len(lines[0]), len(lines))

	for row, line := range lines {
		for column, tile := range line {
			cave.SetTile(datastructures.Coordinate{X: column, Y: row}, tile)
		}
	}
	return
}

func Day15P01() int {
	cave := parseDay15Input()

	_, distance, found := astar.Path(CavePosition{cave: cave, loc: cave.TopLeft()}, CavePosition{cave: cave, loc: cave.BottomRight()})
	if !found {
		panic("Astar couldn't find a solution")
	}

	return int(distance)
}

func Day15P02() int {
	cave := parseDay15Input()
	cave = cave.Extend(5)

	_, distance, found := astar.Path(CavePosition{cave: cave, loc: cave.TopLeft()}, CavePosition{cave: cave, loc: cave.BottomRight()})
	if !found {
		panic("Astar couldn't find a solution")
	}

	return int(distance)
}
