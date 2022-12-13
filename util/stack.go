package util

// ====================
// STACK IMPLEMENTATION
// ====================

type StackEle struct {
	val  interface{}
	next *StackEle
}

type Stack struct {
	Top *StackEle
}

func (s *Stack) Push(d interface{}) {
	(*s).Top = &StackEle{d, (*s).Top}
}

func (s *Stack) Pop() (interface{}, bool) {
	if (*s).Top == nil {
		return nil, false
	}

	d := (*s).Top.val
	(*s).Top = (*s).Top.next

	return d, true
}

func (s *Stack) TopVal() interface{} {
	if (*s).Top == nil {
		return nil
	}

	d := (*s).Top.val
	return d
}

func (s *Stack) Len() int {
	current := s.Top
	len := 0
	if current == nil {
		return len
	}

	for {
		if current.next == nil {
			return len
		}

		len++
		current = current.next
	}
}
