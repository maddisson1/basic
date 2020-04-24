package piscine

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	h := BTreeLevelCount(root)
	for i := 1; i < h+1; i++ {
		printLevel(root, i, f)
	}
}
