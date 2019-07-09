package main

import (
	"fmt"
)

func main() {
	abc := 123
	abcPointer := &abc
	*abcPointer = 456
	fmt.Println(abc) // should be 456
}
