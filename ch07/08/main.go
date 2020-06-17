package main

import (
	"sort"

	"track"
)

func main() {
	track.PrintTracks(track.Tracks)
	byColumns := track.ByColumns(track.Tracks)
	byColumns.SelectCol("year")
	byColumns.SelectCol("title")
	sort.Sort(byColumns)
	track.PrintTracks(track.Tracks)
	sortByStable(track.Tracks)
	track.PrintTracks(track.Tracks)
}

func sortByStable(t []*track.Track) {
	sort.Stable(track.ByArtist(t))
	sort.Stable(track.ByYear(t))
}
