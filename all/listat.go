package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	cnt := 1
	for l != nil {

		if cnt == pos+1 {
			return l
		}
		cnt++
		l = l.Next

	}
	return nil
}
