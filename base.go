package containers

import "iter"

type Elem = any

type UnSizedSeq[E Elem] interface {
	All() (seq iter.Seq2[int, E])
	Values() (seq iter.Seq[E])
}

type SizedSeq[E Elem] interface {
	Size() (size int)
	UnSizedSeq[E]
}

func Zero[E Elem]() iter.Seq[E] {
	return func(yield func(E) bool) {}
}

func Append[E Elem](seq iter.Seq[E], e ...E) iter.Seq[E] {
	return func(yield func(E) bool) {
		for v := range seq {
			if !yield(v) {
				return
			}
		}
		for _, v := range e {
			if !yield(v) {
				return
			}
		}
	}
}

func ConvertSeq2[E Elem](seq iter.Seq[E]) iter.Seq2[int, E] {
	return func(yield func(int, E) bool) {
		idx := 0
		for v := range seq {
			if !yield(idx, v) {
				return
			}
			idx++
		}
	}
}

func Fold[A, B Elem](src iter.Seq[A], f func(B, A) B, zero B) B {
	res := zero
	for v := range src {
		res = f(res, v)
	}
	return res
}

func Map[A, B Elem](src iter.Seq[A], f func(A) B) iter.Seq[B] {
	return Fold(src, func(b iter.Seq[B], a A) iter.Seq[B] {
		return Append(b, f(a))
	}, Zero[B]())
}

func Reduce[A, B Elem](src iter.Seq[A], f func(B, A) B) B {
	var res B
	return Fold(src, f, res)
}

func Filter[A Elem](src iter.Seq[A], f func(A) bool) iter.Seq[A] {
	return Fold(src, func(b iter.Seq[A], a A) iter.Seq[A] {
		if f(a) {
			return Append(b, a)
		}
		return b
	}, Zero[A]())
}

func FlatMap[A, B Elem](src iter.Seq[A], f func(A) iter.Seq[B]) iter.Seq[B] {
	return Fold(src, func(b iter.Seq[B], a A) iter.Seq[B] {
		return Fold(f(a), func(b2 iter.Seq[B], a2 B) iter.Seq[B] {
			return Append(b2, a2)
		}, b)
	}, Zero[B]())
}
