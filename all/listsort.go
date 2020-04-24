package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	if l == nil {
		return nil
	}
	head := l
	not_sorted := true
	for not_sorted {
		curr := head
		var prev *NodeI = nil
		not_sorted = false
		for curr.Next != nil {
			if curr.Data > curr.Next.Data {
				not_sorted = true
				next := curr.Next //Remember next node
				if prev == nil {
					head = curr.Next
				} else {
					prev.Next = next //Set previous node to next
				}
				if curr.Next.Next == nil {
					curr.Next = nil
				} else {
					curr.Next = curr.Next.Next
				}
				next.Next = curr
				prev = next
			} else {
				prev = curr
				curr = curr.Next
			}
		}
	}
	return head
}
