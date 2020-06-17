package main

import (
	"html/template"
	"log"
	"net/http"

	"track"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		track.ByColumns(track.Tracks).Sort(w, req, tplt)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

var tplt = template.Must(template.New("trackTable").Parse(`
<html>
  <body>
    <table>
      <tr>
	    <th><a href="./?by=title">Title</a></th>
	    <th><a href="./?by=artist">Artist</a></th>
	    <th><a href="./?by=album">Album</a></th>
	    <th><a href="./?by=year">Year</a></th>
	    <th><a href="./?by=length">Length</a></th>
	  </tr>
      {{range .}}
      <tr>
        <td>{{.Title}}</td>
        <td>{{.Artist}}</td>
        <td>{{.Album}}</td>
        <td>{{.Year}}</td>
        <td>{{.Length}}</td>
      </tr>
      {{end}}
    </table>
  </body>
</html>`))
