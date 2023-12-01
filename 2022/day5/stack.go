package day5

type Stack struct {
	crates []byte
	length int
}

func (s *Stack) insert(value byte) {
	s.crates = append(s.crates, value)
	s.length++
}

func (s *Stack) pop() byte {
	if s.length == 0 {
		return 0
	}

	out := s.crates[s.length-1]
	s.crates = s.crates[:s.length-1]
	s.length--
	return out
}

func (s *Stack) peek() byte {
	if s.length == 0 {
		return 0
	}
	return s.crates[s.length-1]
}

func (s *Stack) insertStack(newCrates []byte) {
	s.crates = append(s.crates, newCrates...)
}
