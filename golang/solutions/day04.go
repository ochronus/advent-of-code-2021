package solutions

import (
	"fmt"
	"ochronus/aoc2021/utils"
	"strings"
)

type bingoSquare struct {
	row int
	col int
}

type bingoBoard struct {
	values map[string]bingoSquare
	marks  map[int]map[int]bool
}

func createBoard() bingoBoard {
	return bingoBoard{
		values: map[string]bingoSquare{},
		marks: map[int]map[int]bool{
			0: {},
			1: {},
			2: {},
			3: {},
			4: {},
		},
	}
}

func parseInput(lines []string) ([]string, []bingoBoard) {
	var boards []bingoBoard

	numberSequence := strings.Split(lines[0], ",")

	col := 0
	currentBoard := createBoard()
	for i := 2; i < len(lines); i++ {
		if lines[i] == "" { // empty line after each bingo board
			boards = append(boards, currentBoard)
			currentBoard = createBoard()
			col = 0

			continue
		}

		for row, value := range strings.Fields(lines[i]) {
			currentBoard.values[value] = bingoSquare{row, col}
		}

		col++
	}
	boards = append(boards, currentBoard)
	return numberSequence, boards
}

func (board bingoBoard) isWinning() bool {
	// check for fully marked rows or columns
	for row := 0; row < 5; row++ {
		markColCount := 0
		markRowCount := 0
		for col := 0; col < 5; col++ {
			if board.marks[row][col] {
				markColCount++
			}
			if board.marks[col][row] {
				markRowCount++
			}
		}
		if markColCount == 5 || markRowCount == 5 { // found a fully marked row or column
			return true
		}
	}

	return false
}

func (board bingoBoard) getScore(called int) int {
	unmarked := 0

	for value, square := range board.values {
		if !board.marks[square.row][square.col] {
			unmarked += utils.StrToInt(value)
		}
	}

	return unmarked * called
}

func Day04P01() int {
	numberSequence, boards := parseInput(utils.ReadFileLines("../inputs/04.txt"))

	for drawCounter, numberDrawn := range numberSequence {
		for _, board := range boards {
			if square, ok := board.values[numberDrawn]; ok {
				board.marks[square.row][square.col] = true

				if drawCounter > 3 && board.isWinning() {
					return board.getScore(utils.StrToInt(numberDrawn))
				}
			}
		}
	}

	fmt.Println("Didn't find a winning board")
	return 0
}

func Day04P02() int {
	numberSequence, boards := parseInput(utils.ReadFileLines("../inputs/04.txt"))

	for _, numberDrawn := range numberSequence {
		remainingBoards := make([]bingoBoard, 0, len(boards))

		for _, board := range boards {
			if square, ok := board.values[numberDrawn]; ok {
				board.marks[square.row][square.col] = true

				if !board.isWinning() {
					remainingBoards = append(remainingBoards, board)
				} else if len(boards) == 1 {
					return board.getScore(utils.StrToInt(numberDrawn))
				}
			} else {
				remainingBoards = append(remainingBoards, board)
			}
		}

		boards = remainingBoards
	}

	fmt.Println("Didn't find a winning board")
	return 0
}
