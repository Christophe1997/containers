package containers

type Stack[E comparable] interface {
	Sized
	Pop() (Stack[E], E)
	Push(e E) Stack[E]
	Peek() E
	ToList() List[E]
}

type stack[E comparable] struct {
	data LinkedList[E]
}

func NewStack[E comparable]() Stack[E] {
	return &stack[E]{data: NewLinkedList[E]()}
}

func (s *stack[E]) Pop() (Stack[E], E) {
	if s.IsEmpty() {
		panic("Empty Stack")
	} else {
		res := s.data.Back()
		s.data.RemoveBack()
		return s, res
	}
}

func (s *stack[E]) Push(e E) Stack[E] {
	s.data.PushBack(e)
	return s
}

func (s *stack[E]) Peek() E {
	return s.data.Back()
}

func (s *stack[E]) Size() int {
	return s.data.Size()
}

func (s *stack[E]) IsEmpty() bool {
	return s.data.IsEmpty()
}

func (s *stack[E]) ToList() List[E] {
	res := NewLinkedList[E]()
	s.data.ForEach(func(e E) {
		res.Add(e)
	})
	return res
}
