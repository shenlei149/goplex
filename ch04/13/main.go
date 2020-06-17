package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

const queryURL string = "http://www.omdbapi.com/?apikey=1f3d953a&t="

type Movie struct {
	Title    string
	Poster   string
	Response string
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please enter a movie name")
	} else {
		downloadPostor(url.QueryEscape(strings.Join(os.Args[1:], " ")))
	}
}

func downloadPostor(title string) {
	fmt.Println("Movie name:", title)
	url := queryURL + title
	resp, err := http.Get(url)
	checkError(err)

	b, err := ioutil.ReadAll(resp.Body)

	resp.Body.Close()
	checkError(err)

	var poster Movie
	err = json.Unmarshal(b, &poster)
	checkError(err)

	if poster.Response == "False" {
		fmt.Println("Cannot get information for moive", title)
		return
	}

	if poster.Poster == "N/A" {
		fmt.Println("No poster for", title)
		return
	}

	resp, err = http.Get(poster.Poster)
	checkError(err)

	fileName := poster.Title + filepath.Ext(poster.Poster)
	file, err := os.Create(fileName)
	checkError(err)

	_, err = io.Copy(file, resp.Body)
	checkError(err)

	fmt.Println("Downloaded!")
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
