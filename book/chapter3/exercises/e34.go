package chapter3

type Stack struct {
	elements []int
	min      int
}

func (s *Stack) push(el int) {
	s.elements = append(s.elements, el)
	if el < s.min {
		s.min = el
	}
}

func (s *Stack) updateMin() {
	newMin := s.elements[0]
	for _, v := range s.elements {
		if v < newMin {
			newMin = v
		}
	}

	s.min = newMin

}

func (s *Stack) pop() int {
	el := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	if el == s.min {
		s.updateMin()
	}

	return el
}

func (s *Stack) findmin() int {
	return s.min
}
