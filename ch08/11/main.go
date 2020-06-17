package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println(mirroredQuery())
}

func mirroredQuery() string {
	responses := make(chan string, 3)
	ctx, cancel := context.WithCancel(context.Background())
	go func() { responses <- request(ctx, "https://gopl.io") }()
	go func() { responses <- request(ctx, "https://gopl.io") }()
	go func() { responses <- request(ctx, "https://gopl.io") }()
	result := <-responses // return the quickest response
	cancel()
	return result
}

func request(ctx context.Context, url string) string {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err.Error()
	}
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return err.Error()
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return err.Error()
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err.Error()
	}

	return string(bodyBytes)
}
