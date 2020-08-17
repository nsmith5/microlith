package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/nsmith5/microlith/permissions"
)

type API struct {
	ps permissions.Service
}

func (a API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	thing := strings.ReplaceAll(r.URL.Path, "/", "")
	ok := a.ps.CanI(thing)
	if ok {
		fmt.Fprintf(w, "Oh yeahh you can really %s\n", thing)
	} else {
		fmt.Fprintf(w, "Uh oh, looks like you can never %s\n", thing)
	}
}

func main() {
	http.ListenAndServe(":8080", API{permissions.NewSimple()})
}
