package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("Hello world")
	var s = []string{"a", "b", "c", "d"}
	var i = []int{1, 2, 3, 4, 5}
	fmt.Println(test(s))
	fmt.Println(test(i))

	_, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("err:", err.Error())
	} else {
		fmt.Println("successful")
	}
}

func test[T any](s []T) ([]T, []T) {
	var mid = len(s) / 2
	return s[:mid], s[mid:]
}
