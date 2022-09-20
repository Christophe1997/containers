package containers

// RBTree
// Every node has a color either red or black.
// The root of the tree is always black.
// There are no two adjacent red nodes (A red node cannot have a red parent or red child).
// Every path from a node (including root) to any of its descendants NULL nodes has the same number of black nodes.
// All leaf nodes are black nodes.
type RBTree[E comparable] interface {
	Sized
	Search(cmp func(lhs E, rhs E) int, e E) (E, bool)
	Insert(cmp func(lhs E, rhs E) int, e ...E) RBTree[E]
	Height() int
	debug() bool
}

type treeNode[E comparable] struct {
	left   *treeNode[E]
	right  *treeNode[E]
	parent *treeNode[E]
	isRed  bool
	val    E
}

type rbTree[E comparable] struct {
	root *treeNode[E]
	size int
}

func NewRBTree[E comparable]() RBTree[E] {
	return &rbTree[E]{}
}

func (t *rbTree[E]) IsEmpty() bool {
	return t.root == nil
}

func (t *rbTree[E]) Size() int {
	return t.size
}

func (t *rbTree[E]) Search(cmp func(lhs E, rhs E) int, e E) (E, bool) {
	var res E
	cur := t.root
	for cur != nil {
		v := cmp(e, cur.val)
		if v == 0 {
			return cur.val, true
		} else if v < 0 {
			cur = cur.left
		} else {
			cur = cur.right
		}
	}
	return res, false
}

func (t *rbTree[E]) Insert(cmp func(lhs, rhs E) int, e ...E) RBTree[E] {
	res := t
	for _, v := range e {
		res = res.insert(cmp, v)
	}
	return res
}

func (t *rbTree[E]) Height() int {
	return t.root.height()
}

func (t *rbTree[E]) debug() bool {
	rootIsBlack := !t.root.isRed
	return rootIsBlack && t.root.colorOK() && t.root.blackCountOK()
}

// Every path from a node (including root) to any of its descendants NULL nodes has the same number of black nodes.
func (t *treeNode[E]) blackCountOK() bool {
	if t == nil {
		return true
	}
	nodes := t.leaves()
	if len(nodes) == 0 {
		return true
	}
	blackCount := make([]int, len(nodes))
	notAtRoot := len(nodes)
	atRoot := make([]bool, len(nodes))
	for notAtRoot > 0 {
		for idx, node := range nodes {
			if node != t {
				if !node.isRed {
					blackCount[idx]++
				}
				nodes[idx] = node.parent
			} else if !atRoot[idx] {
				atRoot[idx] = true
				notAtRoot--
			}
		}
	}
	count := blackCount[0]
	var res = true
	for _, v := range blackCount[1:] {
		if v != count {
			res = false
		}
	}
	return res && t.left.blackCountOK() && t.right.blackCountOK()
}

func (t *treeNode[E]) leaves() []*treeNode[E] {
	if t == nil {
		return []*treeNode[E]{}
	}
	queue := NewQueue[*treeNode[E]]()
	queue.Enqueue(t)
	var (
		curNode *treeNode[E]
		res     = NewLinkedList[*treeNode[E]]()
	)
	for !queue.IsEmpty() {
		size := queue.Size()
		for idx := 0; idx < size; idx++ {
			queue, curNode = queue.Dequeue()
			if curNode.left == nil && curNode.right == nil {
				res.Add(curNode)
			}
			if curNode.left != nil {
				queue.Enqueue(curNode.left)
			}
			if curNode.right != nil {
				queue.Enqueue(curNode.right)
			}
		}
	}
	return res.ToArray()
}

