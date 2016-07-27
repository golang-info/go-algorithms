package binarysearchtree

import (
	"errors"
)

// TODO: add support for any type, not just int
type Node struct {
	value       int   // Value the node stores
	left, right *Node // pointers to the left and right children of the key
}

// Insert will add a new node to the tree whose root is Node
func (n *Node) Insert(newNode *Node) {
	// New node goes to the right part
	if n.value < newNode.value {
		if n.right == nil {
			n.right = newNode
		} else {
			n.right.Insert(newNode)
		}
	}

	// New node goes to the left part
	if n.value > newNode.value {
		if n.left == nil {
			n.left = newNode
		} else {
			n.left.Insert(newNode)
		}
	}
}

func (n *Node) FindMin() (int, error) {
	if n == nil {
		return 0, errors.New("Empty tree has no Min element.")
	}
	for n.left != nil {
		n = n.left
	}
	return n.value, nil
}

func (n *Node) FindMax() (int, error) {
	if n == nil {
		return 0, errors.New("Empty tree has no Max element.")
	}
	for n.right != nil {
		n = n.right
	}
	return n.value, nil
}

func (n *Node) Delete(element int) (*Node, bool) {
	deleted := false

	if n == nil {
		return n, false
	}

	if n.value < element {
		n.right, deleted = n.right.Delete(element)
	} else if n.value > element {
		n.left, deleted = n.left.Delete(element)
	} else if n.left != nil && n.right != nil {
		candidate, _ := n.right.FindMin()
		n.value = candidate
		n.right, deleted = n.right.Delete(candidate)
		deleted = true
	} else {
		if n.left == nil {
			n = n.right
		} else {
			n = n.left
		}
		deleted = true
	}
	return n, deleted
}

// Walk function calls `f` on every node value
// FIXME: Which is a better as parameter of `f`, `Node` or `node.value`?
func (n *Node) Walk(f func(int)) {
	if n == nil {
		return
	}

	n.left.Walk(f)
	f(n.value)
	n.right.Walk(f)
}

// BinarySearchTree store a root and the tree node numbers.
type BinarySearchTree struct {
	root      *Node // pointer to root of the tree
	nodes_num int   // nodes number: how many nodes are in the tree
}

// Return an initial new tree
func New() *BinarySearchTree {
	return &BinarySearchTree{}
}

// Return how many nodes are in the tree
func (b *BinarySearchTree) Nodes() int {
	return b.nodes_num
}

// Insert an element into the tree
func (b *BinarySearchTree) Insert(element int) {
	node := Node{value: element}

	if b.root == nil {
		b.root = &node
	} else {
		b.root.Insert(&node)
	}
	b.nodes_num += 1
}

// Calls `f` on every node
func (b *BinarySearchTree) Walk(f func(int)) {
	b.root.Walk(f)
}

// Check if an element exists in a tree
func (b *BinarySearchTree) Contains(element int) bool {
	n := b.root
	for n != nil {
		if n.value == element {
			return true
		}

		if n.value < element {
			n = n.right
		} else {
			n = n.left
		}
	}

	return false
}

// Return if the tree is an empty tree
func (b *BinarySearchTree) IsEmpty() bool {
	return b.nodes_num == 0
}

// Find and return a pointer to the node whose value equals to `element`
func (b *BinarySearchTree) Find(element int) *Node {
	n := b.root

	for n != nil {
		if n.value == element {
			return n
		}

		if n.value < element {
			n = n.right
		} else {
			n = n.left
		}
	}

	return nil
}

// Find and return the smallest element in tree
func (b *BinarySearchTree) FindMin() (int, error) {
	return b.root.FindMin()
}

// Find and return the biggest element in tree
func (b *BinarySearchTree) FindMax() (int, error) {
	return b.root.FindMax()
}

// Delete the first appearence of the element
// There are several cases:
// - Node is a leaf, simply delete it
// - Node has one child, point parent to the child, and delete the node
// - Node has two children, replace the leftmost node of the right part with current node,
//   then delete the leftmost node
func (b *BinarySearchTree) Delete(element int) {
	n, deleted := b.root.Delete(element)
	b.root = n
	if deleted {
		b.nodes_num--
	}
}