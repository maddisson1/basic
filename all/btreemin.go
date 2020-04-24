package piscine

func BTreeMin(root *TreeNode) *TreeNode {
	node := root
	min := root
	for node != nil {
		if node.Left != nil {
			min = node.Left
		}
		node = node.Left
	}
	return min
}
