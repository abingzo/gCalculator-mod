package test

import (
	"gCalculator-mod/alg"
	"testing"
)

func TestStack(t *testing.T) {
	stack := alg.NewStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	t.Log(stack.Pop())
	t.Log(stack.Pop())
	t.Log(stack.Peek())
	t.Log(stack.IsEmpty())
}
