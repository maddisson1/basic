package piscine

func ListMerge(l1 *List, l2 *List) {
	if l1.Head == nil {
		l1.Head = l2.Head
		l1.Tail = l2.Tail
		return
	} else if l2.Head == nil {
		return
	}

	curr := l1.Head
	for curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = l2.Head
	l1.Tail = l2.Tail
}
