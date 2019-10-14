package rbst

func (n *Node) visitBreadthFirst(f func([]*Node)) {
	if n == nil || isLeaf(n) {
		return
	}
	visitTier([]*Node{n}, f)
}

func visitTier(tier []*Node, f func([]*Node)) {
	empty := true
	for _, node := range tier {
		if node != nil {
			empty = false
		}
	}
	if empty {
		return
	}

	f(tier)

	var next []*Node
	for i := 0; i < len(tier); i++ {
		n := tier[i]
		if n == nil || isLeaf(n) {
			next = append(next, nil, nil)
			continue
		}

		next = append(next, n.left, n.right)
	}
	visitTier(next, f)
}

func (n *Node) visitDepthFirst(f func(*Node)) {
	if n == nil || isLeaf(n) {
		return
	}
	if n.left != nil {
		n.left.visitDepthFirst(f)
	}
	f(n)
	if n.right != nil {
		n.right.visitDepthFirst(f)
	}
}
