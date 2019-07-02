// Dup2 prints the count and text of lines that appear more than once
// in the input. It reads from stdin or from a list of named files.package main

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		fmt.Fprintln(os.Stdout, "No args passed, reading from std in")
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			defer f.Close()
			countLines(f, counts)
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%s:\t%d\n", line, n)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
