package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Joke struct {
	Setup     string `json:"setup"`
	Punchline string `json:"punchline"`
	Id        int    `json:"id"`
}

func getJoke() Joke {
	url := "https://official-joke-api.appspot.com/random_joke"

	req, _ := http.Get(url)

	body, _ := io.ReadAll(req.Body)

	joke := Joke{}

	json.Unmarshal(body, &joke)

	return joke
}

func printJoke(joke Joke) {
	fmt.Print(joke.Setup)
	fmt.Print("\n...\n")
	fmt.Print("\n...\n")
	fmt.Print("\n...\n")
	fmt.Print(joke.Punchline)
}

func main() {
	joke := getJoke()
	printJoke(joke)
}
