package rbst

func (n *Node) Replace(v Interface) {
	n.replace(newNode(v))
}

func (n *Node) replace(child *Node) {
	child.parent = n.parent
	if n == n.parent.left {
		n.parent.left = child
	} else {
		n.parent.right = child
	}
}
