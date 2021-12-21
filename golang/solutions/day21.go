package solutions

import (
	"ochronus/aoc2021/utils"
)

func Day21P01() int {
	p1 := 5
	p2 := 9
	score1 := 0
	score2 := 0
	cur_val := 1
	rolls := 0
	for {
		rolls += 3
		p1 = (p1 + 3*cur_val + 3) % 10
		cur_val = (cur_val + 3) % 100
		if p1 > 0 {
			score1 += p1
		} else {
			score1 += 10
		}
		if score1 > 999 {
			break
		}
		p2 = (p2 + 3*cur_val + 3) % 10
		rolls += 3
		cur_val = (cur_val + 3) % 100
		if p2 > 0 {
			score2 += p2
		} else {
			score2 += 10
		}
		if score2 > 999 {
			break
		}
	}
	return utils.Min(score1, score2) * rolls

}

var DiceCombos = map[int]int64{6: 7, 5: 6, 7: 6, 4: 3, 8: 3, 3: 1, 9: 1}

func play(starting_positions, starting_scores [2]int64, player int) [2]int64 {
	var wins [2]int64

	for sum_throw, universes := range DiceCombos {
		positions := starting_positions
		scores := starting_scores

		positions[player] = (positions[player]+int64(sum_throw)-1)%10 + 1
		scores[player] += positions[player]

		if scores[player] >= 21 {
			wins[player] += universes
		} else {
			deepwins := play(positions, scores, (player+1)%2)
			for i := int64(0); i < int64(2); i++ {
				wins[i] += deepwins[i] * universes
			}
		}
	}
	return wins
}

func Day21P02() int64 {
	wins := play([2]int64{5, 9}, [2]int64{0, 0}, 0)
	return utils.Max64(wins[0], wins[1])
}
