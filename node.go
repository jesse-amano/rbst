package rbst

import (
	"bytes"
	"fmt"
)

// Node methods are basically transliterated from Wikipedia
// as a way of getting at least a mostly-right answer without
// exhaustive testing
type Node struct {
	parent, left, right *Node
	isRed               bool
	Element             Interface
}

func newNode(v Interface) *Node {
	return &Node{
		Element: v,
		left:    leaf(),
		right:   leaf(),
	}
}

// String renders n and all its descendents into a string,
// with each tier on its own line.
func (n *Node) String() string {
	buf := bytes.NewBufferString("\n")
	yield := func(tier []*Node) {
		for i, n := range tier {
			if i != 0 {
				fmt.Fprint(buf, "|")
			}

			if n == nil {
				fmt.Fprint(buf, "<nil>")
			} else if isLeaf(n) {
				fmt.Fprint(buf, "LEAF")
			} else {
				fmt.Fprintf(buf, "[%v]", n.Element)
				if n.isRed {
					fmt.Fprint(buf, "R")
				} else {
					fmt.Fprint(buf, "B")
				}
			}
		}
		fmt.Fprint(buf, "\n")
	}
	n.visitBreadthFirst(yield)
	return buf.String()
}

// Flatten returns a slice of the descendents of n, sorted from least to greatest
// using the properties of the red/black tree.
func (n *Node) Flatten() []Interface {
	var s []Interface
	yield := func(m *Node) {
		s = append(s, m.Element)
	}
	n.visitDepthFirst(yield)
	return s
}

// Find returns the least descendent whose element is not Less than v.
// It may not be exactly equal to v, if v does not exactly exist as a descendent.
// The caller should perform its own equality checks if equality is required.
func (n *Node) Find(v Interface) *Node {
	if n.Element.Less(v) {
		if n.right == nil {
			return nil
		}
		return n.right.Find(v)
	}

	if n.left == nil {
		return n
	}

	candidate := n.left.Find(v)
	if candidate == nil {
		return n
	}
	return candidate
}
