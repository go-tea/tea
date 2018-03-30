// +build go1.7

package main

import (
	"context"
	"net/http"

	"github.com/go-tea/tea"
)

func main() {
	mux := tea.New()
	mux.CaseSensitive = true

	mux.Get("/ctx/:var", rootHandler)

	http.ListenAndServe(":8080", mux)
}

func rootHandler(rw http.ResponseWriter, req *http.Request) {
	ctx := context.WithValue(req.Context(), "var", tea.GetValue(req, "var"))
	subHandler(rw, req.WithContext(ctx))
}

func subHandler(rw http.ResponseWriter, req *http.Request) {
	val := req.Context().Value("var")
	rw.Write([]byte(val.(string)))
}
