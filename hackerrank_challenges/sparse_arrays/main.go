/*
There is a collection of input strings and a collection of query strings. For each query string, determine how many times it occurs in the list of input strings.

strings = ['ab', 'ab', 'abc']
queries = ['ab', 'abc', 'bc']
results = [2, 1, 0]

Input format:

First line contains an integer n, size of the strings
Each of the next n lines contain a string strings[i]
The next line contains q, the size of the queries
Each of the next q lines contains a string q[i].
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	numLines, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}
	strings := make([]string, numLines)
	for i := 1; i <= numLines; i++ {
		scanner.Scan()
		strings = append(strings, scanner.Text())
	}
	scanner.Scan()
	numQueries, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}
	queries := make([]string, numQueries)
	for i := 1; i <= numQueries; i++ {
		scanner.Scan()
		queries = append(queries, scanner.Text())
	}

	results := make(map[string]int)

	for _, s := range strings { // ['ab', 'ab', 'abc']
		for _, q := range queries { // ['ab', 'abc', 'bc']
			if q == s {
				results[q]++
			}
		}
	}
	for _, value := range results {
		fmt.Println(value)
	}
}
