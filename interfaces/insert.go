package interfaces

func insert(root, n *Node) *Node {
	insertRecurse(root, n)
	insertRepairTree(n)

	for root = n; root.parent != nil; {
		root = root.parent
	}
	return root
}

func insertRecurse(root, n *Node) {
	if root != nil {
		if n.Element.Less(root.Element) {
			if !isLeaf(root.left) {
				insertRecurse(root.left, n)
				return
			}
			root.left = n
		} else {
			if !isLeaf(root.right) {
				insertRecurse(root.right, n)
				return
			}
			root.right = n
		}
	}

	n.parent = root
	n.left = leaf()
	n.right = leaf()
	n.isRed = true
}

func insertRepairTree(n *Node) {
	if n.parent == nil {
		n.isRed = false
	} else if !n.parent.isRed {
		return
	} else if n.aunt() != nil && n.aunt().isRed {
		n.parent.isRed = false
		n.aunt().isRed = false
		n.grandparent().isRed = true
		insertRepairTree(n.grandparent())
	} else {
		p := n.parent
		g := n.grandparent()
		if n == p.right && p == g.left {
			p.rotateLeft()
			n = n.left
		} else if n == p.left && p == g.right {
			p.rotateRight()
			n = n.right
		}

		p = n.parent
		g = n.grandparent()

		if n == p.left {
			g.rotateRight()
		} else {
			g.rotateLeft()
		}
		p.isRed = false
		g.isRed = true
	}
}
