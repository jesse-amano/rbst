package rbst

func leaf() *Node {
	var n Node
	return &n
}

func isLeaf(n *Node) bool {
	if n == nil {
		return true
	}
	return n.Element == nil
}
