# Binary Trees
Binary trees are _recursive_ data structures, as they refer to other binary tree structures within

## Binary Search Tree
Store collections of values with efficient search, insert, and remove operations.  
They help solve challenges introduced by large fixed arrays.  
A binary search tree placed side by side with a sorted array with the same values:
![](binarySearchTreeArrayComparisoin.gif)

This exemplifies two global constraints of binary search trees:
* If node _n_, has a left subtree, all values in the subtree as <= _n_
* If node _n_, has a right subtree, all values in the subtree are >= _n_

A _leaf_ is a node without any subtrees. As such, computer-science tress can be considered to be upside down.  
_Parents_ and _descendants_ are other words used to describe tress.

The order in which an array's values are inserted into the binary tree determines the tree's structure. For example,
inserting 29 into the above tree will result in the following:

![](insert.gif)

The worse case scenario for a binary search tree is revealed when all values are inserted in increasing order:

![](worstCase,gif)

### Removing values from a Binary Search Tree
Removing values from binary search trees is trickier than a linked list. What if the value is the current root node?
The figure below demonstrates the two possible outcomes:

![](removeRootNode.gif)

* Option 1: Find & remove maximum value in the left subtree, use it for new root
* Option 2: Find & remove minimum value in the right subtree, use it for new root

We'll use option 2, here, demonstrated below:

![](removeNode2.gif)

This brings rise to the **subproblem** of removing the min value in a given subtree. **the minimum value in a subtree cannot
have a left child**, otherwise it wouldn't be the smalled value.  
We implement this with our `_remove_min` function.

### Traversing a Binary Tree
_Traversal_ is when we return the values of a binary tree in ascending order.
We can't use the same approach as in liked lists, because we have both right and left references to follow.

## Analyzing Performance of Binary Search Trees
The height of a tree determines the speed of search, insert, and remove operations. _Height_ of a node is the number of 
left/right references needed to get to the most distant descendant leaf node (a leaf node has a height of 0).  
`-1` is used to signify the height of a nonexistent binary node.

In the worse case, the number of nodes traversed in a search operation is based on the height of the tree.

Generally, the runtime performance of insert and search is `O(h`, where `h` = height of the tree.  
Removing a node takes three steps:
1. Locate the node with the value to be removed
2. Locate the minimum value in the node's right subtree
3. Remove the value from the right of the subtree

In the worst case, removal is `3xh`, but it is also generally `O(h)` 

As mentioned earlier, a binary search tree's structure is dependent on the order that the values are inserted/removed.
This means we need a strategy for the tree to determin when it is performing poorly.

## Self-Balancing Binary Trees
AVL trees guarantee that height differences of any node(height of the node's left subtree minus height of the node's right
subtree) is -1, 0, or 1.
Each node stores its height, and inserting a node means the nodes of the affected nodes are computed to ensure that an
unbalanced tree is detected immediately.

After we implement our `computeHeight` and `heightDifference` methods, we add a node with value `27` to the tree:

![](add27ToAvlTree.gif)

And as the recursive `_insert` unwinds, the additional height is added to the ancestors.

_Node rotation_ is conducted when the tree detects an imblanace, which means the height difference exceeds a range of -1 - 1.

Node rotation:  
![](nodeRotation.gif)

Here we can see that the node with a value of 30 breaks the balance that AVL trees promise, it's then conducts a
_rotate right_ to balance itself.

An tree with just three values has four possible unbalanced scenarios:
* Left-left: Demonstrated above, where only a _rotate right_ is needed
* Right-right: The mirror iamge of the above, needing only a _rotate left_
* Left-right: Needs to be reblananced in two steps. First a _rotate left_ which results in a Left-left, and then a
_rotate right_. This two step operation is also known as a _rotate right-left_.
* Right-left: the mirror of Left-right. a _rotate right-left_ will rebalance the tree.

* ![](scendarios.gif)

These re-balance scenarios all have a corresponding helper function to maintain the AVL tree's balance.

## Analyzing performance of Self-Balancing Trees
The worse cas, in regards to rotation methods is `log (N) rotations`. Otherwise, performance for search, insert, and
remove is `O(log N)`.

## Using Binary Trees as (key, value) Symbol Table
We can also use this structure to store key-value pairings, similar to chapter three. Replease the insert() and contains()
functions with put() and get().
It should be noted, however, that this is far less performant than the hash table implementations.

## Using Binary Trees as a Priority Queue
Two benefits to used a binary search tree as a priority queue:
* The structure is dynamic(array-based must first allocate the storage in advance)
* We can iteraate over entries in priority order _without_ dequeing the values.

The `priority` field will be used to locate nodes within the tree. In max binary heap, the highest priority queue can be
located in `O(1)` constant time, because its at the first index. Using binary trees, the top priority node is at the 
most-right node, giving it a runtime of `O(log N` if tree is balanced.

When using binary trees for storage, the only value to remove is the one with the highest priority, done with our 
`_removeMax()` func.