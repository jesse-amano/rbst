// Package interfaces provides a red-black binary search tree
// that uses an interface to describe the elements it contains.
// A potential drawback is that callers need to perform type assertions
// to retrieve elements of an appropriate type, so some boilerplate wrappers
// are needed to make calling code type-safe.
package interfaces

type Interface interface {
	Less(v interface{}) bool
	// Optionally, we could also require a method like
	// IsSameType(v interface{}) bool
	// to do some checking on behalf of the caller,
	// but the burden would still be on the caller to perform
	// type assertions since the type information is unfortunately
	// stateful in this case
}

// Node methods are basically transliterated from Wikipedia
// as a way of getting at least a mostly-right answer without
// exhaustive testing
type Node struct {
	parent, left, right *Node
	isRed               bool
	Element             Interface
}

func NewNode(v Interface) *Node {
	return &Node{
		Element: v,
		left:    leaf(),
		right:   leaf(),
	}
}

func Insert(root *Node, n *Node) *Node {
	insertRecurse(root, n)
	insertRepairTree(n)

	for root = n; root.parent != nil; {
		root = root.parent
	}
	return root
}

func Delete(n *Node) {
	var child *Node
	if isLeaf(n.right) {
		child = n.left
	} else {
		child = n.right
	}

	Replace(n, child)
	if !n.isRed {
		if child.isRed {
			child.isRed = false
		} else {
			deleteCase1(child)
		}
	}
}

func Replace(n, child *Node) {
	child.parent = n.parent
	if n == n.parent.left {
		n.parent.left = child
	} else {
		n.parent.right = child
	}
}
