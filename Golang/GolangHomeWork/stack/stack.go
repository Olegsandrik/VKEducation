package stack

import "errors"

type Stack []string

func (s *Stack) Push(value string) {
	*s = append(*s, value)
}

func (s *Stack) Pop() (string, error) {
	if len(*s) == 0 {
		err := errors.New("стек пуст")
		return "", err
	}
	value := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return value, nil
}

func (s *Stack) Len() int {
	return len(*s)
}

func (s *Stack) Peak() (string, error) {
	if len(*s) == 0 {
		err := errors.New("стек пуст")
		return "", err
	}
	return (*s)[len(*s)-1], nil
}
