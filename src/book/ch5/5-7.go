package main

import "fmt"

func main() {
	slice := []int{10, 20, 30, 40}
	newSlice := slice[1:3]
	fmt.Println("newSlice的元素为: ")
	for i,data := range newSlice {
		fmt.Printf("newSlice[%d] = % d\n", i, data)
	}
}
