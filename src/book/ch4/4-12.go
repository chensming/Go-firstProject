package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i > 3 {
			break
		}
		fmt.Printf("i = %d\n", i)
	}
}