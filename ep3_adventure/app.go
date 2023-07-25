package main

import (
	"encoding/json"
	"html/template"
	"net/http"
	"os"
	"strings"

)

type Adventure struct {
	Title  string        `json:"title"`
	Story  []string      `json:"story"`
	Option []StoryOption `json:"options"`
}

type StoryOption struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

func main() {

	data, err := os.ReadFile("gopher.json")
	if err != nil {
		panic(err)
	}

	arcToAdventureMap := make(map[string]Adventure)
	err = json.Unmarshal(data, &arcToAdventureMap)
	if err != nil {
		panic(err)
	}

	defaultMux()

	tmpl := template.Must(template.ParseFiles("layout.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		arcPath := strings.SplitAfter(r.URL.Path, "/")
		arc := "intro"
		if arcPath[1] != "" {
			arc = arcPath[1]
		}
		data := arcToAdventureMap[arc]
		tmpl.Execute(w, data)
	})
	http.ListenAndServe(":8080", nil)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	return mux
}
