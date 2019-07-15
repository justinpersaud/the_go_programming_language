/*
Write a factorial function that takes a positive integer, N as a parameter and prints the result of N! ( N factorial).

*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func factorial(n int32) int32 {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}
	fmt.Println(factorial(int32(n)))
}
