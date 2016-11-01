package main

import "fmt"

func main() {
	str := "abc"
	perm("", str)
}

func perm(prefix string, str string) {
	n := len(str)
	if n == 0 {
		fmt.Println(prefix)
	} else {
		for i := 0; i < n; i++ {
			perm(prefix+str[i:i+1], str[0:i]+str[(i+1):n])
		}
	}
}