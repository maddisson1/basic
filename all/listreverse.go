package piscine

func ListReverse(l *List) {
	var prev *NodeL
	x := l.Head
	l.Tail = l.Head
	for x != nil {
		next := x.Next
		x.Next = prev
		prev = x
		x = next
	}
	l.Head = prev
}
