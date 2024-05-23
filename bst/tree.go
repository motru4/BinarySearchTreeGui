package bst

import "strconv"

//------------------------------------------------------------------------------

type BST struct {
	size int
	root *Node
}

//------------------------------------------------------------------------------

func (tree *BST) min(node *Node) *Node {

	for node.left != nil {

		node = node.left
	}

	return node
}

//------------------------------------------------------------------------------

func indexOf(slice []int, value int) int {
	for i, v := range slice {
		if v == value {
			return i
		}
	}
	return -1
}

//------------------------------------------------------------------------------

func (tree *BST) find(data int, node *Node) *Node {
	if node == nil {
		return nil
	}

	if data < node.data {
		return tree.find(data, node.left)
	} else if data > node.data {
		return tree.find(data, node.right)
	}

	return node
}

//------------------------------------------------------------------------------

func (tree BST) toTreeString(prefix string, top bool, str string, node *Node) string {

	left := new(Node)
	right := new(Node)

	if node == nil {

		return ""
	}

	left = node.left
	right = node.right

	if right != nil {

		temp := tree.path(top, ""+prefix, "│   ", "    ")
		str = tree.toTreeString(temp, false, str, right)
	}

	str = tree.path(top, str+prefix, "└──", "┌──")
	str = str + " " + strconv.Itoa(node.data) + "\n"

	if left != nil {

		temp := tree.path(top, ""+prefix, "    ", "│   ")
		str = tree.toTreeString(temp, true, str, left)
	}

	return str
}

//------------------------------------------------------------------------------

func (tree BST) path(condition bool, str string, choice1 string, choice2 string) string {

	if condition {

		str += choice1

	} else {

		str += choice2
	}

	return str
}

//------------------------------------------------------------------------------

func (tree BST) postorderStr(node *Node, str string) string {

	if node != nil {
		str = tree.postorderStr(node.left, str)
		str = tree.postorderStr(node.right, str)
		str = str + strconv.Itoa(node.data) + " "
	}

	return str
}

func (tree *BST) postorder(node *Node) []int {
	if node == nil {
		return []int{}
	}

	left := tree.postorder(node.left)
	right := tree.postorder(node.right)

	return append(append(left, right...), node.data)
}

//------------------------------------------------------------------------------

func (tree BST) preorderStr(node *Node, str string) string {

	if node != nil {
		str = str + strconv.Itoa(node.data) + " "
		str = tree.preorderStr(node.left, str)
		str = tree.preorderStr(node.right, str)
	}

	return str
}

func (tree *BST) preorder(node *Node) []int {
	if node == nil {
		return []int{}
	}

	left := tree.preorder(node.left)
	right := tree.preorder(node.right)

	return append(append([]int{node.data}, left...), right...)
}

//------------------------------------------------------------------------------

func (tree BST) inorderStr(node *Node, str string) string {

	if node != nil {
		str = tree.inorderStr(node.left, str)
		str = str + strconv.Itoa(node.data) + " "
		str = tree.inorderStr(node.right, str)
	}

	return str
}

func (tree *BST) inorder(node *Node) []int {
	if node == nil {
		return []int{}
	}

	left := tree.inorder(node.left)
	right := tree.inorder(node.right)

	return append(append(left, node.data), right...)
}
