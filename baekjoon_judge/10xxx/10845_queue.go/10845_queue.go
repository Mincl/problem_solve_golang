package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Queue struct {
	data []int
}

func (q *Queue) Push(d int) {
	q.data = append(q.data, d)
}

func (q *Queue) Pop() int {
	d := q.Front()
	if d != -1 {
		q.data = q.data[1:]
	}
	return d
}

func (q *Queue) Size() int {
	return len(q.data)
}

func (q *Queue) IsEmpty() int {
	if len(q.data) == 0 {
		return 1
	}
	return 0
}

func (q *Queue) Front() int {
	if len(q.data) == 0 {
		return -1
	}
	return q.data[0]
}

func (q *Queue) Back() int {
	if len(q.data) == 0 {
		return -1
	}
	return q.data[len(q.data)-1]
}

func main() {
	var N, d int
	scanner := bufio.NewScanner(os.Stdin)

	queue := new(Queue)
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &N)
	for i := 0; i < N; i++ {
		scanner.Scan()
		parsed := strings.Split(scanner.Text(), " ")
		switch parsed[0] {
		case "push":
			fmt.Sscan(parsed[1], &d)
			queue.Push(d)
		case "pop":
			fmt.Printf("%d\n", queue.Pop())
		case "size":
			fmt.Printf("%d\n", queue.Size())
		case "empty":
			fmt.Printf("%d\n", queue.IsEmpty())
		case "front":
			fmt.Printf("%d\n", queue.Front())
		case "back":
			fmt.Printf("%d\n", queue.Back())
		}
	}
}
