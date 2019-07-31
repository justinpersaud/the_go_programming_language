// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"fmt"
	"log"
	"os"
	"the_go_programming_language/ch04/github"
	"time"
)

//!+

const aMonth = 31 * 24
const aYear = 31 * 365 * 24

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours())
}

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)

	var inMonth, inYear, outYear []*github.Issue
	for _, item := range result.Items {
		// less a month/a year, or more than a year old
		if daysAgo(item.CreatedAt) < aYear {
			inMonth = append(inMonth, item)
		} else if daysAgo(item.CreatedAt) < aMonth {
			inYear = append(inYear, item)
		} else {
			outYear = append(outYear, item)
		}
	}

	fmt.Printf("%d issues less than a month:\n", len(inMonth))
	for _, item := range inMonth {
		fmt.Printf("#%-5d %9.9s %v %.55s\n",
			item.Number, item.User.Login, item.CreatedAt, item.Title)
	}

	fmt.Printf("\n%d issues less than a year:\n", len(inYear))
	for _, item := range inYear {
		fmt.Printf("#%-5d %9.9s %v %.55s\n",
			item.Number, item.User.Login, item.CreatedAt, item.Title)
	}

	fmt.Printf("\n%d issues more than a month:\n", len(outYear))
	for _, item := range outYear {
		fmt.Printf("#%-5d %9.9s %v %.55s\n",
			item.Number, item.User.Login, item.CreatedAt, item.Title)
	}
}
