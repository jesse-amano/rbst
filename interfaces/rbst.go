// Package interfaces provides a red-black binary search tree
// that uses an interface to describe the elements it contains.
// A potential drawback is that callers need to perform type assertions
// to retrieve elements of an appropriate type, so some boilerplate wrappers
// are needed to make calling code type-safe.
package interfaces

// A type, typically an element, that satisfies rbst.Interface can be stored and sorted by the
// types and functions in this package. The methods require that the elements of the collection
// have at least a weak ordering.
type Interface interface {
	Less(v interface{}) bool
	// Optionally, we could also require a method like
	// IsSameType(v interface{}) bool
	// to do some checking on behalf of the caller,
	// but the burden would still be on the caller to perform
	// type assertions since the type information is unfortunately
	// stateful in this case
}

// An rbst.RBST holds the root node of a red-black binary search tree.
type RBST struct {
	root *Node
}

func (r *RBST) Insert(v Interface) {
	if r.root == nil {
		r.root = newNode(v)
		return
	}
	r.root = insert(r.root, newNode(v))
}

func (r *RBST) Flatten() []Interface {
	if r.root == nil || isLeaf(r.root) {
		return []Interface{}
	}
	return r.root.flatten()
}

func (r *RBST) String() string {
	return r.root.String()
}
