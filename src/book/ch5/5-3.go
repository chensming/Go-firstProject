package main

import "fmt"

func modify(a[4] int){
	a[3] = 3
	fmt.Println(a)
}

func main()  {
	a := [4]int{1,2,3,4}
	modify(a)
	fmt.Println(a)
}