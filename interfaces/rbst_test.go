package interfaces

import (
	"bytes"
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"testing"
	"testing/quick"
)

type z int

func (lhs z) Less(rhs interface{}) bool {
	return lhs < rhs.(z)
}

type zs []z

func (s zs) Len() int           { return len(s) }
func (s zs) Less(i, j int) bool { return s[i] < s[j] }
func (s zs) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

type r float64

func (lhs r) Less(rhs interface{}) bool {
	return lhs < rhs.(r)
}

type rs []r

func (s rs) Len() int           { return len(s) }
func (s rs) Less(i, j int) bool { return s[i] < s[j] }
func (s rs) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func TestRbst(t *testing.T) {
	rng := rand.New(rand.NewSource(0))
	testSizes := []int{1, 2, 5, 10, 16, 100, 1024, 99999}
	if testing.Short() {
		testSizes = []int{1, 2, 5, 10, 16}
	}

	var root *Node

	for _, size := range testSizes {
		t.Logf("%d Integers", size)
		integers := make(zs, size)
		for i := 0; i < size; i++ {
			val, ok := quick.Value(reflect.TypeOf(z(0)), rng)
			if ok {
				integers[i] = val.Interface().(z)
				if i == 0 {
					root = NewNode(integers[i])
				} else {
					root = Insert(root, NewNode(integers[i]))
				}
			} else {
				t.Fatalf("test: Error generating values")
			}
		}
		t.Log(toString([]*Node{root}))
		if !blackValidation(root) {
			t.Errorf("uneven distribution of black child nodes")
		}
		if !redValidation(root) {
			t.Errorf("some red children don't have black children")
		}
		sort.Sort(integers)
		zseq := collect(root)
		misplaced := false
		for i := range zseq {
			if zseq[i].(z) != integers[i] {
				misplaced = true
			}
		}
		if misplaced {
			t.Errorf("Elements out of place\nE: %v\nA: %v", integers, zseq)
		}

		t.Logf("%d Floats", size)
		floats := make(rs, size)
		for i := 0; i < size; i++ {
			val, ok := quick.Value(reflect.TypeOf(r(0)), rng)
			if ok {
				floats[i] = val.Interface().(r)
				if i == 0 {
					root = NewNode(floats[i])
				} else {
					root = Insert(root, NewNode(floats[i]))
				}
			} else {
				t.Fatalf("test: Error generating values")
			}
		}
		t.Log(toString([]*Node{root}))
		if !blackValidation(root) {
			t.Errorf("uneven distribution of black child nodes")
		}
		if !redValidation(root) {
			t.Errorf("some red children don't have black children")
		}
		sort.Sort(floats)
		rseq := collect(root)
		for i := range rseq {
			if rseq[i].(r) != floats[i] {
				t.Errorf("Element %f out of place; expected %f", rseq[i], floats[i])
			}
		}
	}
}

func blackValidation(root *Node) bool {
	// Every path from a given node to any leaf node
	// contains the same number of black nodes
	min, max := countBlackPaths(root)
	return min == max
}

func countBlackPaths(root *Node) (min, max int) {
	if root == nil || isLeaf(root) {
		return 0, 0
	}
	leftMin, leftMax := countBlackPaths(root.left)
	rightMin, rightMax := countBlackPaths(root.right)

	if !root.left.isRed {
		leftMin++
		leftMax++
	}

	if !root.right.isRed {
		rightMin++
		rightMax++
	}

	if leftMin < rightMin {
		min = leftMin
	} else {
		min = rightMin
	}
	if leftMax < rightMax {
		max = rightMax
	} else {
		max = leftMax
	}

	return
}

func redValidation(root *Node) bool {
	// If any node is red, both its children are black
	if root.isRed {
		if root.left != nil && root.left.isRed {
			return false
		}
		if root.right != nil && root.right.isRed {
			return false
		}
	}
	if root.left != nil && !redValidation(root.left) {
		return false
	}
	if root.right != nil && !redValidation(root.right) {
		return false
	}
	return true
}

func collect(root *Node) []Interface {
	if root == nil || isLeaf(root) {
		return []Interface{}
	}
	slice := collect(root.left)
	slice = append(slice, root.Element)
	slice = append(slice, collect(root.right)...)
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
