package main

type BinaryTree struct {
	Value int

	Left  *BinaryTree
	Right *BinaryTree
}

func (tree *BinaryTree) InvertBinaryTree() {
	tree.Left, tree.Right = tree.Right, tree.Left

	if tree.Left != nil {
		tree.Left.InvertBinaryTree()
	}

	if tree.Right != nil {
		tree.Right.InvertBinaryTree()
	}
}

func main() {
	tree := &BinaryTree{
		Value: 20,
		Left: &BinaryTree{Value: 10,
			Left: &BinaryTree{Value: 8,
				Left: &BinaryTree{Value: 6}},
			Right: &BinaryTree{Value: 12,
				Left:  &BinaryTree{Value: 11},
				Right: &BinaryTree{Value: 13}}},
		Right: &BinaryTree{Value: 30,
			Left: &BinaryTree{Value: 25,
				Left: &BinaryTree{Value: 23}},
			Right: &BinaryTree{Value: 40}}}

	tree.InvertBinaryTree()
}
