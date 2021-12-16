package datastructures

import (
	"fmt"
	"ochronus/aoc2021/utils"
	"strconv"
)

type Cave struct {
	tiles  []rune
	width  int
	height int
}

func NewCave(w, h int) *Cave {
	return &Cave{
		tiles:  make([]rune, w*h),
		width:  w,
		height: h,
	}
}

func (cave Cave) Size() (int, int) {
	return cave.width, cave.height
}

func (cave Cave) TopLeft() Coordinate {
	return Coordinate{X: 0, Y: 0}
}

func (cave Cave) BottomRight() Coordinate {
	return Coordinate{X: cave.width - 1, Y: cave.height - 1}
}

func (cave *Cave) TileAt(loc Coordinate) (rune, bool) {
	idx, ok := cave.indexOf(loc)
	if !ok {
		return ' ', false
	}

	return cave.tiles[idx], true
}

func (cave *Cave) TileValueAt(loc Coordinate) (int, bool) {
	idx, ok := cave.indexOf(loc)
	if !ok {
		return ' ', false
	}

	return utils.StrToInt(string(cave.tiles[idx])), true
}

func (cave *Cave) isOutOfBounds(loc Coordinate) bool {
	return loc.X < 0 || loc.Y < 0 || loc.X >= cave.width || loc.Y >= cave.height
}

func (cave *Cave) SetTile(loc Coordinate, tile rune) {
	idx, ok := cave.indexOf(loc)
	if !ok {
		panic(fmt.Errorf("out of bounds tile access: [%v]", loc))
	}

	cave.tiles[idx] = tile
}

func (cave *Cave) indexOf(loc Coordinate) (int, bool) {
	return loc.X + (cave.width * loc.Y), !cave.isOutOfBounds(loc)
}

func (cave *Cave) Extend(times int) *Cave {
	w, h := cave.Size()

	result := NewCave(w*times, h*times)

	for XRepeat := 0; XRepeat < times; XRepeat++ {
		for YRepeat := 0; YRepeat < times; YRepeat++ {
			for x := 0; x < w; x++ {
				for y := 0; y < h; y++ {
					tileValue, _ := cave.TileValueAt(Coordinate{X: x, Y: y})

					var prevTileValue int
					var ok bool
					if XRepeat > 0 {
						tmpLoc := Coordinate{X: w*(XRepeat-1) + x, Y: w*YRepeat + y}
						prevTileValue, ok = result.TileValueAt(tmpLoc)
					} else if YRepeat > 0 {
						tmpLoc := Coordinate{X: w*XRepeat + x, Y: w*(YRepeat-1) + y}
						prevTileValue, ok = result.TileValueAt(tmpLoc)
					}

					if ok {
						tileValue = prevTileValue + 1

						if tileValue > 9 {
							tileValue = 1
						}
					}

					n := rune(strconv.Itoa(tileValue)[0])
					newLoc := Coordinate{X: w*XRepeat + x, Y: h*YRepeat + y}
					result.SetTile(newLoc, n)
				}
			}
		}
	}

	return result
}
