package main

import "fmt"

func main()  {
	var a[2] int
	a[0] = 10
	a[1] = 12
	for i := 0; i < len(a); i++ {
		fmt.Printf("a[%d]=%d\n", i, a[i])
	}
}
