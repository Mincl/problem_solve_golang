package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	x int
	y int
}

func (p Pos) Equal(p2 Pos) bool {
	if p.x == p2.x && p.y == p2.y {
		return true
	}
	return false
}

type DirPos struct {
	pos     Pos
	prevDir Pos
}

type Queue struct {
	data []DirPos
}

func (q *Queue) Push(d DirPos) {
	q.data = append(q.data, d)
}

func (q *Queue) Pop() DirPos {
	d := q.data[0]
	q.data = q.data[1:]
	return d
}

func (q *Queue) Size() int {
	return len(q.data)
}

func BFS(pipeMap [][]int, R, C, pipeCnt int, st, ed Pos) bool {
	dirLeft := Pos{0, -1}
	dirRight := Pos{0, 1}
	dirUp := Pos{-1, 0}
	dirDown := Pos{1, 0}
	dirSet := make([]map[Pos]Pos, 8)
	// start
	startSet := [4]Pos{
		dirUp,
		dirDown,
		dirLeft,
		dirRight,
	}
	// dirSet[PipeType][inDirection] = outDirection
	// '|'
	dirSet[1] = make(map[Pos]Pos)
	dirSet[1][dirDown] = dirDown
	dirSet[1][dirUp] = dirUp
	// '-'
	dirSet[2] = make(map[Pos]Pos)
	dirSet[2][dirLeft] = dirLeft
	dirSet[2][dirRight] = dirRight
	// '1'
	dirSet[3] = make(map[Pos]Pos)
	dirSet[3][dirLeft] = dirDown
	dirSet[3][dirUp] = dirRight
	// '2'
	dirSet[4] = make(map[Pos]Pos)
	dirSet[4][dirLeft] = dirUp
	dirSet[4][dirDown] = dirRight
	// '3'
	dirSet[5] = make(map[Pos]Pos)
	dirSet[5][dirRight] = dirUp
	dirSet[5][dirDown] = dirLeft
	// '4'
	dirSet[6] = make(map[Pos]Pos)
	dirSet[6][dirRight] = dirDown
	dirSet[6][dirUp] = dirLeft
	// '+'
	dirSet[7] = make(map[Pos]Pos)
	dirSet[7][dirUp] = dirUp
	dirSet[7][dirDown] = dirDown
	dirSet[7][dirLeft] = dirLeft
	dirSet[7][dirRight] = dirRight

	check := make([][]int, R)
	for k := 0; k < R; k++ {
		check[k] = make([]int, C)
	}

	q := new(Queue)
	q.Push(DirPos{Pos{st.x, st.y}, Pos{0, 0}})
	end := false
	check[st.x][st.y] = 1
	for q.Size() > 0 {
		dirPos := q.Pop()

		if dirPos.pos.Equal(ed) {
			pipeCnt++
			end = true
			break
		}

		k := pipeMap[dirPos.pos.x][dirPos.pos.y]
		if dirPos.pos.Equal(st) {
			for l := 0; l < 4; l++ {
				x := dirPos.pos.x + startSet[l].x
				y := dirPos.pos.y + startSet[l].y
				if CanMove(pipeMap, R, C, x, y) {
					k = pipeMap[x][y]
					if k == -1 {
						continue
					}
					if _, ok := dirSet[k][startSet[l]]; ok {
						q.Push(DirPos{Pos{x, y}, startSet[l]})
						check[x][y]++
						if check[x][y] == 1 {
							pipeCnt--
						}
						// only one path
						break
					}
				}
			}
		} else {
			if nextPos, ok := dirSet[k][dirPos.prevDir]; ok {
				x := dirPos.pos.x + nextPos.x
				y := dirPos.pos.y + nextPos.y
				if CanMove(pipeMap, R, C, x, y) &&
					((k == 7 && check[x][y] <= 1) || k != 7) {
					q.Push(DirPos{Pos{x, y}, nextPos})
					check[x][y]++
					if check[x][y] == 1 {
						pipeCnt--
					}
				}
			}
		}
	}

	if end && pipeCnt == 0 {
		return true
	}
	return false
}

func CanMove(pipeMap [][]int, R, C, x, y int) bool {
	if 0 <= x && x < R && 0 <= y && y < C &&
		pipeMap[x][y] != 0 {
		return true
	}
	return false
}

func main() {
	var R, C int
	var row string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &R, &C)
	pipeCnt := 0
	pipeMap := make([][]int, R)
	st := Pos{0, 0}
	ed := Pos{0, 0}
	for i := 0; i < R; i++ {
		pipeMap[i] = make([]int, C)
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &row)
		for j, val := range row {
			switch val {
			case '.':
				pipeMap[i][j] = 0
			case 'M':
				pipeMap[i][j] = -1
				st.x = i
				st.y = j
			case 'Z':
				pipeMap[i][j] = -1
				ed.x = i
				ed.y = j
			case '|':
				pipeMap[i][j] = 1
				pipeCnt++
			case '-':
				pipeMap[i][j] = 2
				pipeCnt++
			case '1':
				pipeMap[i][j] = 3
				pipeCnt++
			case '2':
				pipeMap[i][j] = 4
				pipeCnt++
			case '3':
				pipeMap[i][j] = 5
				pipeCnt++
			case '4':
				pipeMap[i][j] = 6
				pipeCnt++
			case '+':
				pipeMap[i][j] = 7
				pipeCnt++
			}
		}
	}

	possible := false
	pipeCnt++
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			if pipeMap[i][j] == 0 {
				for k := 1; k <= 7; k++ {
					// make block
					pipeMap[i][j] = k

					possible = BFS(pipeMap, R, C, pipeCnt, st, ed)
					if possible {
						fmt.Printf("%d %d ", i+1, j+1)
						switch k {
						case 1:
							fmt.Printf("|")
						case 2:
							fmt.Printf("-")
						case 3:
							fmt.Printf("1")
						case 4:
							fmt.Printf("2")
						case 5:
							fmt.Printf("3")
						case 6:
							fmt.Printf("4")
						case 7:
							fmt.Printf("+")
						}
						return
					}

					// remove block
					pipeMap[i][j] = 0
				}
			}
		}
	}
}
