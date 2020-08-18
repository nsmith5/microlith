package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/nsmith5/microlith/permissions"
)

type API struct {
	ps permissions.Service
}

func (a API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	thing := strings.ReplaceAll(r.URL.Path, "/", "")
	ok, err := a.ps.CanI(thing)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	if ok {
		fmt.Fprintf(w, "Oh yeahh you can really %s\n", thing)
	} else {
		fmt.Fprintf(w, "Uh oh, looks like you can never %s\n", thing)
	}
}

func main() {
	permsURL := flag.String("perms-url", "", "URL of permissions service")
	flag.Parse()

	switch {
	case len(os.Args) > 1 && os.Args[1] == "permissions":
		// If invoked with `./my-app permissions` we're just running the permissions
		// microservice
		service := permissions.NewSimple()
		handler := permissions.NewRemoteServer(service)
		http.ListenAndServe(":8081", handler)
	default:
		// If invoked as `./my-app` we're running the product main entrypoint
		if *permsURL == "" {
			// No permiossions service URL provided. We're running
			// in monolith mode
			http.ListenAndServe(":8080", API{permissions.NewSimple()})
		} else {
			// Use the permissions URL provided to make a remote client. We're
			// in microservice mode.
			http.ListenAndServe(":8080", API{permissions.NewRemote(*permsURL)})
		}
	}
}
