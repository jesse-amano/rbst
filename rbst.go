// Package rbst provides a red/black binary search tree
// that uses an interface to describe the elements it contains.
// A potential drawback is that callers need to perform type assertions
// to retrieve elements of an appropriate type, so some boilerplate wrappers
// are needed to make calling code type-safe.
package rbst

// A type, typically an element, that satisfies rbst.Interface can be stored and sorted by the
// types and functions in this package. The methods require that the elements of the collection
// have at least a weak ordering.
type Interface interface {
	Less(v interface{}) bool
}

// An rbst.RBST holds the root node of a red-black binary search tree.
type RBST struct {
	root *Node
}

// Insert inserts v into r in order, maintaining the tree's red/black properties.
func (r *RBST) Insert(v Interface) {
	if r.root == nil {
		r.root = newNode(v)
		return
	}
	r.root = insert(r.root, newNode(v))
}

// VisitBreadthFirst visits each node in r, beginning with the root and proceeding
// left-to-right across each tier before visiting the next tier.
func (r *RBST) VisitBreadthFirst(f func([]*Node)) {
	if r.root == nil {
		return
	}
	r.root.visitBreadthFirst(f)
}

// VisitDepthFirst visits each node in r, beginning with the leftmost node and proceeding
// in order until reaching the rightmost node.
func (r *RBST) VisitDepthFirst(f func(*Node)) {
	if r.root == nil {
		return
	}
	r.root.visitDepthFirst(f)
}

// Flatten returns a slice of the elements of r, sorted from least to greatest
// using the properties of the red/black tree.
func (r *RBST) Flatten() []Interface {
	if r.root == nil || isLeaf(r.root) {
		return []Interface{}
	}
	return r.root.Flatten()
}

// String renders each tier of the red/black tree on its own line in a string, beginning with the root.
func (r *RBST) String() string {
	if r.root == nil {
		return leaf().String()
	}
	return r.root.String()
}

// Find returns the least descendent whose element is not Less than v.
// It may not be exactly equal to v, if v does not exactly exist as a descendent.
// The caller should perform its own equality checks if equality is required.
func (r *RBST) Find(v Interface) Interface {
	if r.root == nil {
		return nil
	}
	node := r.root.Find(v)
	if node == nil {
		return nil
	}
	return node.Element
}
