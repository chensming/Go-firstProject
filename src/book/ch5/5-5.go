package main

import "fmt"

func main()  {
	mySlice1 := make([]int, 5)
	fmt.Println("Element of slice is: ")
	for _, data := range mySlice1 {
		fmt.Printf("%d ", data)
	}
	fmt.Println()
}
