package main

import (
	"fmt"
	"strings"
)

func createPointer() *int {
	value := 42
	return &value
}

func main() {
	ptr := createPointer()
	fmt.Println(*ptr) // 输出：42
	fmt.Println(strings.Join([]string{"123123", "sadads", "qweqe"}, "()"))
}
