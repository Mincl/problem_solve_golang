package main

import (
	"fmt"
)

type Stack struct {
	data []int
}

func (s *Stack) Push(value int) error {
	s.data = append(s.data, value)
	return nil
}

func (s *Stack) Pop() (int, error) {
	res := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return res, nil
}

func (s *Stack) Top() int {
	if s.Size() == 0 {
		return -1
	}
	return s.data[len(s.data)-1]
}

func (s *Stack) Size() int {
	return len(s.data)
}

func main() {
	var N int
	fmt.Scanf("%d\n", &N)
	students := make([]int, N)
	for i := range students {
		fmt.Scanf("%d", &students[i])
	}

	if proc(students) {
		fmt.Printf("Nice")
	} else {
		fmt.Printf("Sad")
	}
}

func proc(students []int) bool {
	s := &Stack{make([]int, 0)}
	cur_num := 1
	for i := 0; i < len(students); {
		if cur_num == students[i] {
			cur_num++
			i++
		} else if cur_num == s.Top() {
			s.Pop()
			cur_num++
		} else {
			s.Push(students[i])
			i++
		}
	}

	for s.Size() > 0 {
		v, _ := s.Pop()
		if cur_num == v {
			cur_num++
		} else {
			return false
		}
	}

	return true
}
