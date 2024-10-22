package stacks

// LAST IN FIRST OUT (LIFO)
type Stack struct {
	Items []int
}

func (s *Stack) Push(i int) {
	s.Items = append(s.Items, i)
}

func (s *Stack) Pop() int {
	lastIndex := len(s.Items) - 1
	toRemove := s.Items[lastIndex] // get the last element
	s.Items = s.Items[:lastIndex] // remove the last element
	return toRemove
}
