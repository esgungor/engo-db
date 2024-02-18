package avl

type Node struct {
	record *Record
	height int
	left   *Node
	right  *Node
}

type Record struct {
	Key   string
	Value string
}

func NewNode(r *Record) *Node {
	return &Node{
		record: r,
		height: 1,
	}
}

func Insert(n *Node, r Record) *Node {
	if n == nil {
		n = NewNode(&r)
		return n
	}
	if n.record.Key > r.Key {
		n.left = Insert(n.left, r)
	} else if n.record.Key < r.Key {
		n.right = Insert(n.right, r)

	} else {
		return n
	}
	n.height = 1 + max(height(n.left), height(n.right))

	balance := getBalance(n)

	if balance > 1 && r.Key < n.left.record.Key {
		return rightRotate(n)
	}
	if balance < -1 && r.Key > n.right.record.Key {
		return leftRotate(n)
	}
	if balance > 1 && r.Key > n.left.record.Key {
		n.left = leftRotate(n.left)
		return rightRotate(n)
	}
	if balance < -1 && r.Key < n.right.record.Key {
		n.right = rightRotate(n.right)
		return leftRotate(n)
	}
	return n
}

func rightRotate(n *Node) *Node {

	x := n.left
	t2 := x.right
	x.right = n
	n.left = t2

	n.height = max(height(n.left), height(n.right)) + 1
	x.height = max(height(x.left), height(x.right)) + 1
	return x
}

func leftRotate(n *Node) *Node {
	y := n.right
	t2 := y.left

	y.left = n
	n.right = t2

	n.height = max(height(n.left), height(n.right)) + 1
	y.height = max(height(y.left), height(y.right)) + 1
	return y
}

func height(n *Node) int {
	if n == nil {
		return 0
	}
	return n.height
}

func getBalance(n *Node) int {
	if n == nil {
		return 0
	}
	return height(n.left) - height(n.right)
}
