package jokes

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Name struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func GetName() (*Name, error) {

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
