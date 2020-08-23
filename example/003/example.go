package main

import (
	"net/http"

	"github.com/go-tea/middleware"
	"github.com/go-tea/tea"
	"github.com/gorilla/mux"
	"github.com/julienschmidt/httprouter"
)

func main() {
	boneSub := tea.New()

	gorrilaSub := mux.NewRouter()
	httprouterSub := httprouter.New()

	boneSub.Get("/test", func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte("Hello from bone mux"))
	})

	gorrilaSub.HandleFunc("/test", func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte("Hello from gorilla mux"))
	})

	httprouterSub.GET("/test", func(rw http.ResponseWriter, req *http.Request, _ httprouter.Params) {
		rw.Write([]byte("Hello from httprouter mux"))
	})

	muxx := tea.New(middleware.RealIP, middleware.Recoverer, middleware.Logger).Prefix("/api")
	muxx.SubRoute("/bone", boneSub)
	muxx.SubRoute("/gorilla", gorrilaSub)
	muxx.SubRoute("/http", httprouterSub)

	http.ListenAndServe(":8080", muxx)
}
