package jokes

import (
	"fmt"
	"log"
	"net/http"
)

func HandleJoke(w http.ResponseWriter, r *http.Request) {
	name, err := GetName()
	if err != nil {
		log.Println(err)
	}

	joke, err := GetJoke(name)
	if err != nil {
		log.Println(err)
	}

	log.Println(r.Method, r.RequestURI)

	//TODO: this meets the specification of the API. In reality I would expect
	//there to be more structure in the response. I would also check with the
	//stakeholder/project manager/product manager to verify.
	fmt.Fprintln(w, joke.Value.Joke)
}
