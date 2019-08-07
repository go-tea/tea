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

type key interface{}
type value interface{}

func rootHandler(rw http.ResponseWriter, req *http.Request) {

	var k key
	var v value

	k = "var"
	v = tea.GetValue(req, "var")

	ctx := context.WithValue(req.Context(), k, v)
	subHandler(rw, req.WithContext(ctx))
}

func subHandler(rw http.ResponseWriter, req *http.Request) {
	val := req.Context().Value("var")
	rw.Write([]byte(val.(string)))
}
