package main

import (
	"fmt"
	"os"
	"time"
)

type Joke struct {
	Setup     string
	Punchline string
}

func main() {
	joke_type := ""
	if len(os.Args) > 1 {
		joke_type = os.Args[1] + "/"
	}

	jokes := []Joke{}
	err := getJoke(joke_type, &jokes)

	if err != nil {
		fmt.Println("Something went wrong: ", err)
	} else {
		fmt.Println(jokes[0].Setup)
		time.Sleep(3 * time.Second)
		fmt.Println(jokes[0].Punchline)
	}
}
