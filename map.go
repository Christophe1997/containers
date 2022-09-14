package containers

type Map[K, V comparable] interface {
	Sized

	ForEach(do func(key K, val V))

	Put(key K, val V) Map[K, V]
	Get(key K) (V, bool)
	Keys() Set[K]
	Values() Collection[V]
}
