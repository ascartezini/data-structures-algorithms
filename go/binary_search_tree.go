package main

type BST struct {
	left, right *BST
	val         int
}

func isPresent(bst *BST, num int) bool {
	// return isNodePresentRecursive(bst, num)
	return isNodePresentIterative(bst, num)
}

// recursive
// Average: O(log(n)) time | O(log(n)) space
// Worst: O(n) time | O(n) space
func isNodePresentRecursive(bst *BST, num int) bool {
	present := false

	if bst != nil {
		if num > bst.val {
			isNodePresentRecursive(bst.right, num)
		} else if num < bst.val {
			isNodePresentRecursive(bst.left, num)
		} else if num == bst.val {
			present = true
		}
	}

	return present
}

// iterative
// Average: O(log(n)) time | O(1) space
// Worst: O(n) time | O(1) space
func isNodePresentIterative(bst *BST, num int) bool {
	present := false
	rootNode := bst

	for rootNode != nil {
		if num > rootNode.val {
			rootNode = rootNode.right
		} else if num < rootNode.val {
			rootNode = rootNode.left
		} else if num == rootNode.val {
			present = true
			break
		}
	}

	return present
}
