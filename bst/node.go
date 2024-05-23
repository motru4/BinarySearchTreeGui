package bst

import "errors"

//------------------------------------------------------------------------------

var (
	ErrReplaceNodeNotSatisfy error = errors.New("bst: The new value of the node to be replaced must satisfy the binary search tree structure")
	ErrDuplicatedNode        error = errors.New("bst: Found duplicated value on tree")
	ErrNodeNotFound          error = errors.New("bst: Node not found")
)

//------------------------------------------------------------------------------

type Node struct {
	data  int
	left  *Node
	right *Node
}

//------------------------------------------------------------------------------

func (tree *BST) add(data int, node *Node) (*Node, error) {

	if node == nil {

		tree.size++

		return &Node{data: data}, nil
	}

	if node.data == data {

		return nil, ErrDuplicatedNode
	}

	if data < node.data {

		left, err := tree.add(data, node.left)

		if err != nil {
			return nil, err
		}

		node.left = left
	} else if data > node.data {

		right, err := tree.add(data, node.right)

		if err != nil {
			return nil, err
		}

		node.right = right
	}

	return node, nil
}

//------------------------------------------------------------------------------

func (tree *BST) delete(data int, node *Node) (*Node, error) {

	if node == nil {

		return nil, ErrNodeNotFound
	}

	if data < node.data {

		left, err := tree.delete(data, node.left)

		if err != nil {
			return nil, err
		}

		node.left = left
	} else if data > node.data {

		right, err := tree.delete(data, node.right)

		if err != nil {
			return nil, err
		}

		node.right = right
	} else {

		if node.left != nil && node.right != nil {

			min := tree.min(node.right)
			node.data = min.data
			right, err := tree.delete(min.data, node.right)

			if err != nil {
				return nil, err
			}

			node.right = right
		} else {

			if node.left == nil && node.right == nil {

				node = nil

			} else if node.left == nil {

				node = node.right

			} else {

				node = node.left
			}

			tree.size--
		}
	}

	return node, nil
}

//------------------------------------------------------------------------------

func (tree *BST) change(data int, newdata int) (*Node, error) {
	node := tree.root

	if node == nil {
		return nil, ErrNodeNotFound
	}

	var inorder []int = tree.inorder(node)
	var index int = indexOf(inorder, data)
	var nodeToChange *Node = tree.find(data, node)

	if index == -1 {
		return nil, ErrNodeNotFound
	}

	if index == 0 {
		if inorder[index+1] > newdata {
			nodeToChange.data = newdata
		} else {
			return nil, ErrReplaceNodeNotSatisfy
		}
	} else if index == len(inorder)-1 {
		if inorder[index-1] < newdata {
			nodeToChange.data = newdata
		} else {
			return nil, ErrReplaceNodeNotSatisfy
		}
	} else {
		if inorder[index+1] > newdata && inorder[index-1] < newdata {
			nodeToChange.data = newdata
		} else {
			return nil, ErrReplaceNodeNotSatisfy
		}
	}

	return node, nil
}
