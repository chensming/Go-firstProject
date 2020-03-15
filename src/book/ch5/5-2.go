package main

import "fmt"

func main() {
	var a [2] int
	a[0] = 10
	a[1] = 12
	for i, v := range a {
		fmt.Println(i, ": ", v)
	}
}
