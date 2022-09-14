package containers

type Set[E comparable] interface {
	Collection[E]
	Union(other Set[E]) Set[E]
	Intersection(other Set[E]) Set[E]
	Difference(other Set[E]) Set[E]
}

type hashSet[E comparable] struct {
	data map[E]bool
}

func NewHashSet[E comparable]() Set[E] {
	return &hashSet[E]{data: make(map[E]bool)}
}

func (s *hashSet[E]) Union(other Set[E]) Set[E] {
	res := NewHashSet[E]()
	s.ForEach(func(e E) {
		res.Add(e)
	})
	other.ForEach(func(e E) {
		res.Add(e)
	})
	return res
}

func (s *hashSet[E]) Intersection(other Set[E]) Set[E] {
	res := NewHashSet[E]()
	s.ForEach(func(e E) {
		if other.Contains(e) {
			res.Add(e)
		}
	})
	return res
}

func (s *hashSet[E]) Difference(other Set[E]) Set[E] {
	res := NewHashSet[E]()
	s.ForEach(func(e E) {
		if !other.Contains(e) {
			res.Add(e)
		}
	})
	return res
}

func (s *hashSet[E]) ForEach(do func(E)) {
	for k := range s.data {
		do(k)
	}
}

func (s *hashSet[E]) Iterator() Iterator[E] {
	res := NewLinkedList[E]()
	s.ForEach(func(e E) {
		res.Add(e)
	})
	return res.Iterator()
}

func (s *hashSet[E]) Size() int {
	return len(s.data)
}

func (s *hashSet[E]) IsEmpty() bool {
	return s.Size() == 0
}

func (s *hashSet[E]) Add(e ...E) {
	for _, v := range e {
		s.data[v] = true
	}
}

func (s *hashSet[E]) Contains(e E) bool {
	_, ok := s.data[e]
	return ok
}

func (s *hashSet[E]) ContainsWith(cmp func(lhs E, rhs E) bool, e E) bool {
	for k := range s.data {
		if cmp(k, e) {
			return true
		}
	}
	return false
}

func (s *hashSet[E]) Remove(e ...E) {
	for _, v := range e {
		delete(s.data, v)
	}
}

func (s *hashSet[E]) RemoveWith(cmp func(lhs E, rhs E) bool, e ...E) {
	for _, v := range e {
		for k := range s.data {
			if cmp(k, v) {
				delete(s.data, k)
			}
		}
	}
}
