package main

import (
	"encoding/json"
	"net/http"
	"time"
)

var req = &http.Client{Timeout: 10 * time.Second}

func getJoke(joke_type string, jokes *[]Joke) error {
	resp, err := req.Get("https://official-joke-api.appspot.com/jokes/" + joke_type + "random")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if joke_type == "" {
		joke := Joke{}
		error := json.NewDecoder(resp.Body).Decode(&joke)
		*jokes = append(*jokes, joke)
		return error
	} else {
		return json.NewDecoder(resp.Body).Decode(jokes)
	}
}
