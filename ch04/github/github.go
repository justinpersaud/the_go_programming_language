package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const (
	// BaseURL is the v3 endpoint for GitHub.
	BaseURL = "https://api.github.com/"
)

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type IssueResult struct {
	URL   string `json:"url"`
	Title string `json:"title"`
	Body  string `json:"body"`
	State string `json:"state"`
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string `json:"title"`
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

type Repository struct {
	Owner string
	Name  string
}

// SearchIssues queries the GitHub issue tracker.
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(BaseURL + "search/issues" + "?q=" + q)
	if err != nil {
		return nil, err
	}

	// We must close resp.Body on all execution paths.
	// (Chapter 5 presents 'defer', which makes this simpler.)
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

// CreateIssue creates an issue on GitHub
func CreateIssue(issue Issue, repo Repository) (*IssueResult, error) {
	endpoint := fmt.Sprintf("%s%s%s%s%s%s", BaseURL, "repos/", repo.Owner, "/", repo.Name, "/issues")
	data, err := json.Marshal(issue)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "token "+os.Getenv("GITHUB_TOKEN"))

	resp, err := client.Do(req)

	if resp.StatusCode != http.StatusCreated {
		resp.Body.Close()
		return nil, fmt.Errorf("issue create failed: %s", resp.Status)
	}

	var result IssueResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

// ReadIssue fetches info from a GitHub issue.
func ReadIssue(issue Issue, repo Repository) (*Issue, error) {
	endpoint := fmt.Sprintf("%s%s%s%s%s%s%s%d", BaseURL, "repos/", repo.Owner, "/", repo.Name, "/issues", "/", issue.Number)
	resp, err := http.Get(endpoint)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("error fetching issue: %s, %s", endpoint, resp.Status)
	}
	var result Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

// // EditIssue updates a GitHub issue.
// func EditIssue(issue Issue, repo Repository) (*IssueResult, error) {
// 	endpoint := fmt.Sprintf("%s%s%s%s%s%s%s%d", BaseURL, "repos/", repo.Owner, "/", repo.Name, "/issues", "/", issue.Number)
// 	data, err := json.Marshal(issue)
// 	if err != nil {
// 		return nil, err
// 	}
// 	client := &http.Client{}
// 	req, err := http.NewRequest("PATCH", endpoint, bytes.NewBuffer(data))
// 	if err != nil {
// 		return nil, err
// 	}
// 	req.Header.Add("Content-Type", "application/json")
// 	req.Header.Add("Authorization", "token "+os.Getenv("GITHUB_TOKEN"))

// 	resp, err := client.Do(req)

// 	if resp.StatusCode != http.StatusOK {
// 		resp.Body.Close()
// 		return nil, fmt.Errorf("issue edit failed: %s", resp.Status)
// 	}
// 	var result IssueResult
// 	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
// 		resp.Body.Close()
// 		return nil, err
// 	}
// 	resp.Body.Close()
// 	return &result, nil
// }
