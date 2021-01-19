package function

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

// To marshal the reponse in JSON format
// Export the fields in the struct.
type Response struct {
	Artistname string `json:"artist"`
	Featartist string `json:"feat"`
	Songname   string `json:"song"`
	Lyrics     string `json:"lyrics"`
}

func format(text string) string {
	t := strings.ToLower(text)
	t = strings.ReplaceAll(t, " ", "")
	return t
}

func FetchLyrics(artistname, songname string) (response *Response) {

	resp := &Response{}
	c := colly.NewCollector()

	c.OnHTML(".col-xs-12.col-lg-8.text-center", func(e *colly.HTMLElement) {

		var depth int = 8

		artistname := e.DOM.Find(".lyricsh").Text()
		resp.Artistname = strings.ReplaceAll(strings.TrimSpace(artistname), "Lyrics", "")

		songname := strings.ReplaceAll(e.DOM.Find(".col-xs-12.col-lg-8.text-center > b").Text(), "\"", "")
		resp.Songname = strings.TrimSpace(songname)

		feat := e.DOM.Find(".feat")

		if feat.Length() > 0 {
			depth = 10
			resp.Featartist = feat.Text()
		}

		lyrics := e.DOM.Find(".col-xs-12.col-lg-8.text-center > div:nth-child(" + strconv.Itoa(depth) + ")").Text()
		resp.Lyrics = strings.TrimSpace(lyrics)

	})

	url := fmt.Sprintf("https://www.azlyrics.com/lyrics/%s/%s.html", artistname, songname)
	c.Visit(url)
	return resp
}

// Handle a function invocation
func Handle(w http.ResponseWriter, req *http.Request) {

	query := req.URL.Query()
	artist := format(query.Get("artist"))
	song := format(query.Get("song"))

	r := FetchLyrics(artist, song)
	m, _ := json.Marshal(r)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(string(m))))
}
