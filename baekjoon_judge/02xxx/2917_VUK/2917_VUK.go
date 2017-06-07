package main

import (
	"bufio"
	"fmt"
	"os"
)

type Direction struct {
	x int
	y int
}

type Axis Direction

type Queue struct {
	data []Axis
}

func (s *Queue) Push(a Axis) {
	s.data = append(s.data, a)
}

func (s *Queue) Pop() Axis {
	a := s.data[0]
	s.data = s.data[1:]
	return a
}

func (s *Queue) Size() int {
	return len(s.data)
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	var N, M int
	var temp string
	var st, ed Axis
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &N, &M)
	field := make([][]int, N)
	check := make([][]int, N)
	for i := 0; i < N; i++ {
		field[i] = make([]int, M)
		check[i] = make([]int, M)
		for j := 0; j < M; j++ {
			field[i][j] = 1 << 31
		}
	}
	s := new(Queue)
	for i := 0; i < N; i++ {
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &temp)
		for j := 0; j < M; j++ {
			switch temp[j] {
			case '+':
				s.Push(Axis{i, j})
				field[i][j] = 0
			case 'V':
				st.x = i
				st.y = j
			case 'J':
				ed.x = i
				ed.y = j
			}
		}
	}

	dir := [4]Direction{
		Direction{1, 0},
		Direction{-1, 0},
		Direction{0, 1},
		Direction{0, -1},
	}

	for s.Size() > 0 {
		a := s.Pop()
		for i := 0; i < 4; i++ {
			x := a.x + dir[i].x
			y := a.y + dir[i].y
			if 0 <= x && x < N && 0 <= y && y < M &&
				field[x][y] == 1<<31 {
				s.Push(Axis{x, y})
				field[x][y] = field[a.x][a.y] + 1
			}
		}
	}

	// for i := 0; i < N; i++ {
	// 	fmt.Println(field[i])
	// }

	left := 1
	right := 1000
	max := 0
	k := 0
	goal := false
	for left <= right {
		k = (left + right) / 2
		goal = false
		if field[st.x][st.y] < k {
			goal = false
		} else {
			for i := 0; i < N; i++ {
				for j := 0; j < M; j++ {
					check[i][j] = 0
				}
			}

			// BFS
			s.Push(Axis{st.x, st.y})
			check[st.x][st.y] = 1
			for s.Size() > 0 {
				a := s.Pop()
				for i := 0; i < 4; i++ {
					x := a.x + dir[i].x
					y := a.y + dir[i].y
					if 0 <= x && x < N && 0 <= y && y < M &&
						check[x][y] == 0 && field[x][y] >= k {
						check[x][y] = 1
						s.Push(Axis{x, y})
					}
				}
			}

			if check[ed.x][ed.y] != 0 {
				goal = true
				max = Max(max, k)
			} else {
				goal = false
			}
		}
		if goal {
			left = k + 1
		} else {
			right = k - 1
		}
	}

	fmt.Printf("%d\n", max)
}
