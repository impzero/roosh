package roosh

import (
	"errors"
)

type Calculator interface {
	Add() (int, error)
	Sub() (int, error)
	Div() (int, error)
	Mul() (int, error)
}

type StackCalculator struct {
	Memory *Stack
}

func NewStackCalculator(mem *Stack) *StackCalculator {
	return &StackCalculator{Memory: mem}
}

// Add adds the top two elements from the stack
func (sc *StackCalculator) Add() (int, error) {
	if len(*sc.Memory) < 2 {
		return 0, errors.New("not enough elements in the stack")
	}

	a := sc.Memory.Pop()
	b := sc.Memory.Pop()
	res := a + b
	sc.Memory.Push(res)
	return res, nil
}

// Sub subtracts the top two elements from the stack
func (sc *StackCalculator) Sub() (int, error) {
	if len(*sc.Memory) < 2 {
		return 0, errors.New("not enough elements in the stack")
	}

	a := sc.Memory.Pop()
	b := sc.Memory.Pop()
	res := a - b
	sc.Memory.Push(res)
	return res, nil
}

// Mul multiplies the top two elements from memory
func (sc *StackCalculator) Mul() (int, error) {
	if len(*sc.Memory) < 2 {
		return 0, errors.New("not enough elements in the stack")
	}

	a := sc.Memory.Pop()
	b := sc.Memory.Pop()
	res := a * b
	sc.Memory.Push(res)
	return res, nil
}

// Div divides the top most element from the second top most element
func (sc *StackCalculator) Div() (int, error) {
	if len(*sc.Memory) < 2 {
		return 0, errors.New("not enough elements in the stack")
	}

	a := sc.Memory.Pop()
	b := sc.Memory.Pop()
	res := a / b
	sc.Memory.Push(res)
	return res, nil
}

// Print prints the top most element from memory
func (sc *StackCalculator) Print() int {
	return sc.Memory.Fetch()[0]
}

type Stack []int

func (s *Stack) Fetch() []int {
	return *s
}

func (s *Stack) Push(n int) {
	*s = append(*s, n)
}

func (s *Stack) Pop() int {
	if len(*s) == 0 {
		return 0
	}

	n := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return n
}
