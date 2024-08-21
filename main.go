package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

var jokes = []string{
	"Why don't scientists trust atoms? Because they make up everything!",
	"Why did the chicken join a band? Because it had the drumsticks!",
}

var names = []string{"John", "Jane"}
var occupations = []string{"developer", "teacher"}
var devices = []string{"laptop", "smartphone"}
var bodyParts = []string{"head", "wrist"}
var moods = []string{"happy", "sad"}
var actions = []string{"playing music", "displaying a funny meme"}

func jokeHandler(w http.ResponseWriter, r *http.Request) {
	joke := jokes[rand.Intn(len(jokes))]
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, joke)
}

func madlibHandler(w http.ResponseWriter, r *http.Request) {
	madlib := fmt.Sprintf(
		"%s is a %d-year-old %s who has been struggling with a lot of job-related stress. "+
			"He/she decided to try a new application to relieve stress, which runs on a %s to help improve his/her mood. "+
			"The application senses his/her mood through a device he/she wears on his/her %s. "+
			"When the device senses that he/she is %s, it responds by %s.",
		names[rand.Intn(len(names))],
		20+rand.Intn(40),
		occupations[rand.Intn(len(occupations))],
		devices[rand.Intn(len(devices))],
		bodyParts[rand.Intn(len(bodyParts))],
		moods[rand.Intn(len(moods))],
		actions[rand.Intn(len(actions))],
	)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, madlib)
}

func main() {
	http.HandleFunc("/joke", jokeHandler)
	http.HandleFunc("/madlib", madlibHandler)

	log.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
