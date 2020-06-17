package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	tplt.Execute(w, db) // ignore error
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if price, ok := db[item]; ok {
		fmt.Fprintf(w, "%s\n", price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

var tplt = template.Must(template.New("trackTable").Parse(`
<html>
  <body>
    <table>
      <tr>
	    <th>Item</th>
	    <th>Price</th>
	  </tr>
      {{ range $key, $value := . }}
      <tr>
        <td>{{$key}}</td>
        <td>{{$value}}</td>
      </tr>
      {{end}}
    </table>
  </body>
</html>`))
