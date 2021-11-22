package jokes

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Joke struct {
	Type  string `json:"type"`
	Value *Value `json:"value"`
}

type Value struct {
	ID         int      `json:"id"`
	Joke       string   `json:"joke"`
	Categories []string `json:"categories"`
}

func GetJoke(name *Name) (*Joke, error) {

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
