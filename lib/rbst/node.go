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

func (n *Node) String() string {
	return toString([]*Node{n})
}

func newNode(v Interface) *Node {
	return &Node{
		Element: v,
		left:    leaf(),
		right:   leaf(),
	}
}

func (n *Node) flatten() []Interface {
	if n == nil || isLeaf(n) {
		return []Interface{}
	}
	slice := n.left.flatten()
	slice = append(slice, n.Element)
	slice = append(slice, n.right.flatten()...)
	return slice
}

func toString(nodes []*Node) (str string) {
	empty := true
	for _, node := range nodes {
		if node != nil {
			empty = false
		}
	}
	if empty {
		return ""
	}

	var next []*Node
	buf := bytes.NewBufferString("\n")

	defer func() {
		r := recover()
		if r != nil {
			fmt.Fprint(buf, "PANIC")
			str = buf.String()
		}
	}()

	for i := 0; i < len(nodes); i++ {
		if i != 0 {
			fmt.Fprint(buf, "|")
		}
		n := nodes[i]
		if n == nil {
			fmt.Fprint(buf, "<nil>")
			next = append(next, nil, nil)
		} else if isLeaf(n) {
			fmt.Fprint(buf, "LEAF")
			next = append(next, nil, nil)
		} else {
			next = append(next, n.left, n.right)
			fmt.Fprintf(buf, "[%v]", n.Element)
			if n.isRed {
				fmt.Fprint(buf, "R")
			} else {
				fmt.Fprint(buf, "B")
			}
		}
	}

	fmt.Fprint(buf, toString(next))
	return buf.String()
}
