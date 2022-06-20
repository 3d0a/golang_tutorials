package main

import (
	"fmt"
)

type Stack struct {
	arr []string
}

func (s *Stack) Push(elem string) {
	s.arr = append(s.arr, elem)
}

func (s *Stack) Pop() string {
	last_elem := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return last_elem
}

func (s *Stack) IsEmpty() bool {
	return len(s.arr) == 0
}

func isValid(s string) bool {
	stack := Stack{}
	for _, elem := range s {
		strElem := string(elem)
		switch {
		case strElem == "{":
			stack.Push("}")
		case strElem == "(":
			stack.Push(")")
		case strElem == "[":
			stack.Push("]")
		case stack.IsEmpty() || stack.Pop() != strElem:
			return false
		}
	}
	return stack.IsEmpty()
}

func main() {
	str := "{}()[]"
	fmt.Println(isValid(str))
}
