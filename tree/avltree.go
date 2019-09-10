/*
AVL Tree implementation
1.Insertation
2.Display
3.Deletion
*/

package main

import (
	"fmt"
)

// AVLTreeNode contains tree node detail
type AVLTreeNode struct {
	data, height int
	left, right  *AVLTreeNode
}

// AVLTree node root
type AVLTree struct {
	avl *AVLTreeNode
}

// max returns maximum value
func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// getHeight returns height of the tree
func (an *AVLTreeNode) getHeight() int {
	if an == nil {
		return 0
	}

	return an.height
}

// setHeight sets height of a node
func (an *AVLTreeNode) setHeight() int {
	if an == nil {
		return 0
	}

	lh := 0
	rh := 0

	if an.left != nil {
		lh = an.left.getHeight()
	}

	if an.right != nil {
		rh = an.right.getHeight()
	}

	return 1 + max(lh, rh)
}

// leftRotate construct tree after left roatation of a node
func (an *AVLTreeNode) leftRotate() *AVLTreeNode {
	newRoot := an.right
	an.right = newRoot.left
	newRoot.left = an
	an.height = an.setHeight()
	newRoot.height = newRoot.setHeight()

	return newRoot
}

// rightRotation construct tree after right rotation of a tree
func (an *AVLTreeNode) rightRotate() *AVLTreeNode {
	newRoot := an.left
	an.left = newRoot.right
	newRoot.right = an
	an.height = an.setHeight()
	newRoot.height = newRoot.setHeight()

	return newRoot
}

// Add node to a AVL tree
func (an *AVLTreeNode) Add(val int) *AVLTreeNode {
	if an == nil {
		return &AVLTreeNode{data: val, height: 1, left: nil, right: nil}
	}

	if an.data > val {
		an.left = an.left.Add(val)
	} else if an.data <= val {
		an.right = an.right.Add(val)
	}

	return an.RecalculateHeight()
}

// RecalculateHeight to recalculate height of each node of a tree
func (an *AVLTreeNode) RecalculateHeight() *AVLTreeNode {
	bf := an.left.getHeight() - an.right.getHeight()
	if bf > 1 { // left heavy
		if an.left.left.getHeight() >= an.left.right.getHeight() {
			an = an.rightRotate()
		} else {
			an.left = an.left.leftRotate()
			an = an.rightRotate()
		}
	} else if bf < -1 { // right heavy
		if an.right.right.getHeight() >= an.right.left.getHeight() {
			an = an.leftRotate()
		} else {
			an.right = an.right.rightRotate()
			an = an.leftRotate()
		}
	}

	an.height = an.setHeight()

	return an
}

// Remove node from a AVL tree
func (an *AVLTreeNode) Remove(val int) *AVLTreeNode {
	if an == nil {
		return nil
	}

	if an.data < val {
		an.right = an.right.Remove(val)
	} else if an.data > val {
		an.left = an.left.Remove(val)
	} else {
		if an.left != nil && an.right != nil {
			// replace value with smallest node of the right subtree
			rightMinVal := an.right.Smallest()
			an.data = rightMinVal
			an.right = an.right.Remove(rightMinVal)
		} else if an.left != nil {
			an = an.left
		} else if an.right != nil {
			an = an.right
		} else {
			an = nil
			return an
		}
	}

	return an.RecalculateHeight()
}

// Smallest returns smallest value of a tree
func (an *AVLTreeNode) Smallest() int {
	if an.left != nil {
		return an.left.Smallest()
	}

	return an.data
}

// InOrder prints node values after InOrder traversal
func (an *AVLTreeNode) InOrder() {
	if an == nil {
		return
	}

	an.left.InOrder()
	fmt.Println(an.data)
	an.right.InOrder()
}

// PreOrdere prints node values after PreOrder traversal
func (an *AVLTreeNode) PreOrder() {
	if an == nil {
		return
	}

	fmt.Println(an.data)
	an.left.PreOrder()
	an.right.PreOrder()
}

// Add node to a AVL tree
func (a *AVLTree) Add(val int) {
	a.avl = a.avl.Add(val)
}

// Remove node from a AVL tree
func (a *AVLTree) Remove(val int) {
	a.avl = a.avl.Remove(val)
}

// InOrderTraversal of AVL tree
func (a *AVLTree) InOrderTraversal() {
	a.avl.InOrder()
}

// PreOrderTraversal of AVL tree
func (a *AVLTree) PreOrderTraversal() {
	a.avl.PreOrder()
}

func main() {
	avl := &AVLTree{}
	avl.Add(10)
	avl.Add(20)
	avl.Add(30)
	avl.Add(40)
	avl.Add(50)
	avl.Add(25)

	fmt.Println("InOrder Traversal:")
	avl.InOrderTraversal()
	fmt.Println("PreOrder Traversal:")
	avl.PreOrderTraversal()

	avl.Remove(20)
	fmt.Println("PreOrder traversal after removing 20")
	avl.PreOrderTraversal()
}

/*
InOrder Traversal:
10
20
25
30
40
50
PreOrder Traversal:
30
20
10
25
40
50
PreOrder traversal after removing 20
30
25
10
40
50
*/
