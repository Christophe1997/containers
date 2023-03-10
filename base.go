package containers

type Iterable[E any] interface {
	ForEach(do func(E))
	Iterator() Iterator[E]
}

type Iterator[E any] interface {
	HasNext() bool
	Next() E
}

type Sized interface {
	Size() int
	IsEmpty() bool
}

type Collection[E comparable] interface {
	Iterable[E]
	Sized

	Add(e ...E)
	Contains(e E) bool
	ContainsWith(cmp func(lhs, rhs E) bool, e E) bool
	Remove(e ...E)
	RemoveWith(cmp func(lhs, rhs E) bool, e ...E)
}

func Copy[E comparable](dest Collection[E], src Collection[E]) {
	src.ForEach(func(e E) {
		dest.Add(e)
	})
}

func Fmap[A comparable, B comparable](f func(a A) B) func(ls List[A]) List[B] {
	return func(ls List[A]) List[B] {
		res := NewLinkedList[B]()
		ls.ForEach(func(a A) {
			res.Add(f(a))
		})
		return res
	}
}
