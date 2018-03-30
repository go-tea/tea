package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-tea/tea"
)

var (
	mux = tea.New(Serve, Wrap)
)

func Wrap(mux *tea.Mux) *tea.Mux {
	return mux.Prefix("/api")
}

func Serve(mux *tea.Mux) *tea.Mux {
	mux.Serve = func(rw http.ResponseWriter, req *http.Request) {
		tr := time.Now()
		mux.DefaultServe(rw, req)
		fmt.Println("Serve request from", req.RemoteAddr, "in", time.Since(tr))
	}
	return mux
}

func main() {
	// Custom 404
	mux.NotFound(Handler404)
	// Handle with any http method, Handle takes http.Handler as argument.
	//	mux.Handle("/index", http.HandlerFunc(homeHandler))
	mux.Handle("/index", homeHandler)
	mux.Handle("/index/:var/info/:test", varHandler)
	// Get, Post etc... takes http.HandlerFunc as argument.
	mux.Post("/home", homeHandler)
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
	rw.Write([]byte("These are not the droids you're looking for ..."))
}
