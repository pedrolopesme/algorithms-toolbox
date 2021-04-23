package main

type Node struct {
	val   int
	left  *Node
	right *Node
}

type BinarySearchTree struct {
	root *Node
}

func (bst *BinarySearchTree) Insert(val int) {
	newNode := Node{val: val}

	if bst.root == nil {
		bst.root = newNode
	} else {
		iNode := bst.root
		for iNode != nil {
			if val < iNode.val {
				if iNode.left == nil {
					iNode.left = newNode
					break
				}
				iNode = iNode.left
			} else {
				if iNode.right == nil {
					iNode.right = newNode
					break
				}
				iNode = iNode.right
			}
		}
	}
}

func (bst *BinarySearchTree) Lookup(int val) *Node {
	iNode = bst.root
	for iNode != null {
		if iNode.val == val {
			break
		}

		if val < iNode.val {
			iNode = iNode.left
		} else {
			iNode = iNode.right
		}
	}
	return iNode
}

func main() {
	// silence is gold.
}
