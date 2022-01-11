package main

import (
	"fmt"
	"math"
)

type Node struct {
	key    int
	value  string
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

// put adds a new Node with the given key-value pair, or replaces the value if key exists.
func (bt *BinaryTree) put(k int, v string) {
	bt.root = bt._put(bt.root, k, v)
}

// _put adds or updates a node to the subtree rooted at node
func (bt *BinaryTree) _put(node *Node, k int, v string) *Node {
	if node == nil {
		return &Node{key: k, value: v}
	}
	if k == node.key {
		node.value = v
		return node
	}

	if k < node.key {
		node.left = bt._put(node.left, k, v)
		node = resolveLeftLeaning(node)
	} else {
		node.right = bt._put(node.right, k, v)
		node = resolveRightLeaning(node)
	}

	node.computeHeight()

	return node
}

// remove removes a value from BinaryTree
func (bt *BinaryTree) remove(key int) {
	bt.root = bt._remove(bt.root, key)
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

func (bt *BinaryTree) _remove(node *Node, key int) *Node {
	if node == nil {
		return nil
	}

	switch {
	case key < node.key:
		node.left = bt._remove(node.left, key)
		node = resolveRightLeaning(node)
	case key > node.key:
		node.right = bt._remove(node.right, key)
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

func (bt *BinaryTree) contains(key int) bool {
	has := bt.get(key)

	return has != ""
}

// _contains iterates over a BinaryTree, returning true if it contains the given values.
func (bt *BinaryTree) get(target int) string {
	node := bt.root
	for node != nil {
		if target == node.key {
			return node.value
		}
		if target < node.key {
			node = node.left
		} else {
			node = node.right
		}
	}

	return ""
}

func (bt *BinaryTree) iterate() {
	bt._inorder(bt.root)
}

func (bt *BinaryTree) _inorder(node *Node) {
	if node == nil {
		return
	}

	bt._inorder(node.left)
	fmt.Println(node.key, node.value)
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
	bTree := BinaryTree{}

	M := map[int]string{
		5:  "Boron",
		20: "Calcium",
		53: "Iodine",
		58: "Cerium",
		76: "Osmium",
	}

	fmt.Println(bTree)

	for k, v := range M {
		bTree.put(k, v)
	}

	bTree.put(79, "Gold")

	bTree.iterate()

	fmt.Printf("has 19: %v\n", bTree.contains(20))
	fmt.Printf("has 99: %v\n", bTree.contains(99))

	bTree.remove(20)
	fmt.Printf("has 19: %v\n", bTree.contains(19))

	fmt.Println(bTree.root.height)

	fmt.Println(bTree.get(79))
}
