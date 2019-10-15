// package kv demonstrates how complex types may be passed through simple, legible wrapper packages
// to interact with general-purpose rbst logic.
package kv

import (
	"errors"
	"strings"

	"github.com/google/uuid"
	"github.com/jesse-amano/rbst"
)

var (
	ErrInsertNil = errors.New("rbst: Cannot insert nil value.")
)

// Value is an example of a complicated node value in a red/black tree.
// Its identifier is treated as the node key, so lookups by identifier
// can take advantage of the search tree's structure.
type Value struct {
	id        *uuid.UUID
	numbers   []float64
	documents map[string]interface{}
}

// Less compares the identifiers of two Values.
// Accepting an interface{} type and not checking it during conversion
// makes this theoretically vulnerable to panics, but other functions
// in this package guarantee Less will never be called with a value
// not of type Value.
func (v Value) Less(o interface{}) bool {
	return strings.Compare(v.id.String(), o.(Value).id.String()) < 0
}

// String satisfies the common and conventional "stringer" interface.
// In particular, it makes a pretty-printed KeyValueTree much prettier
// by restricting output to just the identifier.
// It's still a little ugly. Sorry. That is kind of how it goes when
// elements are best identified by long strings.
func (v Value) String() string {
	if v.id == nil {
		return "ø"
	}
	return v.id.String()
}

// KeyValueTree is a search tree containing Values using their identifier field
// as a search key.
type KeyValueTree struct {
	tree rbst.RBST
}

// Find finds the value identified by id in O(log(n)) time, thanks to
// binary search tree structure.
func (t *KeyValueTree) Find(id *uuid.UUID) *Value {
	if id == nil {
		return nil
	}
	val := t.tree.Find(&Value{id: id})
	if val == nil {
		return nil
	}
	return val.(*Value)
}

// Flatten returns a slice containing all elements of the tree,
// sorted in ascending order.
func (t *KeyValueTree) Flatten() []*Value {
	kvs := t.tree.Flatten()
	vals := make([]*Value, len(kvs))
	for i := range vals {
		vals[i] = kvs[i].(*Value)
	}
	return vals
}

// Insert adds an element to the tree. Its signature guarantees compile-
// time checking that all elements added to the tree are non-nil and of
// the same type. This isn't necessarily a requirement for all search trees,
// but it's a common requirement for many. Fortunately, Go1 does not force us
// to choose until we know which case we are in.
func (t *KeyValueTree) Insert(v *Value) error {
	if v == nil || v.id == nil {
		return ErrInsertNil
	}
	t.tree.Insert(v)
	return nil
}

// String returns the tree represented as a string.
func (t *KeyValueTree) String() string {
	return t.tree.String()
}

// VisitBreadthFirst applies a function to each row of the tree,
// starting from the root.
func (t *KeyValueTree) VisitBreadthFirst(f func([]*Value)) {
	t.tree.VisitBreadthFirst(func(s []*rbst.Node) {
		vals := make([]*Value, len(s))
		for i := range vals {
			vals[i] = s[i].Element.(*Value)
		}
		f(vals)
	})
}

// VisitDepthFirst applies a function to each node of the tree,
// starting from the left and working toward the right.
func (t *KeyValueTree) VisitDepthFirst(f func(*Value)) {
	t.tree.VisitDepthFirst(func(n *rbst.Node) {
		if n == nil {
			return
		}
		f(n.Element.(*Value))
	})
}
