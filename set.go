package lisq

type Set[E comparable] interface {
	Union(other Set[E]) Set[E]
	Intersection(other Set[E]) Set[E]
	Difference(other Set[E]) Set[E]
}
