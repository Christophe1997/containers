package lisq

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

func Fold[A, B Elem](f func(B, A) B, zero B, src iter.Seq[A]) B {
	res := zero
	for v := range src {
		res = f(res, v)
	}
	return res
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

func Map[A, B Elem](f func(A) B, src iter.Seq[A]) iter.Seq[B] {
	return Fold(func(b iter.Seq[B], a A) iter.Seq[B] {
		return Append(b, f(a))
	}, Zero[B](), src)
}

func Map2[A, B, C Elem](f func(A, B) C, src1 iter.Seq[A], src2 iter.Seq[B]) iter.Seq[C] {
	next1, stop1 := iter.Pull(src1)
	next2, stop2 := iter.Pull(src2)
	return func(yield func(C) bool) {
		defer stop1()
		defer stop2()
		a, hasNext1 := next1()
		b, hasNext2 := next2()
		for hasNext1 && hasNext2 {
			if !yield(f(a, b)) {
				return
			}
			a, hasNext1 = next1()
			b, hasNext2 = next2()
		}
	}
}

func Filter[A Elem](f func(A) bool, src iter.Seq[A]) iter.Seq[A] {
	return Fold(func(b iter.Seq[A], a A) iter.Seq[A] {
		if f(a) {
			return Append(b, a)
		}
		return b
	}, Zero[A](), src)
}

func Reduce[A, B Elem](f func(B, A) B, src iter.Seq[A]) B {
	var res B
	return Fold(f, res, src)
}

func FlatMap[A, B Elem](f func(A) iter.Seq[B], src iter.Seq[A]) iter.Seq[B] {
	return Fold(func(b iter.Seq[B], a A) iter.Seq[B] {
		return Fold(func(b2 iter.Seq[B], a2 B) iter.Seq[B] {
			return Append(b2, a2)
		}, b, f(a))
	}, Zero[B](), src)
}

func Take[A Elem](n int, src iter.Seq[A]) iter.Seq[A] {
	return func(yield func(A) bool) {
		for idx, v := range ConvertSeq2(src) {
			if idx >= n || !yield(v) {
				return
			}
		}
	}
}

var NaturalNumbers iter.Seq[int] = func(yield func(int) bool) {
	i := 0
	for {
		if !yield(i) {
			return
		}
		i++
	}
}
