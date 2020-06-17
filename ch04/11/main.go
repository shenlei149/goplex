package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 4 {
		usage()
	}

	reader = bufio.NewReader(os.Stdin)

	cmd := os.Args[1]
	owner := os.Args[2]
	repo := os.Args[3]

	switch cmd {
	case "create":
		createIssue(owner, repo)
	case "read":
		readIssue(owner, repo)
	case "update":
		updateIssue(owner, repo)
	case "delete":
		// can not find delete API on https://developer.github.com/v3/issues/
	}
}

func usage() {
	fmt.Println("Usage: create|read|update|delete owner repo")
	os.Exit(1)
}

const APIURL = "https://api.github.com"

var reader *bufio.Reader

func doHttpRequest(issue Issue, url, httpMethod string) {
	jsonRequest, _ := json.Marshal(issue)
	fmt.Println(string(jsonRequest))

	req, _ := http.NewRequest(httpMethod, url, bytes.NewBuffer(jsonRequest))
	req.SetBasicAuth("shenlei149", os.Getenv("GITHUB_PASS"))
	fmt.Println("Pass:", os.Getenv("GITHUB_PASS"))

	resp, err := http.DefaultClient.Do(req)
	fmt.Println(err)
	fmt.Println(resp.Status)

	resp.Body.Close()
}

func createIssue(owner, repo string) {
	fmt.Println("Enter the title:")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)
	fmt.Println(title)

	fmt.Println("Enter the body:")
	body, _ := reader.ReadString('\n')
	body = strings.TrimSpace(body)
	fmt.Println(body)

	createRequest := Issue{Title: title, Body: body}

	url := strings.Join([]string{APIURL, "repos", owner, repo, "issues"}, "/")
	fmt.Println(url)

	doHttpRequest(createRequest, url, http.MethodPost)
}

func listIssue(owner, repo string) {
	url := strings.Join([]string{APIURL, "repos", owner, repo, "issues"}, "/")
	fmt.Println(url)

	resp, _ := http.Get(url)
	fmt.Println(resp.Status)

	var result []Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		fmt.Println(err)
		resp.Body.Close()
	}

	resp.Body.Close()

	for _, issue := range result {
		fmt.Println(issue.Number, issue.Title)
	}
}

func getSingleIssue(owner, repo string) Issue {
	fmt.Println("Enter the number of issue:")
	number, _ := reader.ReadString('\n')
	number = strings.TrimSpace(number)

	url := strings.Join([]string{APIURL, "repos", owner, repo, "issues", number}, "/")
	fmt.Println(url)

	resp, _ := http.Get(url)
	fmt.Println(resp.Status)

	var result Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		fmt.Println(err)
		resp.Body.Close()
	}

	resp.Body.Close()

	return result
}

func readIssue(owner, repo string) {
	listIssue(owner, repo)

	issue := getSingleIssue(owner, repo)
	fmt.Println(issue.Body)
}

func updateIssue(owner, repo string) {
	listIssue(owner, repo)

	issue := getSingleIssue(owner, repo)
	fmt.Println(issue.Body)

	fmt.Println("Enter the new title of issue:")
	title, _ := reader.ReadString('\n')
	newtitle := strings.TrimSpace(title)

	fmt.Println("Enter the new content of issue:")
	body, _ := reader.ReadString('\n')
	newbody := strings.TrimSpace(body)

	if len(newtitle) > 0 || len(newbody) > 0 {
		if len(newtitle) > 0 {
			issue.Title = newtitle
		}

		if len(newbody) > 0 {
			issue.Body = newbody
		}

		url := strings.Join([]string{APIURL, "repos", owner, repo, "issues",
			strconv.Itoa(issue.Number)}, "/")
		fmt.Println(url)

		doHttpRequest(issue, url, http.MethodPatch)
	}
}

type Issue struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	Number int
}
