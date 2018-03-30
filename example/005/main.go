// +build go1.7

package main

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-tea/middleware"
	"github.com/go-tea/tea"
)

var mux *tea.Mux

func main() {
	mux = tea.New()
	mux.CaseSensitive = true

	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)

	mux.Get("/ctx/:var", handler)
	mux.Get("/ctx2/:var", handler2)
	mux.Get("/ctx3/:var/#id^[0-9]$", handler3)
	mux.Get("/ctx4/:var/:id", handler3).Eval(isString, isNumber)

	http.ListenAndServe(":8080", mux)
}

func handler(rw http.ResponseWriter, req *http.Request) {
	ctx := context.WithValue(req.Context(), "var", tea.GetValue(req, "var"))
	subHandler(rw, req.WithContext(ctx))
}

func handler2(rw http.ResponseWriter, req *http.Request) {
	params, err := tea.Params(req)

	if err != nil {
		fmt.Errorf("params.Params %s", err)
	}
	val := params.Get("var")
	ctx := context.WithValue(req.Context(), "var", val)
	subHandler(rw, req.WithContext(ctx))
}

func handler3(rw http.ResponseWriter, req *http.Request) {
	params, err := tea.Params(req)

	if err != nil {
		fmt.Errorf("params.Params %s", err)
	}
	val := params.Get("var")
	id := params.Get("id")
	fmt.Fprintf(rw, "URL.Path = %q\n", req.URL.Path)
	rw.Write([]byte(val + "\n"))
	rw.Write([]byte(id))

}

func subHandler(rw http.ResponseWriter, req *http.Request) {
	val := req.Context().Value("var")
	rw.Write([]byte(val.(string)))

}

// Evaluator which check if the url parameters is a number
func isNumber(str string) bool {
	if _, err := strconv.Atoi(str); err == nil {
		return true
	}
	return false
}

func isString(str string) bool {
	return true
}
