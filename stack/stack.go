package stack

type Stack struct {
	Items []interface{}
	Top   int
}

//constant time amoritized.
func (s *Stack) Push(v int) {
	s.Items = append(s.Items, v)
	s.Top = len(s.Items)
}

//constant time
func (s *Stack) Peek() interface{} {
	return s.Items[s.Top-1]
}

//constant time
func (s *Stack) Pop() interface{} {
	s.Top = s.Top - 1
	v := s.Items[s.Top]
	s.Items = append(s.Items[:s.Top])
	return v
}
