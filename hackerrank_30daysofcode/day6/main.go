package main

import (
	"bufio"
	"fmt"
	"os"
)

func oddEvenSplit(data []string) {
	var odd, even string
	for _, str := range data {
		odd = ""
		even = ""
		for i, w := range str {
			if i%2 == 0 {
				even += string(w)
			} else {
				odd += string(w)
			}
		}
		fmt.Printf("%v %v\n", even, odd)
	}
}

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	data := make([]string, 0)
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	oddEvenSplit(data[1:])
}
