package main

import "fmt"

type Stack struct {
	size  int
	stack []byte
}

func NewStack() *Stack {
	return &Stack{0, make([]byte, 10000)}
}

func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

func (s *Stack) Size() int {
	return s.size
}

func (s *Stack) Push(b byte) {
	s.stack[s.size] = b
	s.size++
}

func (s *Stack) Pop() {
	if s.size == 0 {
		return
	}

	s.size--
}

func (s *Stack) Top() byte {
	if s.size == 0 {
		return 0
	}

	return s.stack[s.size-1]
}

func IsBalancedParentheses(str string) bool {
	s := NewStack()
	for _, c := range str {
		switch c {
		case '(':
			s.Push(')')

		case '{':
			s.Push('}')

		case '[':
			s.Push(']')

		default:
			if c != rune(s.Top()) {
				return false
			}
			s.Pop()
		}
	}

	return s.IsEmpty()
}

func main() {
	for {
		fmt.Print("Please enter the string: ")
		var str string
		fmt.Scan(&str)
		if str == "q" || str == "Q" {
			break
		}

		b := IsBalancedParentheses(str)
		fmt.Print("The string is ")
		if !b {
			fmt.Print("not ")
		}
		fmt.Println("balenced!")
	}
}
