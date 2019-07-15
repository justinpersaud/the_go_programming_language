/*
Given n names and phone numbers, assemble a phone book that maps friends' names to their respective phone numbers. You will then be given an unknown number of names to query your phone book for. For each name queried, print the associated entry from your phone book on a new line in the form name=phoneNumber; if an entry for name is not found, print Not found instead.

Note: Your phone book should be a Dictionary/Map/HashMap data structure.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// splits a string separated by a space on a single line into two return values
func splitString(s string) []string {
	return strings.Split(s, " ")

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// the first input should be the number of keys and values we are adding
	scanner.Scan()
	entries, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}
	phonebook := make(map[string]string)
	for k := 1; k <= entries; k++ {
		scanner.Scan()
		slice := splitString(scanner.Text())
		// assume we're only providing new entries
		phonebook[slice[0]] = slice[1]
	}

	// read an unknown number of lines which are checking if the keys are in the phonebook
	for scanner.Scan() {
		_, ok := phonebook[scanner.Text()]
		if ok {
			fmt.Printf("%s=%s\n", scanner.Text(), phonebook[scanner.Text()])
		} else {
			fmt.Println("Not found")
		}
	}
}
