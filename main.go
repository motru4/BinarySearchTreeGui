package main

import (
	"wor/bst"
	"wor/guiApp"
)

func main() {

	tree := &bst.BST{}

	guiApp.StartGUI(tree)

}
