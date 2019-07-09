/*Task
Given an integer, , perform the following conditional actions:

If  is odd, print Weird
If  is even and in the inclusive range of  to , print Not Weird
If  is even and in the inclusive range of  to , print Weird
If  is even and greater than , print Not Weird
Complete the stub code provided in your editor to print whether or not  is weird.
*/

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	NTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	N := int32(NTemp)
	isWeird(N)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func isWeird(n int32) {
	switch {
	case !(n%2 == 0):
		fmt.Println("Weird")
	case n%2 == 0 && (n >= 6 && n <= 20):
		fmt.Println("Weird")
	case n%2 == 0 && (n >= 2 && n <= 5):
		fmt.Println("Not Weird")
	case n%2 == 0 && n > 20:
		fmt.Println("Not Weird")
	default:
		fmt.Println("No match") // this should never happen
	}
}
