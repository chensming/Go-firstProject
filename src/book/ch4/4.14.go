package main

import "fmt"

func main() {
	fmt.Println("111")
	goto next
	fmt.Println("222")
next:
	fmt.Print("333")
}
