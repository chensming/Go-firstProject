package main

import "fmt"

func main() {
	score := 65
	switch {
	case score > 90:
		fmt.Println("优秀")
	case score > 80:
		fmt.Println("良好")
	case score > 70, score > 60:
		fmt.Println("一般")
	default:
		fmt.Println("差")
	}
}

//一般