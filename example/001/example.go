package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-tea/tea"
	"github.com/go-tea/tea/serve"
)

func main() {

	mux := tea.New(serve.RealIP, serve.Logger)

	mux.NotFound(Handler404)
	mux.Get("/home", homeHandler)
	mux.Get("/home/:var", varHandler)

	mux.Get("/test/*", func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte(req.RequestURI))
	})

	// Start Listening
	log.Fatal(mux.ListenAndServe(":8080"))
}

func homeHandler(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("WELCOME HOME"))
}

func varHandler(rw http.ResponseWriter, req *http.Request) {
	varr := tea.GetValue(req, "var")
	test := tea.GetValue(req, "test")

	var args = struct {
		First  string
		Second string
	}{varr, test}

	if err := json.NewEncoder(rw).Encode(&args); err != nil {
		panic(err)
	}
}

func Handler404(rw http.ResponseWriter, req *http.Request) {
	rw.WriteHeader(http.StatusNotFound)
	rw.Write([]byte("These are not the droids you're looking for ..."))

}
