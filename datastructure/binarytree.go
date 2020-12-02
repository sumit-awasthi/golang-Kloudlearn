//binary search tree
package main

type Tree struct {
	node *Node
}
//root value
func (t *Tree) insert(value int) *Tree {
	if t.node == nil {
		t.node = &Node{value: value}
	} else {
		t.node.insert(value)
	}

	return t
}

type Node struct {
	value int
	left  *Node
	right *Node
}

func (n *Node) insert(value int) {
	if value <= n.value {
		if n.left == nil {
			n.left = &Node{value: value}
		} else {
			n.left.insert(value)
		}
	} else {
		if n.right == nil {
			n.right = &Node{value: value}
		} else {
			n.right.insert(value)
		}
	}
}
//  traversal  or searching for an element
func (n *Node) exists(value int) bool {
	if n == nil {
		return false
	}

	if n.value == value {
		return true
	}

	println("value", n.value)
	if value <= n.value {
		return n.left.exists(value)
	} else {
		return n.right.exists(value)
	}
}

func printNode(n *Node) {
	if n == nil {
		return
	}
	println(n.value)
	printNode(n.left)
	printNode(n.right)
}

func main() {
	t := &Tree{}
	t.insert(10)
	t.insert(8)
	t.insert(20)
	t.insert(9)
	t.insert(0)
	t.insert(15)
	t.insert(25)

	printNode(t.node)
	println(t.node.exists(26)) //searching for an element in the tree
}
