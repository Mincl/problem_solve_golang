package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Queue struct {
	point []Direction
}

type Direction struct {
	x int
	y int
}

func (q *Queue) Size() int {
	return len(q.point)
}
func (q *Queue) Push(x, y int) {
	q.point = append(q.point, Direction{x, y})
}
func (q *Queue) Pop() Direction {
	dir := q.point[0]
	q.point = q.point[1:]
	return dir
}

func tomatoDay(box [][]int, N, M int) int {
	max := 0
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			if box[i][j] == 0 {
				return -1
			}
			max = int(math.Max(float64(max), float64(box[i][j])))
		}
	}
	return max
}

func main() {
	var M, N int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	fmt.Sscan(scanner.Text(), &M)
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &N)
	box := make([][]int, N)
	for i := 0; i < N; i++ {
		box[i] = make([]int, M)
		for j := 0; j < M; j++ {
			scanner.Scan()
			fmt.Sscan(scanner.Text(), &box[i][j])
		}
	}

	direction := []Direction{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	queue := new(Queue)
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			if box[i][j] == 1 {
				queue.Push(i, j)
			}
		}
	}

	for queue.Size() > 0 {
		point := queue.Pop()
		for _, dir := range direction {
			if 0 <= point.x+dir.x && point.x+dir.x < N &&
				0 <= point.y+dir.y && point.y+dir.y < M &&
				box[point.x+dir.x][point.y+dir.y] == 0 {
				box[point.x+dir.x][point.y+dir.y] = box[point.x][point.y] + 1
				queue.Push(point.x+dir.x, point.y+dir.y)
			}
		}
	}
	day := tomatoDay(box, N, M)
	if day > 0 {
		day--
	}
	fmt.Printf("%d\n", day)
}
