package piscine

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	if n1 == nil && n2 != nil {
		return n2
	} else if n1 != nil && n2 == nil {
		return n1
	} else if n1 == nil && n2 == nil {
		return nil
	}
	curr1 := n1
	curr2 := n2

	var new *NodeI = nil
	var head *NodeI = nil

	for curr1 != nil && curr2 != nil {
		if curr1.Data > curr2.Data {
			if new == nil {
				new = curr2
				head = new
			} else {
				new.Next = curr2
				new = new.Next
			}
			curr2 = curr2.Next

		} else {
			if new == nil {
				new = curr1
				head = new
			} else {
				new.Next = curr1
				new = new.Next
			}
			curr1 = curr1.Next
		}
	}

	if curr1 != nil {
		for curr1 != nil {
			new.Next = curr1
			new = new.Next
			curr1 = curr1.Next
		}
	} else if curr2 != nil {
		for curr2 != nil {
			new.Next = curr2
			new = new.Next
			curr2 = curr2.Next
		}
	}
	return head
}
