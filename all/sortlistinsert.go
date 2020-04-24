package piscine

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	if l == nil {
		return nil
	}
	curr := l
	head := l
	n := &NodeI{Data: data_ref}
	var prev *NodeI = nil
	for curr.Next != nil {
		if curr.Data > data_ref {
			if prev == nil {
				head = n
				n.Next = curr
			} else {
				prev.Next = n
				n.Next = curr
			}
			return head
		}
		prev = curr
		curr = curr.Next
	}
	if curr.Data > data_ref {
		prev.Next = n
		n.Next = curr
	} else {
		curr.Next = n
	}
	return head
}
