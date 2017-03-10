package main

import (
	"fmt"
)

type Vector struct {
	x int
	y int
}

func win_check(board [][]int) bool {
	n := 6
	m := 7
	direction := [8]Vector{
		{0, 1}, {0, -1}, {1, 0}, {-1, 0},
		{1, 1}, {1, -1}, {-1, 1}, {-1, -1},
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] == 0 {
				continue
			}
			for k := 0; k < 8; k++ {
				kimbab_cnt := 0
				for l := 1; l <= 3; l++ {
					x := i + direction[k].x*l
					y := j + direction[k].y*l
					if 0 > x || x >= n ||
						0 > y || y >= m {
						break
					}
					if board[i][j] == board[x][y] {
						kimbab_cnt++
					}
				}
				if kimbab_cnt == 3 {
					return true
				}
			}
		}
	}
	return false
}

func main() {
	var s, j, win_turn int
	var winner string
	board := make([][]int, 6)
	board_cnt := make([]int, 7)
	for i := 0; i < 6; i++ {
		board[i] = make([]int, 7)
	}
	for i := 0; i < 21; i++ {
		fmt.Scan(&s, &j)
		board[board_cnt[s-1]][s-1] = 1
		if win_turn == 0 && win_check(board) {
			winner = "sk"
			win_turn = i + 1
		}
		board_cnt[s-1]++
		board[board_cnt[j-1]][j-1] = 2
		board_cnt[j-1]++
		if win_turn == 0 && win_check(board) {
			winner = "ji"
			win_turn = i + 1
		}
	}
	if win_turn == 0 {
		fmt.Print("ss")
	} else {
		fmt.Printf("%s %d", winner, win_turn)
	}
}
