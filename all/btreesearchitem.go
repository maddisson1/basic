package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Data == elem {
		return root
	}
	l := BTreeSearchItem(root.Left, elem)
	if l != nil {
		return l
	}
	r := BTreeSearchItem(root.Right, elem)
	return r
}
