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

	boneSub.Use(middleware.RequestID)
	boneSub.Use(middleware.RealIP)
	boneSub.Use(middleware.Logger)
	boneSub.Use(middleware.Recoverer)

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

	muxx := tea.New().Prefix("/api")

	//	muxx.Use(middleware.RequestID)
	//	muxx.Use(middleware.RealIP)
	//	muxx.Use(middleware.Logger)
	//	muxx.Use(middleware.Recoverer)

	muxx.SubRoute("/bone", boneSub)
	muxx.SubRoute("/gorilla", gorrilaSub)
	muxx.SubRoute("/http", httprouterSub)

	http.ListenAndServe(":8080", muxx)
}
