package bst

import (
	"testing"
)

func TestInsert(t *testing.T) {
	tree := &BST{}

	// Test inserting a node to an empty tree
	err := tree.Insert(5)
	if err != nil || tree.root.data != 5 || tree.size != 1 {
		t.Errorf("Failed to insert node to an empty tree")
	}

	// Test inserting a node with a value less than the root
	err = tree.Insert(3)
	if err != nil || tree.root.left.data != 3 || tree.size != 2 {
		t.Errorf("Failed to insert node with a value less than the root")
	}

	// Test inserting a node with a value greater than the root
	err = tree.Insert(7)
	if err != nil || tree.root.right.data != 7 || tree.size != 3 {
		t.Errorf("Failed to insert node with a value greater than the root")
	}

	// Test inserting a node with a value equal to an existing node
	err = tree.Insert(7)
	if err != ErrDuplicatedNode {
		t.Errorf("Failed to return error when inserting a node with a value equal to an existing node")
	}
}

func TestRemove(t *testing.T) {
	tree := &BST{}

	// Test removing a node from an empty tree
	err := tree.Remove(5)
	if err != ErrNodeNotFound {
		t.Errorf("Failed to return error when removing a node from an empty tree")
	}

	// Add some nodes to the tree
	tree.Insert(5)
	tree.Insert(3)
	tree.Insert(7)

	// Test removing a leaf node
	err = tree.Remove(3)
	if err != nil || tree.root.left != nil || tree.size != 2 {
		t.Errorf("Failed to remove a leaf node")
	}

	// Test removing a node with one child
	err = tree.Remove(7)
	if err != nil || tree.root.right != nil || tree.size != 1 {
		t.Errorf("Failed to remove a node with one child")
	}

	// Test removing a node with two children
	tree.Insert(3)
	tree.Insert(7)
	err = tree.Remove(5)
	if err != nil || tree.root.data != 7 || tree.root.right != nil || tree.size != 2 {
		t.Errorf("Failed to remove a node with two children")
	}
}
