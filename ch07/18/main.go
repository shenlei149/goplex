package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

func main() {
	dec := xml.NewDecoder(os.Stdin)
	var stack []*Element // stack of element names
	var root Node
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlbuild: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			cur := &Element{tok.Name, tok.Attr, nil}
			if len(stack) == 0 {
				root = cur
			} else {
				parent := stack[len(stack)-1]
				parent.Children = append(parent.Children, cur)
			}
			stack = append(stack, cur) // push
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop
		case xml.CharData:
			parent := stack[len(stack)-1]
			parent.Children = append(parent.Children, CharData(tok))
		}
	}

	fmt.Println(root)
}

type Node interface{} // CharData or *Element
type CharData string
type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}

func (n *Element) String() string {
	b := &bytes.Buffer{}
	visit(n, b, 0)
	return b.String()
}

func visit(n Node, w io.Writer, depth int) {
	switch n := n.(type) {
	case *Element:
		fmt.Fprintf(w, "%*s%s %s\n", depth*4, "", n.Type.Local, n.Attr)
		for _, child := range n.Children {
			visit(child, w, depth+1)
		}
	case CharData:
		fmt.Fprintf(w, "%*s%q\n", depth*4, "", n)
	}
}
