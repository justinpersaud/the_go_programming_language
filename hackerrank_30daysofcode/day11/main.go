/*

Given a 6 x 6 2D Array, A:

1 1 1 0 0 0
0 1 0 0 0 0
1 1 1 0 0 0
0 0 0 0 0 0
0 0 0 0 0 0
0 0 0 0 0 0

We define an hourglass in A to be a subset of values with indices falling in this pattern in A's graphical representation:

a b c
  d
e f g

There are 16 hourglasses in A, and an hourglass sum is the sum of an hourglass' values.

Task:
Calculate the hourglass sum for every hourglass in A, then print the maximum hourglass sum.

Input Format:
There are 6 lines of input, where each line contains 6 space-separated integers desciribg 2D Array A; every value in A will be in the inclusive range of -9 to 9.

Sample Input:
1 1 1 0 0 0
0 1 0 0 0 0
1 1 1 0 0 0
0 0 2 4 4 0
0 0 0 2 0 0
0 0 1 2 4 0

Sample Output:
19

x = 0
y = 0
arr[x][y] + arr[x][y+1] + arr[x][y+2]
arr[x+1][y+1]
arr[x+2][y] + arr[x+2][y+1] + arr[x+2][y+2]
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// stringSliceIntSlice takes a slice of strings and returns a slice of integers
func StringSliceIntSlice(s []string) []int {
	var arr []int
	var val int
	var err error
	for i := 0; i < len(s); i++ {
		val, err = strconv.Atoi(s[i])
		if err != nil {
			panic(err)
		}
		arr = append(arr, val)
	}
	return arr
}

func HourglassSum(arr [][]int) int {
	var sum, largestSum int
	for x := 0; x <= 3; x++ {
		for y := 0; y <= 3; y++ {
			sum = 0
			sum += arr[x][y] + arr[x][y+1] + arr[x][y+2]
			sum += arr[x+1][y+1]
			sum += arr[x+2][y] + arr[x+2][y+1] + arr[x+2][y+2]
			if sum > largestSum {
				largestSum = sum
			}
		}
	}
	return largestSum
}

func StringsToSlice(s string) []string {
	var output []string
	output = append(strings.Split(s, " "))
	return output
}

func main() {
	// read 6 lines from stdin to create our data structure
	scanner := bufio.NewScanner(os.Stdin)
	arr := make([][]int, 6)
	var input []string
	for scanner.Scan() {
		input = StringsToSlice(scanner.Text())
		arr = append(arr, StringSliceIntSlice(input))
	}
	// calcuate sum of all hourglasses in the 2d array
	result := HourglassSum(arr)
	fmt.Println(result)
}
