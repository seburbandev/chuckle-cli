package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_BuildApiQuery_Returns_QueryType_Custom(t *testing.T) {

	jokeType := "programming"
	jokeCount := 1
	url := fmt.Sprintf("https://official-joke-api.appspot.com/jokes/%v/random", jokeType)
	expected := ApiQuery{Url: url, QueryType: "custom", JokeType: "programming", JokeCount: jokeCount}

	result := buildApiQuery(jokeType, jokeCount)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("buildApiQuery(%v) = %v, want %v", jokeType, result, expected)
	}

}

func Test_BuildApiQuery_Returns_QueryType_Default(t *testing.T) {

	jokeType := ""
	jokeCount := 1
	url := fmt.Sprintf("https://official-joke-api.appspot.com/jokes/random/%v", jokeCount)
	expected := ApiQuery{Url: url, QueryType: "default", JokeType: "", JokeCount: jokeCount}

	result := buildApiQuery(jokeType, jokeCount)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("buildApiQuery(%v) = %v, want %v", jokeType, result, expected)
	}
}
