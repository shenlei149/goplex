package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
)

var issues []Issue
var issuesPage *template.Template
var listTemplate string = `
<h1>issues</h1>
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>User</th>
  <th>Title</th>
</tr>
{{range .}}
<tr>
  <td><a href='/issue/{{.Number}}'>{{.Number}}</a></td>
  <td><a href='/user/{{.User.Login}}'>{{.User.Login}}</a></td>
  <td><a href='/issue/{{.Number}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`

var users map[string]*User = make(map[string]*User)
var userPage *template.Template
var userTemplate string = `
<h1>{{.Login}}</h1>
<img src='{{.AvatarURL}}' />
<br />
<a target="_blank" href='{{.HTMLURL}}'>{{.Login}}</a>
`

var issueByID map[int]*Issue = make(map[int]*Issue)
var issuePage *template.Template
var issueTemplate string = `
<h1>{{.Number}}</h1>
<a target="_blank" href='{{.HTMLURL}}'>{{.Title}}</a>
<br />
<p>{{.Body}}</p>
`

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: exe owner repo")
		return
	}

	issuesPage = template.Must(template.New("issuelist").Parse(listTemplate))
	userPage = template.Must(template.New("user").Parse(userTemplate))
	issuePage = template.Must(template.New("issue").Parse(issueTemplate))

	issues = listIssue(os.Args[1], os.Args[2])
	for i, issue := range issues {
		users[issue.User.Login] = issue.User
		issueByID[issue.Number] = &issues[i]
	}

	http.HandleFunc("/user/", handleUser)
	http.HandleFunc("/issue/", handleIssue)
	http.HandleFunc("/", handle)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handle(w http.ResponseWriter, r *http.Request) {
	if err := issuesPage.Execute(w, issues); err != nil {
		w.Write([]byte("Fail to execute your requst."))
	}
}

func handleUser(w http.ResponseWriter, r *http.Request) {
	userLogin := path.Base(r.URL.Path)
	if err := userPage.Execute(w, users[userLogin]); err != nil {
		w.Write([]byte("Fail to execute your requst."))
	}
}

func handleIssue(w http.ResponseWriter, r *http.Request) {
	issueID, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		w.Write([]byte("Fail to execute your requst."))
	}

	if err = issuePage.Execute(w, issueByID[issueID]); err != nil {
		w.Write([]byte("Fail to execute your requst."))
	}
}

const APIURL = "https://api.github.com"

func listIssue(owner, repo string) []Issue {
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

	return result
}

type Issue struct {
	Title   string
	Body    string
	Number  int
	HTMLURL string `json:"html_url"`
	User    *User
}

type User struct {
	Login     string
	HTMLURL   string `json:"html_url"`
	AvatarURL string `json:"avatar_url"`
}
