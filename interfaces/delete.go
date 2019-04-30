package interfaces

// Delete removes a node from its RBST.
func (n *Node) Delete() {
	var child *Node
	if isLeaf(n.right) {
		child = n.left
	} else {
		child = n.right
	}

	n.replace(child)
	if !n.isRed {
		if child.isRed {
			child.isRed = false
		} else {
			deleteCase1(child)
		}
	}
}

func deleteCase1(n *Node) {
	if n.parent != nil {
		deleteCase2(n)
	}
}

func deleteCase2(n *Node) {
	s := n.sibling()
	if s.isRed {
		n.parent.isRed = true
		s.isRed = false
		if n == n.parent.left {
			n.parent.rotateLeft()
		} else {
			n.parent.rotateRight()
		}
	}
	deleteCase3(n)
}

func deleteCase3(n *Node) {
	s := n.sibling()

	if !(n.parent.isRed || s.isRed || s.left.isRed || s.right.isRed) {
		s.isRed = true
		deleteCase1(n.parent)
	} else {
		deleteCase4(n)
	}
}

func deleteCase4(n *Node) {
	s := n.sibling()

	if n.parent.isRed && !s.isRed && !s.left.isRed && !s.right.isRed {
		s.isRed = true
		n.parent.isRed = false
	} else {
		deleteCase5(n)
	}
}

func deleteCase5(n *Node) {
	s := n.sibling()
	if !s.isRed {
		if n == n.parent.left && !s.right.isRed && s.left.isRed {
			s.isRed = true
			s.left.isRed = false
			s.rotateRight()
		} else if n == n.parent.right && !s.left.isRed && s.right.isRed {
			s.isRed = true
			s.right.isRed = false
			s.rotateLeft()
		}
	}
	deleteCase6(n)
}

func deleteCase6(n *Node) {
	s := n.sibling()
	s.isRed = n.parent.isRed
	n.parent.isRed = false

	if n == n.parent.left {
		s.right.isRed = false
		n.parent.rotateLeft()
	} else {
		s.left.isRed = false
		n.parent.rotateRight()
	}
}
