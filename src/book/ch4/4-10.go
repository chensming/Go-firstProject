package main

import "fmt"

func main() {
	str := "abcd"
	for index, value := range str {
		fmt.Printf("str[%d] = %c\n", index, value)
	}
}
