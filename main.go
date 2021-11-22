package main

import (
	"log"
	"net/http"

	"wildfire/jokes"
)

func main() {
	http.HandleFunc("/api/v0/joke", jokes.HandleJoke)
	if err := http.ListenAndServe(":5000", nil); err != nil {
		log.Println(err)
	}
}

//TODO: consider other logging options. Logrus has been good in the past.
//log is lightweight but may not support future needs when it comes to logging.
//Since no specific needs are specified we use the simplest/smallest solution
//available out of the box

//TODO: consider using an api key system for purposes of tracking use and/or
//abuse.

//TODO: consider using a router. With the addition of more and more endpoints a
//good router library would add readability and allow easier additions for new
//endpoints
