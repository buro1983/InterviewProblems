/*
Binary tree operations
Insert
Delete
Search
*/

package main

import (
	"fmt"
)

// BinaryTreeNode contains BST node data and next children address
type BinaryTreeNode struct {
	data        int
	left, right *BinaryTreeNode
}

// BiinaryTree construct BST
type BinaryTree struct {
	btree *BinaryTreeNode
}

// Add and node in BST
func (btn *BinaryTreeNode) Add(val int) *BinaryTreeNode {
	if btn == nil {
		return &BinaryTreeNode{data: val, left: nil, right: nil}
	}

	if btn.data >= val {
		btn.left = btn.left.Add(val)
	} else {
		btn.right = btn.right.Add(val)
	}

	return btn
}

// InOrder traversal of BST
func (bt *BinaryTree) InOrder() {
	if bt.btree == nil {
		return
	}

	bt.btree.InOrder()
}

// InOrder traversal of BST
func (btn *BinaryTreeNode) InOrder() {
	if btn == nil {
		return
	}

	btn.left.InOrder()
	fmt.Println(btn.data)
	btn.right.InOrder()
}

// Add in BST
func (bt *BinaryTree) Add(val int) {
	bt.btree = bt.btree.Add(val)
}

// Delete node from BST
func (bt *BinaryTree) Delete(val int) {
	if bt.btree == nil {
		fmt.Println("BST is empty")
		return
	}

	bt.btree = bt.btree.Del(val)
}

// NextInOrderSucc returns next inorder successor of a node
func (btn BinaryTreeNode) NextInOrderSucc() int {
	if btn.left == nil {
		return btn.data
	}

	return btn.left.NextInOrderSucc()
}

// Del a node on BST
func (btn *BinaryTreeNode) Del(val int) *BinaryTreeNode {
	if btn == nil {
		fmt.Println("Value", val, "to be deleted not found")
		return btn
	}

	switch {
	case btn.data < val:
		btn.right = btn.right.Del(val)

	case btn.data > val:
		btn.left = btn.left.Del(val)

	default:
		if btn.left == nil && btn.right == nil { // node chindren
			return nil
		} else if btn.left == nil { // no left children
			return btn.right
		} else if btn.right == nil { // no right children
			return btn.left
		} else { // both node has children
			btn.data = btn.right.NextInOrderSucc()
			btn.right = btn.right.Del(btn.data)
		}
	}

	return btn
}

// Search value in BST
func (bt BinaryTree) Search(val int) bool {
	if bt.btree == nil {
		return false
	}

	return bt.btree.Search(val)
}

func (btn *BinaryTreeNode) Search(val int) bool {
	if btn == nil {
		return false
	}

	if btn.data < val {
		return btn.right.Search(val)
	} else if btn.data > val {
		return btn.left.Search(val)
	}

	return true
}

func main() {

	btree := &BinaryTree{}
	btree.Add(50)
	btree.Add(30)
	btree.Add(20)
	btree.Add(40)
	btree.Add(70)
	btree.Add(60)

	fmt.Println("Print InOrder after adding data")
	btree.InOrder()

	fmt.Println("Add 12 and print tree")
	btree.Add(80)
	btree.InOrder()

	btree.Delete(20)
	fmt.Println("BST after deletion 20")
	btree.InOrder()

	fmt.Println("Search for 70:", btree.Search(70))
	fmt.Println("Search for -70:", btree.Search(-70))
	fmt.Println("Search for 1000:", btree.Search(1000))

	btree.Delete(30)
	fmt.Println("BST after deletion 30")
	btree.InOrder()

	btree.Delete(70)
	fmt.Println("BST after deletion 70")
	btree.InOrder()
}

/*
Print InOrder after adding data
20
30
40
50
60
70
Add 12 and print tree
20
30
40
50
60
70
80
BST after deletion 20
30
40
50
60
70
80
Search for 70: true
Search for -70: false
Search for 1000: false
BST after deletion 30
40
50
60
70
80
BST after deletion 70
40
50
60
80
*/
