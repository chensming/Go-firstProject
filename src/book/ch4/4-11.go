package main

import "fmt"

func main() {
	str := "abcd"
	for index := range str{
		fmt.Printf("str[%d] = %c\n", index, str[index])
	}
	for index, _ := range str {
		fmt.Printf("str[%d] = %c\n", index, str[index])
	}
}
