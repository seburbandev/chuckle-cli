package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
)

func main() {

	jokeType := flag.String("type", "", "Type of joke to fetch")
	jokeCount := 1

	flag.Parse()

	apiQuery := buildApiQuery(*jokeType, jokeCount)

	joke, errorMessage := getJoke(apiQuery)

	if errorMessage != "" {
		fmt.Println(errorMessage)
		return
	}

	printJoke(joke)
}

type Joke struct {
	Setup     string `json:"setup"`
	Punchline string `json:"punchline"`
	Type      string `json:"type"`
	Id        int    `json:"id"`
}

type ApiQuery struct {
	Url       string
	QueryType string
	JokeType  string
	JokeCount int
}

func getJoke(aq ApiQuery) (Joke, string) {

	req, err := http.Get(aq.Url)
	errorMessage := ""

	if err != nil {
		errorMessage = fmt.Sprintf("Error occurred while sending the HTTP request: %v", err.Error())
		return Joke{}, errorMessage
	}

	if req.StatusCode != http.StatusOK {
		errorMessage = fmt.Sprintf("Received non 200 status code: %v", req.Status)
		return Joke{}, errorMessage
	}

	body, err := io.ReadAll(req.Body)
	if err != nil {
		errorMessage = fmt.Sprintf("Error occurred while reading the response body: %v", err.Error())
		return Joke{}, errorMessage
	}

	defer req.Body.Close()

	if len(body) == 0 || string(body) == "[]" {
		errorMessage = fmt.Sprintf("Error: could not find joke of type '%v'", aq.JokeType)
		return Joke{}, errorMessage
	}

	var jokes []Joke
	err = json.Unmarshal(body, &jokes)
	if err != nil {
		errorMessage = fmt.Sprintf("Error occurred while parsing the JSON response: %v", err.Error())
		return Joke{}, errorMessage
	}

	return jokes[0], errorMessage
}

func printJoke(joke Joke) {
	fmt.Println(joke.Setup)
	fmt.Println("...")
	fmt.Println("")
	fmt.Println("...")
	fmt.Println("")
	fmt.Println("...")
	fmt.Println(joke.Punchline)
}

func buildApiQuery(jokeType string, jokeCount int) ApiQuery {

	queryType := "default"

	if jokeType != "" {
		queryType = "custom"
	}

	aq := ApiQuery{
		QueryType: queryType,
		JokeType:  jokeType,
		JokeCount: jokeCount,
	}

	baseUrl := "https://official-joke-api.appspot.com"

	if aq.QueryType == "default" {
		baseUrl += fmt.Sprintf("/jokes/random/%v", aq.JokeCount)
	} else {
		baseUrl += fmt.Sprintf("/jokes/%v/random", aq.JokeType)
	}

	aq.Url = baseUrl

	return aq
}
