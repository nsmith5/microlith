package main

import (
	"fmt"
	"net/http"
	"strings"
)

var okthings = map[string]struct{}{
	"walk": struct{}{},
	"talk": struct{}{},

	// Nah, we're not that cool
	// "walk the walk": struct{}
}

type API struct{}

func (a API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	thing := strings.ReplaceAll(r.URL.Path, "/", "")
	_, ok := okthings[thing]
	if ok {
		fmt.Fprintf(w, "Oh yeahh you can really %s\n", thing)
	} else {
		fmt.Fprintf(w, "Uh oh, looks like you can never %s\n", thing)
	}
}

func main() {
	http.ListenAndServe(":8080", API{})
}
