package main

import (
	"errors"
	"fmt"
	"log"
	"math"
)

// Node is used by BinaryTree to implement a priority queue.
type Node struct {
	priority int
	value    string
	left     *Node
	right    *Node
	height   int
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

// BinaryTree contains the root nodes and methods to manipulate the tree.
type BinaryTree struct {
	root *Node
}

// insert adds an entry into BinaryTree.
func (bt *BinaryTree) insert(p int, v string) {
	bt.root = bt._insert(bt.root, p, v)
}

// _insert adds a Node the the tree with (priority, value) pair.
func (bt *BinaryTree) _insert(node *Node, p int, v string) *Node {
	if node == nil {
		return &Node{priority: p, value: v}
	}

	if p <= node.priority {
		node.left = bt._insert(node.left, p, v)
		node = resolveLeftLeaning(node)
	} else {
		node.right = bt._insert(node.right, p, v)
		node = resolveRightLeaning(node)
	}

	node.computeHeight()

	return node
}

// iterate traverses the elements of BinaryTree.
func (bt *BinaryTree) iterate() {
	bt._inorder(bt.root)
}

// _inorder traveres the tree in order.
func (bt *BinaryTree) _inorder(node *Node) {
	if node == nil {
		return
	}

	bt._inorder(node.left)
	fmt.Println(node.priority, node.value)
	bt._inorder(node.right)
}

// PriorityQueue uses BinaryTree to store entries.
type PriorityQueue struct {
	tree *BinaryTree
	amt  int
}

// length returns the number of entries in PriorityQueue.
func (pq *PriorityQueue) length() int {
	return pq.amt
}

// isEmpty returns whether PriorityQueue is empty or not.
func (pq *PriorityQueue) isEmpty() bool {
	return pq.amt == 0
}

// enqueue adds an entry to the priority queue.
func (pq *PriorityQueue) enqueue(p int, v string) {
	pq.tree.insert(p, v)
	pq.amt++
}

// _removeMax removes the max entry, and ensures AVL ordering.
func (pq *PriorityQueue) _removeMax(node *Node) (string, *Node) {
	if node.right == nil {
		return node.value, node.left
	}

	var value string
	value, node.right = pq._removeMax(node.right)
	node = resolveLeftLeaning(node)

	node.computeHeight()

	return value, node
}

// dequeue removes and returns the value of the node with the highest priority.
func (pq *PriorityQueue) dequeue() (value string) {
	if pq.amt == 0 {
		log.Fatal(errors.New("PriorityQueue is empty"))
	}

	value, pq.tree.root = pq._removeMax(pq.tree.root)
	pq.amt--

	return value
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
