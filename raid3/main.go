package main

import (
	"bufio"
	"fmt"
	"os"
)

func printArr(arr [][]string) {
	for i, s := range arr {
		fmt.Print(s, " ")
		if i%3 == 2 && i < len(arr)-1 {
			fmt.Print(" || ")
		}
	}
	fmt.Println()
}

func main() {
	file, _ := os.Stdin.Stat()
	if file.Mode()&os.ModeNamedPipe > 0 {
		reader := bufio.NewReader(os.Stdin)
		arr, _ := reader.ReadBytes('\x10')
		if len(arr) > 1 {
			x, y := 0, 0
			for _, s := range arr {
				if s == 10 {
					y++
				}
			}
			for i := 0; arr[i] != 10; i++ {
				x++
			}
			newarr := [][]string{}
			if arr[0] == 'o' {
				newarr = append(newarr, []string{"raid1a"}, []string{itoa(x)}, []string{itoa(y)})
			}
			if arr[0] == '/' {
				newarr = append(newarr, []string{"raid1b"}, []string{itoa(x)}, []string{itoa(y)})
			}
			if x == 1 && y == 1 && arr[0] == 'A' {
				newarr = append(newarr, []string{"raid1c"}, []string{itoa(x)}, []string{itoa(y)})
				newarr = append(newarr, []string{"raid1d"}, []string{itoa(x)}, []string{itoa(y)})
				newarr = append(newarr, []string{"raid1e"}, []string{itoa(x)}, []string{itoa(y)})
			} else if x == 1 && arr[0] == 'A' && arr[2*(y-x)] == 'A' {
				newarr = append(newarr, []string{"raid1c"}, []string{itoa(x)}, []string{itoa(y)})
			} else if x == 1 && arr[0] == 'A' && arr[2*(y-x)] == 'C' {
				newarr = append(newarr, []string{"raid1c"}, []string{itoa(x)}, []string{itoa(y)})
				newarr = append(newarr, []string{"raid1e"}, []string{itoa(x)}, []string{itoa(y)})
			} else if y == 1 && arr[0] == 'A' && arr[x-1] == 'C' {
				newarr = append(newarr, []string{"raid1d"}, []string{itoa(x)}, []string{itoa(y)})
				newarr = append(newarr, []string{"raid1e"}, []string{itoa(x)}, []string{itoa(y)})
			} else if y == 1 && arr[0] == 'A' && arr[x-1] == 'A' {
				newarr = append(newarr, []string{"raid1c"}, []string{itoa(x)}, []string{itoa(y)})
			} else if arr[0] == 'A' && arr[x-1] == 'A' && arr[x*y-1] == 'C' && arr[x*y+x-2] == 'C' {
				newarr = append(newarr, []string{"raid1c"}, []string{itoa(x)}, []string{itoa(y)})
			} else if arr[0] == 'A' && arr[x-1] == 'C' && arr[x*y-1] == 'A' && arr[x*y+x-2] == 'C' {
				newarr = append(newarr, []string{"raid1d"}, []string{itoa(x)}, []string{itoa(y)})
			} else if arr[0] == 'A' && arr[x-1] == 'C' && arr[x*y-1] == 'C' && arr[x*y+x-2] == 'A' {
				newarr = append(newarr, []string{"raid1c"}, []string{itoa(x)}, []string{itoa(y)})
			}
			printArr(newarr)
		} else {
			fmt.Println()
		}
	} else {
		fmt.Println("No pipe data")
	}
}

func powersize(n int) int {
	powersize := 0
	for n != 0 {
		n /= 10
		powersize++
	}
	if powersize != 0 {
		powersize--
	}
	return powersize
}

func power(x int, n int) int {
	answer := 1
	for i := 0; i < n; i++ {
		answer *= x
	}
	return answer
}

func itoa(n int) string {
	str := ""
	if n < 0 {
		str += "-"
		if n == -9223372036854775808 {
			str += "9"
			n %= 1000000000000000000
		}
		n *= -1
	}
	powsize := powersize(n)
	for powsize > -1 {
		p := power(10, powsize)
		str += string(rune((n / p) + 48))
		n %= p
		powsize--
	}
	return str
}
