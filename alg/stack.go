package alg

/*
	简单的栈结构实现
*/

import "container/list"

type Stack struct {
	list *list.List
}

func NewStack() *Stack {
	return &Stack{list: list.New()}
}

func (s *Stack) Pop() interface{} {
	return s.list.Remove(s.list.Back())
}

func (s *Stack) Push(val interface{}) {
	s.list.PushBack(val)
}

func (s *Stack) IsEmpty() bool {
	return s.list.Len() == 0
}

func (s *Stack) Len() int {
	return s.list.Len()
}

func (s *Stack) Peek() interface{} {
	return s.list.Back().Value
}
