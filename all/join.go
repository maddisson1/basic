package piscine

func Join(strs []string, sep string) string {

	res := ""
	nf := false

	for _, s := range strs {

		if nf {
			res += sep
		}
		nf = true
		res += s
	}
	return res
}
