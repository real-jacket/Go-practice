package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

// Movie is a info of movie
type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

var titles []struct{ Title string }

//IssuesURL about github issues
const IssuesURL = "https://api.github.com/search/issues"

//IssuesSearchReasult what we need
type IssuesSearchReasult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

// Issue struct
type Issue struct {
	Number   int
	HTMLURL  string `json:"html_url"`
	Title    string
	State    string
	User     *User
	CreateAt time.Time `json:"created_at"`
	Body     string    // in Markdown format
}

// User info
type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

// SearchIssues queries the GitHub issue tracker.
func SearchIssues(terms []string) (*IssuesSearchReasult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchReasult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}

	resp.Body.Close()
	return &result, nil
}

func main() {
	var movies = []Movie{
		{
			Title: "Casablance",
			Year:  1942,
			Color: false,
			Actors: []string{
				"Humpher Bogart",
				"Ingrid Begerman",
			},
		},
		{
			Title: "Cool Hand Luke",
			Year:  1967,
			Color: true,
			Actors: []string{
				"Paul Newman",
			},
		},
		{
			Title: "Bullitt",
			Year:  1968,
			Color: true,
			Actors: []string{
				"Steve McQueen",
				"Jacqueline Bisset",
			},
		},
	}

	data, err := json.MarshalIndent(movies, "", "    ")
	if err != nil {
		log.Fatalf("JSON　marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)

	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("JSON　unmarshaling failed: %s", err)
	}

	fmt.Println(titles)

	result, err := SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
}
