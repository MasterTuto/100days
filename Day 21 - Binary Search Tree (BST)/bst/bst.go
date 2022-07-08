package bst

import (
	"golang.org/x/exp/constraints"
)

type BST[T constraints.Ordered] struct {
	count int
	root  *BSTNode[T]
}

type BSTNode[T constraints.Ordered] struct {
	left  *BSTNode[T]
	value *T
	right *BSTNode[T]
}

func New[T constraints.Ordered]() *BST[T] {
	return &BST[T]{
		count: 0,
		root:  nil,
	}
}

func (bst BST[T]) Count() int {
	return bst.count
}

func (bst BST[T]) Empty() bool {
	return bst.root == nil
}

func (bst *BST[T]) Insert(value T) {
	if bst.Empty() {
		bst.root = &BSTNode[T]{
			left:  nil,
			value: &value,
			right: nil,
		}

		bst.count++

		return
	}

	currentNode := bst.root
	for {
		if value < *currentNode.value {
			if currentNode.left == nil {
				currentNode.left = &BSTNode[T]{left: nil, value: &value, right: nil}
				break
			} else {
				currentNode = currentNode.left
			}
		} else {
			if currentNode.right == nil {
				currentNode.right = &BSTNode[T]{left: nil, value: &value, right: nil}
				break
			} else {
				currentNode = currentNode.right
			}
		}
	}

	bst.count += 1
}

func preOrder[T constraints.Ordered](c chan T, node *BSTNode[T], isRoot bool) {
	c <- *node.value

	if node.left != nil {
		preOrder(c, node.left, false)
	}

	if node.right != nil {
		preOrder(c, node.right, false)
	}

	if isRoot {
		close(c)
	}
}

func inOrder[T constraints.Ordered](c chan T, node *BSTNode[T], isRoot bool) {
	if node.left != nil {
		inOrder(c, node.left, false)
	}

	c <- *node.value

	if node.right != nil {
		inOrder(c, node.right, false)
	}

	if isRoot {
		close(c)
	}
}

func postOrder[T constraints.Ordered](c chan T, node *BSTNode[T], isRoot bool) {

	if node.left != nil {
		postOrder(c, node.left, false)
	}

	if node.right != nil {
		postOrder(c, node.right, false)
	}

	c <- *node.value

	if isRoot {
		close(c)
	}
}

func (bst BST[T]) PreOrder() <-chan T {
	c := make(chan T)

	if bst.Empty() {
		close(c)
	} else {
		go preOrder(c, bst.root, true)
	}

	return c
}

func (bst BST[T]) InOrder() <-chan T {
	c := make(chan T)

	if bst.Empty() {
		close(c)
	} else {
		go inOrder(c, bst.root, true)
	}

	return c
}

func (bst BST[T]) PostOrder() <-chan T {
	c := make(chan T)

	if bst.Empty() {
		close(c)
	} else {
		go postOrder(c, bst.root, true)
	}

	return c
}

func (bst BST[T]) Contains(value T) bool {
	for node := bst.root; node != nil; {

		if value == *node.value {
			return true
		} else if value < *node.value {
			node = node.left
		} else {
			node = node.right
		}
	}

	return false
}

func searchSmaller[T constraints.Ordered](node *BSTNode[T]) *BSTNode[T] {
	cur := node
	for ; cur.left != nil; cur = cur.left {
	}

	return cur
}

func (bst *BST[T]) removeNode(node, parent *BSTNode[T], isNodeLeft bool) {
	var substitute *BSTNode[T]

	if node.IsLeaf() {
		substitute = nil
	} else if node.left != nil && node.right != nil {
		substitute = searchSmaller(node.right)
	} else if node.left != nil {
		substitute = node.left
	} else if node.right != nil {
		substitute = node.right
	}

	if parent == nil {
		if substitute == nil {
			bst.root = nil
		} else {
			bst.root.value = substitute.value
		}
		return
	}

	if isNodeLeft {
		if substitute != nil {
			parent.left.value = substitute.value
		} else {
			parent.left = nil
		}
	} else {
		if substitute != nil {
			parent.right.value = substitute.value
		} else {
			parent.right = nil
		}
	}
}

func (bst *BST[T]) Remove(value T) bool {

	if bst.Empty() {
		return false
	}

	var parentNode *BSTNode[T] = nil
	node := bst.root

	for node != nil {
		if value == *node.value {
			break
		} else if value < *node.value {
			parentNode = node
			node = node.left
		} else {
			parentNode = node
			node = node.right
		}
	}

	if node == nil {
		return false
	}

	if parentNode != nil {
		bst.removeNode(node, parentNode, value < *parentNode.value)
	} else {
		bst.removeNode(node, parentNode, false)
	}

	bst.count--
	return true
}

func (n BSTNode[T]) IsLeaf() bool {
	return n.left == nil && n.right == nil
}
