package main

import (
	"fmt"
)

// Node defines the structure used to build a BinaryTree.
type Node struct {
	value int
	left  *Node
	right *Node
}

type BinaryTree struct {
	root *Node
}

// insert inserts a value into BinaryTree.
func (bt *BinaryTree) insert(val int) {
	bt.root = bt._insert(bt.root, val)
}

// _insert sets the bt.root as the subtree that results when inserting val in the subtree rooted at bt.root.
func (bt *BinaryTree) _insert(node *Node, val int) *Node {
	if node == nil {
		return &Node{value: val}
	}

	if val <= node.value {
		node.left = bt._insert(node.left, val)
	} else {
		node.right = bt._insert(node.right, val)
	}

	return node
}

// _contains iterates over a BinaryTree, returning true if it contains the given values.
func (bt *BinaryTree) _contains(target int) bool {
	node := bt.root
	for node != nil {
		if target == node.value {
			return true
		}
		if target < node.value {
			node = node.left
		} else {
			node = node.right
		}
	}

	return false
}

// _removeMin is a helper function that removes the minimum value in a subtree rooted at node.
func (bt *BinaryTree) _removeMin(node *Node) *Node {
	if node.left == nil {
		return node.right
	}

	node.left = bt._removeMin(node.left)

	return node
}

// remove removes a value from BinaryTree
func (bt *BinaryTree) remove(val int) {
	bt.root = bt._remove(bt.root, val)
}

// _remove removes a value from a subtree rooted at node and returns the resulting subtree.
func (bt *BinaryTree) _remove(node *Node, val int) *Node {
	if node == nil {
		return nil
	}

	switch {
	case val < node.value:
		node.left = bt._remove(node.left, val)
	case val > node.value:
		node.right = bt._remove(node.right, val)
	default:
		// handle if node is a leaf
		if node.left == nil {
			return node.right
		}
		if node.right == nil {
			return node.left
		}

		original := node // remember the original reference to node

		// find the smallest child in the right subtree
		node = node.right
		for node.left != nil {
			node = node.left
		}

		node.right = bt._removeMin(original.right)
		node.left = original.left // stitch the subtree back together
	}

	return node
}

func (bt *BinaryTree) iterate() {
	bt._inorder(bt.root)
}

func (bt *BinaryTree) _inorder(node *Node) {
	if node == nil {
		return
	}

	bt._inorder(node.left)
	fmt.Println(node.value)
	bt._inorder(node.right)
}

func main() {
	// bTree := NewBinaryTree()
	bTree := BinaryTree{}

	A := []int{19, 14, 15, 53, 58, 3, 26}

	fmt.Println(bTree)
	fmt.Println(bTree.root)

	for _, i := range A {
		bTree.insert(i)
		// fmt.Println(bTree.root.value)
	}

	bTree.iterate()

	fmt.Printf("has 19: %v\n", bTree._contains(19))
	fmt.Printf("has 99: %v\n", bTree._contains(99))

	bTree.remove(19)
	fmt.Printf("has 19: %v\n", bTree._contains(19))

	bTree.iterate()
}
