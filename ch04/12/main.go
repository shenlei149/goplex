package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	comics := loadInfo()

	index := buildIndex(comics)

	for {
		fmt.Println("Enter the key word:")
		reader := bufio.NewReader(os.Stdin)
		words, _ := reader.ReadString('\n')
		result := search(strings.Fields(strings.TrimSpace(words)), index)

		for _, item := range result {
			fmt.Println(item.Num, item.Title)
		}
	}
}

func search(words []string, index map[string][]Comic) []Comic {
	set := make(map[int]Comic)

	for _, word := range words {
		for _, comic := range index[word] {
			set[comic.Num] = comic
		}
	}

	result := make([]Comic, 0, len(set))
	for _, comic := range set {
		result = append(result, comic)
	}

	return result
}

func buildIndex(comics []Comic) map[string][]Comic {
	result := make(map[string][]Comic)

	for _, comic := range comics {
		for _, word := range strings.Fields(comic.Transcript) {
			result[word] = append(result[word], comic)
		}
	}

	return result
}

func loadInfo() []Comic {
	var comics []Comic
	for i := 0; i < 10; i++ {
		url := fmt.Sprintf("https://xkcd.com/%d/info.0.json", i)
		fmt.Println(url)

		resp, err := http.Get(url)
		if err != nil {
			resp.Body.Close()
		}

		var comic Comic
		if err := json.NewDecoder(resp.Body).Decode(&comic); err != nil {
			fmt.Println(err)
			resp.Body.Close()
			continue
		}

		resp.Body.Close()

		comics = append(comics, comic)
	}

	return comics
}

type Comic struct {
	Title      string
	Transcript string
	Num        int
}
