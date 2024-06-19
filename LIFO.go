package main

type LIFO struct {
	stack []interface{}
}

func (s *LIFO) push(element interface{}) {
	s.stack = append(s.stack, element)
}

func (s *LIFO) pop() interface{} {
	deleted := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]

	return deleted
}
