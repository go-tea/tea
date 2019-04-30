package main

import (
	"io/ioutil"
	"net/http"

	"github.com/go-tea/tea"
)

func main() {
	mux := tea.New()

	mux.NotFound(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusTeapot)
	})

	mux.Get("/", defaultHandler)
	mux.Get("/reg/#var^[a-z]$/#var2^[0-9]$", ShowVar)
	mux.Get("/test", defaultHandler)
	mux.Handler("/file/", http.StripPrefix("/file/", http.FileServer(http.Dir("assets"))))

	http.ListenAndServe(":8080", mux)
}

func defaultHandler(rw http.ResponseWriter, req *http.Request) {
	file, _ := ioutil.ReadFile("index.html")
	rw.Write(file)
}

func ShowVar(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte(tea.GetAllValues(req)["var"]))
}
