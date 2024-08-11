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

func getJoke() (Joke, string) {
	url := "https://official-joke-api.appspot.com/random_joke"

	req, err := http.Get(url)

	if err != nil {
		errorMessage := fmt.Sprintf("Error occurred while sending the HTTP request: %v", err.Error())
		return Joke{}, errorMessage
	}

	if req.StatusCode != http.StatusOK {
		errorMessage := fmt.Sprintf("Received non 200 status code: %v", req.Status)
		return Joke{}, errorMessage
	}

	body, err := io.ReadAll(req.Body)
	if err != nil {
		errorMessage := fmt.Sprintf("Error occurred while reading the response body: %v", err.Error())
		return Joke{}, errorMessage
	}

	defer req.Body.Close()

	joke := Joke{}
	err = json.Unmarshal(body, &joke)
	if err != nil {
		errorMessage := fmt.Sprintf("Error occurred while parsing the JSON response: %v", err.Error())
		return Joke{}, errorMessage
	}

	return joke, ""
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

func main() {
	joke, errorMessage := getJoke()

	if errorMessage != "" {
		fmt.Println(errorMessage)
		return
	}

	printJoke(joke)
}
