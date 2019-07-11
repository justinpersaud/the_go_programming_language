/*
Given a string, S, of length N that is indexed from 0 to N-1, print its even-indexed and odd-indexed characters as 2 space-separated strings on a single line (see the Sample below for more detail).

Note: 0 is considered to be an even index.

Sample Input

2
Hacker
Rank
Sample Output

Hce akr
Rn ak


slice = ["Hacker", "Rank"]
*/

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
