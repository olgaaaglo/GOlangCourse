package main 

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"sort"
)

var numRecords = flag.Int("nr", 0, "numRecords")
var sortBy = flag.String("sb", "", "sortBy")

type Song struct {
	Rank int
    Title string
    Artist string
    Album string
    Year string
}

type Songs struct {
	Songs []Song
}

func (songs* Songs) read() {
	file, err := os.Open("songs.json")
	if err != nil {
		fmt.Print(err)
	}

	data := make([]byte, 100000)
	count, err := file.Read(data)
	if err != nil {
		fmt.Println("error:", err)
	}

	err = json.Unmarshal(data[:count], songs)
    if err != nil {
		fmt.Println("error:", err)
    }
}

func (songs Songs) sort() {
	if *sortBy != "" {
		sort.Slice(songs.Songs, func(i, j int) bool {
			switch *sortBy {
			case "Rank":
				return songs.Songs[i].Rank < songs.Songs[j].Rank
			case "Title":
				return songs.Songs[i].Title < songs.Songs[j].Title
			case "Artist":
				return songs.Songs[i].Artist < songs.Songs[j].Artist
			case "Album":
				return songs.Songs[i].Album < songs.Songs[j].Album
			case "Year":
				return songs.Songs[i].Year < songs.Songs[j].Year
			}
			return false
		})
	}

	if *numRecords > 0 {
		for i := 0; i < *numRecords; i++ {
			fmt.Println(songs.Songs[i])
		}
	}
}

func main() {
	flag.Parse()

	var songs Songs  
	songs.read()
	songs.sort()
	
}