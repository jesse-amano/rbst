package rbst

func (n *Node) rotateLeft() {
	nnew := n.right
	p := n.parent
	n.right = nnew.left // panics if nnew is nil
	nnew.left = n
	n.parent = nnew

	if n.right != nil {
		n.right.parent = n
	}
	if p != nil {
		if n == p.left {
			p.left = nnew
		} else {
			p.right = nnew
		}
	}
	nnew.parent = p
}

func (n *Node) rotateRight() {
	nnew := n.left
	p := n.parent
	n.left = nnew.right // panics if nnew is nil
	nnew.right = n
	n.parent = nnew

	if n.left != nil {
		n.left.parent = n
	}
	if p != nil {
		if n == p.left {
			p.left = nnew
		} else {
			p.right = nnew
		}
	}
	nnew.parent = p
}
