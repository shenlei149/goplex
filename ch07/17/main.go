package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

func main() {
	input := toMyElements(os.Args[1:])

	dec := xml.NewDecoder(os.Stdin)
	var stack []myElement // stack of element names
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, toMyElement(tok)) // push
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop
		case xml.CharData:
			if containsAll(stack, input) {
				fmt.Printf("%s: %s\n", Join(stack, " "), tok)
			}
		}
	}
}

type myElement struct {
	name  string
	id    string
	class []string
}

func toMyElement(e xml.StartElement) myElement {
	name := e.Name.Local
	var id string
	var class []string
	for _, attr := range e.Attr {
		switch attr.Name.Local {
		case "id":
			id = attr.Value
		case "class":
			class = strings.Fields(attr.Value)
		}
	}

	return myElement{name, id, class}
}

func toMyElements(s []string) []myElement {
	var ret []myElement
	for _, v := range s {
		ret = append(ret, stringToMyElement(v))
	}
	return ret
}

// name#id.class1.class2
func stringToMyElement(s string) myElement {
	sep := regexp.MustCompile(`#|\.`)
	words := sep.Split(s, -1)
	name := words[0]
	var id string
	var class []string
	if len(words) > 1 {
		if strings.Contains(s, "#") {
			id = words[1]
			class = words[2:]
		} else {
			class = words[1:]
		}
	}
	return myElement{name, id, class}
}

// containsAll reports whether x contains the elements of y, in order.
func containsAll(x, input []myElement) bool {
	for len(input) <= len(x) {
		if len(input) == 0 {
			return true
		}
		if satisfy(x[0], input[0]) {
			input = input[1:]
		}
		x = x[1:]
	}
	return false
}

func satisfy(html, input myElement) bool {
	if html.name != input.name {
		return false
	}

	if input.id != "" && html.id != input.id {
		return false
	}

	if len(input.class) > 0 && !containsAllString(html.class, input.class) {
		return false
	}

	return true
}

func containsAllString(x, y []string) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		if x[0] == y[0] {
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}

func Join(s []myElement, sep string) string {
	var ret []string
	for _, v := range s {
		ret = append(ret, v.String())
	}

	return strings.Join(ret, sep)
}

func (e *myElement) String() string {
	ret := e.name
	if e.id != "" {
		ret += "#" + e.id
	} else if len(e.class) > 0 {
		ret += "." + strings.Join(e.class, ".")
	}
	return ret
}
