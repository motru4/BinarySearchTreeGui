package bst

//------------------------------------------------------------------------------

func (tree *BST) Size(size ...int) int {
	if len(size) > 0 {
		tree.size = size[0]
	}
	return tree.size
}

//------------------------------------------------------------------------------

func (tree *BST) Root(root ...*Node) *Node {
	if len(root) > 0 {
		tree.root = root[0]
	}
	return tree.root
}

//------------------------------------------------------------------------------

func (tree *BST) Insert(data int) error {

	root, err := tree.add(data, tree.root)

	if err != nil {
		return err
	}

	tree.root = root
	return nil
}

//------------------------------------------------------------------------------

func (tree *BST) Remove(data int) error {

	root, err := tree.delete(data, tree.root)

	if err != nil {
		return err
	}

	tree.root = root
	return nil
}

//------------------------------------------------------------------------------

func (tree *BST) Replace(data int, newdata int) error {

	root, err := tree.change(data, newdata)

	if err != nil {
		return err
	}

	tree.root = root
	return nil
}

//------------------------------------------------------------------------------

func (tree BST) Inorder() []int {

	inorder := tree.inorder(tree.root)

	return inorder
}

func (tree BST) InorderStr() string {

	postorder := tree.inorderStr(tree.root, "")

	return postorder
}

//------------------------------------------------------------------------------

func (tree BST) Preorder() []int {

	preorder := tree.preorder(tree.root)

	return preorder
}

func (tree BST) PreorderStr() string {

	postorder := tree.preorderStr(tree.root, "")

	return postorder
}

//------------------------------------------------------------------------------

func (tree BST) Postorder() []int {

	postorder := tree.postorder(tree.root)

	return postorder
}

func (tree BST) PostorderStr() string {

	postorder := tree.postorderStr(tree.root, "")

	return postorder
}

//------------------------------------------------------------------------------

func (tree BST) TreeString() string {

	str := tree.toTreeString("", true, "", tree.root)
	return str
}
