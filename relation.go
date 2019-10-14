package rbst

func (n *Node) grandparent() *Node {
	p := n.parent
	if p == nil {
		return nil
	}
	return p.parent
}

func (n *Node) sibling() *Node {
	p := n.parent
	if p == nil {
		return nil
	}
	if n == p.left {
		return p.right
	}
	return p.left
}

func (n *Node) aunt() *Node {
	p := n.parent
	if p == nil {
		return nil
	}
	return p.sibling()
}
