package main

import "fmt"

func main() {
	num := 3
	switch num {
	case 1:
		fmt.Println("按下的是1楼")
	case 2:
		fmt.Println("按下的是2楼")
	case 3:
		fmt.Println("按下的是3楼")
	default:
		fmt.Println("按下的是x楼")
	}

}
