package piscine

func BTreeMax(root *TreeNode) *TreeNode {
	node := root
	max := root
	for node != nil {
		if node.Right != nil {
			max = node.Right
		}
		node = node.Right
	}
	return max
}
