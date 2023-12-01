package day7

import (
	"errors"
)

type Stack struct {
	path   []string
	length int
}

func (s *Stack) insert(dir string) {
	s.length++
	s.path = append(s.path, dir)
}

func (s *Stack) pop() error {
	if s.length == 0 {
		return errors.New("cannot pop stack - is empty")
	}

	s.length--
	s.path = s.path[:s.length]
	return nil
}

func (s *Stack) reset() {
	s.path = []string{"/"}
	s.length = 1
}
