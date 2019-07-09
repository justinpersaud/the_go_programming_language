package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var _ = strconv.Itoa // Ignore this comment. You can still use the package "strconv".

	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank "

	scanner := bufio.NewScanner(os.Stdin)
	// Declare second integer, double, and String variables.
	var var1 uint64
	var var2 float64
	var var3 string

	// Read and save an integer, double, and String to your variables.

	scanner.Scan()
	var1, _ = strconv.ParseUint(scanner.Text(), 10, 64)
	scanner.Scan()
	var2, _ = strconv.ParseFloat(scanner.Text(), 64)
	scanner.Scan()
	var3 = scanner.Text()

	// Print the sum of both integer variables on a new line.
	fmt.Println(var1 + i)
	// Print the sum of the double variables on a new line.
	fmt.Printf("%.1f\n", d+var2)
	// Concatenate and print the String variables on a new line
	// The 's' variable above should be printed first.
	fmt.Println(s + var3)
}
