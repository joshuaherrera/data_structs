package stack

type Stack struct {
	Items []int
	Top   int
}

//constant time amoritized.
func (s *Stack) Push(v int) {
	s.Items = append(s.Items, v)
	s.Top = len(s.Items)
}

//constant time
func (s *Stack) Peek() int {
	return s.Items[s.Top-1]
}

//constant time
func (s *Stack) Pop() int {
	s.Top = s.Top - 1
	v := s.Items[s.Top]
	s.Items = append(s.Items[:s.Top])
	return v
}
