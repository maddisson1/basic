package piscine

func ConcatParams(args []string) string {

	str := ""
	flag := false

	for _, char := range args {

		if flag {

			str += "\n"
		}
		flag = true
		str += char
	}
	return str
}
