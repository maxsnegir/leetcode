package main

import (
	"fmt"
)

func main() {
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	fmt.Println(obj.Top())
	fmt.Println(obj.Pop())
	fmt.Println(obj.Empty())
}

type MyStack struct {
	q []int
}

func Constructor() MyStack {
	return MyStack{
		q: make([]int, 0),
	}

}

func (s *MyStack) Push(x int) {
	s.q = append(s.q, x)
}

func (s *MyStack) Pop() int {
	val := s.q[len(s.q)-1]
	s.q = s.q[:len(s.q)-1]
	return val
}

func (s *MyStack) Top() int {
	return s.q[len(s.q)-1]
}

func (s *MyStack) Empty() bool {
	return len(s.q) == 0
}
