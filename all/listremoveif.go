package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	if l.Head == nil {
		return
	}

	var prev *NodeL = nil
	curr := l.Head

	for curr != nil {
		if curr.Data != data_ref {
			prev = curr
			curr = curr.Next
			continue
		}
		curr = curr.Next
		if prev == nil {
			l.Head = curr
		} else {
			prev.Next = curr
		}
	}
}
