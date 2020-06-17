package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

func main() {
	result, err := searchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	var oneMonth []*Issue
	var oneYear []*Issue
	var others []*Issue
	for _, item := range result.Items {
		days := time.Since(item.CreatedAt).Hours() / 24
		if days < 30 {
			oneMonth = append(oneMonth, item)
		} else if days < 365 {
			oneYear = append(oneYear, item)
		} else {
			others = append(others, item)
		}
	}

	fmt.Printf("%d issues:\n", result.TotalCount)
	fmt.Println("Less than a month")
	printIssues(oneMonth)
	fmt.Println("Less than a year")
	printIssues(oneYear)
	fmt.Println("More than a year")
	printIssues(others)
}

func printIssues(items []*Issue) {
	for _, item := range items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}

func searchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}

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

const IssuesURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}