// There are no two adjacent red nodes (A red node cannot have a red parent or red child).
// All leaf nodes are black nodes.
func (t *treeNode[E]) colorOK() bool {
	if t == nil {
		return true
	} else if t.left == nil && t.right == nil {
		return !t.isRed
	} else if t.isRed {
		if t.left == nil {
			return !t.right.isRed && t.right.colorOK()
		} else if t.right == nil {
			return !t.left.isRed && t.left.colorOK()
		} else {
			return !t.left.isRed && !t.right.isRed && t.left.colorOK() && t.right.colorOK()
		}
	}
	return true
}

func (t *treeNode[E]) height() int {
	if t == nil {
		return 0
	}
	leftHeight := t.left.height()
	rightHeight := t.right.height()
	if leftHeight > rightHeight {
		return leftHeight + 1
	} else {
		return rightHeight + 1
	}
}

func (t *treeNode[E]) findRoot() *treeNode[E] {
	cur := t
	for cur.parent != nil {
		cur = cur.parent
	}
	return cur
}

func (t *rbTree[E]) insert(cmp func(lhs, rhs E) int, e E) *rbTree[E] {
	res := &treeNode[E]{isRed: true, val: e}
	if t.root == nil {
		res.isRed = false
		t.root = res
		return t
	}

	var parent *treeNode[E]
	cur := t.root
	for cur != nil {
		parent = cur
		v := cmp(e, cur.val)
		if v == 0 {
			return t
		} else if v < 0 {
			cur = cur.left
		} else if v > 0 {
			cur = cur.right
		}
	}
	if parent.left == cur {
		parent.left = res
	} else {
		parent.right = res
	}
	res.parent = parent

	res.balance()
	t.root = t.root.findRoot()
	t.size++
	return t
}

// rotation left-left case
func (t *treeNode[E]) rotationRight() {
	if t == nil || t.parent == nil {
		return
	}
	var (
		parent = t.parent
		grandP = parent.parent
		right  = t.right
	)
	// if nil then grandP is root
	if grandP != nil {
		if grandP.left == parent {
			grandP.left = t
		} else {
			grandP.right = t
		}
	}
	// rotation
	t.right = parent
	parent.left = right

	// reset parent
	t.parent = grandP
	parent.parent = t
	right.parent = parent
}

// rotation right-right case
func (t *treeNode[E]) rotationLeft() {
	if t == nil || t.parent == nil {
		return
	}
	var (
		parent = t.parent
		grandP = parent.parent
		left   = t.left
	)
	// if nil then grandP is root
	if grandP != nil {
		if grandP.left == parent {
			grandP.left = t
		} else {
			grandP.right = t
		}
	}

	// rotation
	t.left = parent
	parent.right = left

	// reset parent
	t.parent = grandP
	parent.parent = t
	left.parent = parent
}

func (t *treeNode[E]) balance() {
	cur := t
	if cur == nil || !cur.isRed {
		return
	}
	var (
		parent        = cur.parent
		grapdP, uncle *treeNode[E]
	)
	if parent == nil {
		cur.isRed = false
		return
	} else if !parent.isRed {
		return
	}
	// must exist, because root is black
	grapdP = parent.parent
	if grapdP.left == parent {
		uncle = grapdP.right
	} else {
		uncle = grapdP.left
	}

	if uncle.isRed {
		parent.isRed = false
		uncle.isRed = false
		grapdP.isRed = true
		grapdP.balance()
	} else {
		isLeft1 := grapdP.left == parent
		isLeft2 := parent.left == cur

		switch {
		// left-left
		case isLeft1 && isLeft2:
			parent.rotationRight()
			parent.isRed = false
			grapdP.isRed = true
		// left-right
		case isLeft1 && !isLeft2:
			cur.rotationLeft()
			cur.rotationRight()
			cur.isRed = false
			grapdP.isRed = true
		// right-right
		case !isLeft1 && !isLeft2:
			parent.rotationLeft()
			parent.isRed = false
			grapdP.isRed = true
		// right-left
		case !isLeft1 && isLeft2:
			cur.rotationRight()
			cur.rotationLeft()
			cur.isRed = false
			grapdP.isRed = true
		}
	}
}
