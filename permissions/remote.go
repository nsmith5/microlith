package permissions

import (
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type server struct {
	Service
}

func NewRemoteServer(s Service) http.Handler {
	return server{s}
}

func (s server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	ok, _ := s.CanI(string(body))
	if ok {
		io.WriteString(w, "true")
	} else {
		io.WriteString(w, "false")
	}
}

type client struct {
	addr string
}

func (c client) CanI(thing string) (bool, error) {
	resp, err := http.Post(c.addr, "text/plain", strings.NewReader(thing))
	if err != nil {
		return false, err
	}
	body, _ := ioutil.ReadAll(resp.Body)
	if string(body) == "true" {
		return true, nil
	}
	return false, nil
}

func NewRemote(addr string) Service {
	return client{addr}
}
