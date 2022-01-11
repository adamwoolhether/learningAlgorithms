package main

import (
	"fmt"
	"math"
)

type Node struct {
	value  int
	left   *Node
	right  *Node
	height int
}

// heightDifference computes the difference in height between the left and right subtree
func (n *Node) heightDifference() int {
	var leftHeight, rightHeight int
	if n.left != nil {
		leftHeight = n.left.height
	} else {
		leftHeight = -1
	}

	if n.right != nil {
		rightHeight = n.right.height
	} else {
		rightHeight = -1
	}

	return leftHeight - rightHeight
}

// computeHeight updates the height of a node. It assumes that child nodes (if any) report their correct height.
func (n *Node) computeHeight() {
	var leftHeight, rightHeight int

	if n.left != nil {
		leftHeight = n.left.height
	} else {
		leftHeight = -1
	}

	if n.right != nil {
		rightHeight = n.right.height
	} else {
		rightHeight = -1
	}

	n.height = 1 + int(math.Max(float64(leftHeight), float64(rightHeight)))
}

type BinaryTree struct {
	root *Node
}

func (bt *BinaryTree) insert(val int) {
	bt.root = bt._insert(bt.root, val)
}

func (bt *BinaryTree) _insert(node *Node, val int) *Node {
	if node == nil {
		return &Node{value: val}
	}

	if val <= node.value {
		node.left = bt._insert(node.left, val)
		node = resolveLeftLeaning(node)
	} else {
		node.right = bt._insert(node.right, val)
		node = resolveRightLeaning(node)
	}

	node.computeHeight()

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

// remove removes a value from BinaryTree
func (bt *BinaryTree) remove(val int) {
	bt.root = bt._remove(bt.root, val)
}

func (bt *BinaryTree) _removeMin(node *Node) *Node {
	if node.left == nil {
		return node.right
	}

	node.left = bt._removeMin(node.left)
	node = resolveRightLeaning(node)

	node.computeHeight()

	return node
}

// _remove removes a value from a subtree rooted at node and returns the resulting subtree.
func (bt *BinaryTree) _remove(node *Node, val int) *Node {
	if node == nil {
		return nil
	}

	switch {
	case val < node.value:
		node.left = bt._remove(node.left, val)
		node = resolveRightLeaning(node)
	case val > node.value:
		node.right = bt._remove(node.right, val)
		node = resolveLeftLeaning(node)
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
		node = resolveLeftLeaning(node)
	}

	node.computeHeight()

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

// resolveLeftLeaning re-balances and return a new node root for nodes that are left-learning.
func resolveLeftLeaning(node *Node) *Node {
	if node.heightDifference() == 2 {
		if node.left.heightDifference() >= 0 {
			node = rotateRight(node)
		} else {
			node = rotateLeftRight(node)
		}
	}

	return node
}

// resolveRightLeaning re-balances and return a new node root for nodes that are right-learning.
func resolveRightLeaning(node *Node) *Node {
	if node.heightDifference() == -2 {
		if node.right.heightDifference() <= 0 {
			node = rotateLeft(node)
		} else {
			node = rotateRightLeft(node)
		}
	}

	return node
}

func rotateRight(node *Node) *Node {
	newRoot := node.left
	grandson := newRoot.right
	node.left = grandson
	newRoot.right = node

	node.computeHeight()

	return newRoot
}

func rotateLeft(node *Node) *Node {
	newRoot := node.right
	grandson := newRoot.left
	node.right = grandson
	newRoot.left = node

	node.computeHeight()

	return newRoot
}

func rotateLeftRight(node *Node) *Node {
	child := node.left
	newRoot := child.right
	grand1 := newRoot.left
	grand2 := newRoot.right
	child.right = grand1
	node.left = grand2

	newRoot.left = child
	newRoot.right = node

	child.computeHeight()
	node.computeHeight()

	return newRoot
}

func rotateRightLeft(node *Node) *Node {
	child := node.right
	newRoot := child.left
	grand1 := newRoot.left
	grand2 := newRoot.right
	child.left = grand2
	node.right = grand1

	newRoot.left = node
	newRoot.right = child

	child.computeHeight()
	node.computeHeight()

	return newRoot
}

func main() {
	// bTree := NewBinaryTree()
	bTree := BinaryTree{}

	A := []int{19, 14, 15, 53, 58, 3, 26}

	fmt.Println(bTree)

	for _, i := range A {
		bTree.insert(i)
	}

	bTree.insert(29)
	bTree.insert(27)

	bTree.iterate()

	fmt.Printf("has 19: %v\n", bTree._contains(19))
	fmt.Printf("has 99: %v\n", bTree._contains(99))

	bTree.remove(19)
	fmt.Printf("has 19: %v\n", bTree._contains(19))

	fmt.Println(bTree.root.height)

	avlTreeArray := []int{10, 30, 50}
	avlTree := BinaryTree{}

	for _, i := range avlTreeArray {
		avlTree.insert(i)
	}

	avlTree.iterate()
}
