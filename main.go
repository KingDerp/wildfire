package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/joke", handleJoke)
	if err := http.ListenAndServe(":5000", nil); err != nil {
		//TODO: this is not a good way to handle this error. We don't want to
		//kill the program every time there is a bad or unrecognized call from
		//the third party api
		log.Fatalln(err)
	}
}

func handleJoke(w http.ResponseWriter, r *http.Request) {
	name, err := getName()
	if err != nil {
		log.Println(err.Error())
	}

	joke, err := getJoke(name)
	if err != nil {
		log.Println(err.Error())
	}

	log.Println(r.Method, r.RequestURI)
	fmt.Fprintln(w, joke.Value.Joke)
}

func getJoke(name *Name) (*Joke, error) {

	req, err := http.NewRequest("GET", "http://api.icndb.com/jokes/random", bytes.NewBuffer([]byte{}))
	if err != nil {
		log.Println("error creating New Joke http request object")
		return nil, err
	}

	q := req.URL.Query()
	q.Add("firstName", name.FirstName)
	q.Add("lastName", name.LastName)
	q.Add("limitTo", "nerdy")
	req.URL.RawQuery = q.Encode()

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("error getting new joke")
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("error converting response body to bytes")
		return nil, err
	}

	joke := &Joke{}

	err = json.Unmarshal(body, joke)
	if err != nil {
		log.Println("error demarshalling joke")
		return nil, err
	}

	return joke, nil
}

func getName() (*Name, error) {

	resp, err := http.Get("https://names.mcquay.me/api/v0/")
	if err != nil {
		log.Println("error getting names")
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("error converting response body to bytes")
		return nil, err
	}

	name := &Name{}

	err = json.Unmarshal(body, name)
	if err != nil {
		log.Println("error demarshalling name")
		return nil, err
	}

	return name, nil
}

type Name struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type Joke struct {
	Type  string `json:"type"`
	Value *Value `json:"value"`
}

type Value struct {
	ID         int      `json:"id"`
	Joke       string   `json:"joke"`
	Categories []string `json:"categories"`
}

//TODO: consider other logging options. log is lightweight but may not support
//future needs when it comes to logging. Since no specific needs are specified
//we use the simplest/smallest solution available out of the box

//TODO: consider using an api key system for purposes of tracking use and/or
//abuse.
