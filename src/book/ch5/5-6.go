package main

import "fmt"

func main()  {
	slice := []int{10, 20, 30, 40, 50}
	slice[1] = 25
	fmt.Println("输出切片中的元素:")
	for i,data := range slice {
		fmt.Printf("slice[%d] = %d\n", i, data)
	}
}
