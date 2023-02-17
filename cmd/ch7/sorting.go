package ch7

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"
)

// Track represnts the data associated to a music track
type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"The Woven Web", "Animals As Leaders", "The Joy of Motion", 2016, length("4m08s")},
	{"Patterns", "Syncatto", "Patterns", 2023, length("4m02s")},
	{"Colony", "Alluvial", "The Deep Longing for Annihilation", 2017, length("2m57s")},
	{"Callisto", "Returning We Hear the Larks", "Larks", 2015, length("4m05s")},
}

// calculates takes the duration of a song fomatted as
// XmYs and returns a duration
func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

// printTracks prints a slice of pointers to Track structs in a table
func PrintTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush()
}

// byArtist defines a type for an slice of tracks
// that will be used to satisfy the sort interface
// in a particular ordering (by artits)
type byArtist []*Track

func (x byArtist) Len() int           { return len(x) }
func (x byArtist) Less(i, j int) bool { return x[i].Artist < x[j].Artist }
func (x byArtist) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type byYear []*Track

func (x byYear) Len() int           { return len(x) }
func (x byYear) Less(i, j int) bool { return x[i].Year < x[j].Year }
func (x byYear) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type customSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

func (x customSort) Len() int           { return len(x.t) }
func (x customSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x customSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }
