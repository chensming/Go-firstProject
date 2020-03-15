package main

import "fmt"

func main(){
	slice := []int{10,20,30,40,50}
	newSlice := slice[1:3]
	newSlice[1] = 4
	fmt.Println("The element of the slice: ")
	for i,data := range slice{
		fmt.Printf("slice[%d] = %d\n", i, data)
	}
	fmt.Println("The element of the newSlice: ")
	for v, data2 := range newSlice {
		fmt.Printf("newSlice[%d] = %d\n", v, data2)
	}
}