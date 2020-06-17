package track

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var Tracks = []*Track{
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func PrintTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column widths and print table
}

type ByArtist []*Track

func (x ByArtist) Len() int           { return len(x) }
func (x ByArtist) Less(i, j int) bool { return x[i].Artist < x[j].Artist }
func (x ByArtist) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type ByYear []*Track

func (x ByYear) Len() int           { return len(x) }
func (x ByYear) Less(i, j int) bool { return x[i].Year < x[j].Year }
func (x ByYear) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type less func(x, y *Track) bool

func lessByTitle(x, y *Track) bool  { return x.Title < y.Title }
func lessByArtist(x, y *Track) bool { return x.Artist < y.Artist }
func lessByAlbum(x, y *Track) bool  { return x.Album < y.Album }
func lessByYear(x, y *Track) bool   { return x.Year < y.Year }
func lessByLength(x, y *Track) bool { return x.Length < y.Length }

type byColumns struct {
	tracks  []*Track
	columns []*column
}

type column struct {
	name string
	less less
}

func ByColumns(t []*Track) *byColumns {
	return &byColumns{tracks: t}
}

func (c *byColumns) Len() int {
	return len(c.tracks)
}

func (c *byColumns) Swap(i, j int) {
	c.tracks[i], c.tracks[j] = c.tracks[j], c.tracks[i]
}

func (c *byColumns) Less(i, j int) bool {
	x, y := c.tracks[i], c.tracks[j]
	for i := 0; i < len(c.columns); i++ {
		f := c.columns[i].less
		if f(x, y) {
			return true
		} else if f(y, x) {
			return false
		} // m == n
	}
	return false
}

func (c *byColumns) SelectCol(s string) {
	var f less
	switch s {
	case "title":
		f = lessByTitle
	case "artist":
		f = lessByArtist
	case "album":
		f = lessByAlbum
	case "year":
		f = lessByYear
	case "length":
		f = lessByLength
	default:
		s = "title"
		f = lessByTitle
	}

	for i, col := range c.columns {
		if col.name == s {
			if i != 0 {
				c.columns[0], c.columns[i] = c.columns[i], c.columns[0]
			}
			return
		}
	}

	c.columns = append(c.columns, &column{s, f})
	c.columns[0], c.columns[len(c.columns)-1] = c.columns[len(c.columns)-1], c.columns[0]
}

func (c *byColumns) Sort(w http.ResponseWriter, req *http.Request, tplt *template.Template) {
	col := req.URL.Query().Get("by")
	if col != "" {
		c.SelectCol(col)
		sort.Sort(c)
	}

	if err := tplt.Execute(w, c.tracks); err != nil {
		log.Fatal(err)
	}
}
