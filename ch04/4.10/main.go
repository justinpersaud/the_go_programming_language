/*
Build a tool that lets users create, read, update, and delete GitHub issues from the command line, invoking their preferred text editor when substantial text input is required.
*/

package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"the_go_programming_language/ch04/github"

	"github.com/urfave/cli"
)

func main() {
	// hardcode the repository for now
	repository := github.Repository{
		Owner: "justinpersaud",
		Name:  "the_go_programming_language",
	}

	issue := github.Issue{
		Title: "My test issue",
	}

	app := cli.NewApp()
	app.Name = "GitHub Issue Editor"
	app.Usage = "Lets you create, read, update and delete GitHub issues."

	searchFlags := []cli.Flag{
		cli.StringSliceFlag{
			Name:  "query",
			Value: &cli.StringSlice{""},
			Usage: "Searches GitHub issues",
		},
	}
	issueFlags := []cli.Flag{
		cli.StringFlag{
			Name:  "number",
			Value: "",
			Usage: "Issue to fetch",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "create",
			Usage: "Creates GitHub issues",
			Action: func(c *cli.Context) error {
				result, err := github.CreateIssue(issue, repository)
				if err != nil {
					return err
				}
				fmt.Println(result)
				return nil
			},
		},
		{
			Name:  "search",
			Usage: "Searches GitHub issues",
			Flags: searchFlags,
			Action: func(c *cli.Context) error {
				result, err := github.SearchIssues(c.StringSlice("query"))
				if err != nil {
					return err
				}
				// fmt.Println(result)
				fmt.Printf("%d issues matching query:\n", result.TotalCount)
				for _, item := range result.Items {
					fmt.Printf("#%-5d %9.9s %v %.55s\n",
						item.Number, item.User.Login, item.CreatedAt, item.Title)
				}
				return nil
			},
		},
		{
			Name:  "get",
			Usage: "Gets information about a GitHub issue",
			Flags: issueFlags,
			Action: func(c *cli.Context) error {
				issueNum, err := strconv.Atoi(c.String("number"))
				if err != nil {
					panic(err)
				}
				issue := github.Issue{Number: issueNum}
				result, err := github.ReadIssue(issue, repository)
				if err != nil {
					return err
				}
				// fmt.Println(result)
				fmt.Printf("#%-5d %9.9s %v %.55s\n",
					result.Number, result.User.Login, result.CreatedAt, result.Title)
				return nil
			},
		},
		// {
		// 	Name:  "edit",
		// 	Usage: "Edit a GitHub issue",
		// 	Flags: issueFlags,
		// 	Action: func(c *cli.Context) error {
		// 		issueNum, err := strconv.Atoi(c.String("number"))
		// 		if err != nil {
		// 			panic(err)
		// 		}
		// 		issue := github.Issue{Number: issueNum}
		// 		result, err := github.ReadIssue(issue, repository)
		// 		if err != nil {
		// 			return err
		// 		}
		// 		// fmt.Println(result)
		// 		fmt.Printf("#%-5d %9.9s %v %.55s\n",
		// 			result.Number, result.User.Login, result.CreatedAt, result.Title)
		// 		return nil
		// 	},
		// },
	}

	// start the app
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
