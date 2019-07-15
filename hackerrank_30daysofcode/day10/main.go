/*
Given a base-10 integer, n, convert it to binary (base-2). Then find and print the base-10 integer denoting the maximum number of consecutive 1's in n's binary representation.

Input format: n int
Constraints: 1 <= n <= 10^6

>5
1
>13
2
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// ConsecutiveOne returns the number of consecutive 1s in a string formatted binary number
func ConsecutiveOne(s string) int {
	num, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}
	binary := strconv.FormatInt(num, 2)
	count := 0
	var curCount int
	for _, v := range binary { // 110110111
		if string(v) == "1" {
			curCount++
			if curCount > count {
				count = curCount
			}
		} else {
			if string(v) == "0" {
				curCount = 0
			}
		}
	}
	return count
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	binary := scanner.Text()
	fmt.Println(ConsecutiveOne(binary))
}
