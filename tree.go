package containers

type binaryTree[E comparable] struct {
	left  *binaryTree[E]
	right *binaryTree[E]
	val   E
}
