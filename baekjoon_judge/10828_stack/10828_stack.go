package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Stack struct {
	data []int
}

func (s *Stack) Push(d int) {
	s.data = append(s.data, d)
}

func (s *Stack) Pop() int {
	d := s.Top()
	if d != -1 {
		s.data = s.data[:len(s.data)-1]
	}
	return d
}

func (s *Stack) Size() int {
	return len(s.data)
}

func (s *Stack) IsEmpty() int {
	if len(s.data) == 0 {
		return 1
	}
	return 0
}

func (s *Stack) Top() int {
	if len(s.data) == 0 {
		return -1
	}
	return s.data[len(s.data)-1]
}

func main() {
	var N int
	var d int
	scanner := bufio.NewScanner(os.Stdin)

	stack := new(Stack)
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &N)
	for i := 0; i < N; i++ {
		scanner.Scan()
		parse := strings.Split(scanner.Text(), " ")
		switch parse[0] {
		case "push":
			fmt.Sscan(parse[1], &d)
			stack.Push(d)
		case "pop":
			fmt.Printf("%d\n", stack.Pop())
		case "size":
			fmt.Printf("%d\n", stack.Size())
		case "empty":
			fmt.Printf("%d\n", stack.IsEmpty())
		case "top":
			fmt.Printf("%d\n", stack.Top())
		}
	}
}
