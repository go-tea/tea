package main

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/go-tea/middleware"
	"github.com/go-tea/rh"
	"github.com/go-tea/tea"
)

func main() {
	teaSub := tea.New()

	teaSub.Use(middleware.RequestID)
	teaSub.Use(middleware.RealIP)
	teaSub.Use(middleware.Logger)
	teaSub.Use(middleware.Recoverer)

	rhSub := rh.Mux

	rg, _ := regexp.Compile("foo.*")

	rhSub.HandleFunc(rg, func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte("Hello from rh mux"))
	})

	mux := tea.New().Prefix("/api")
	mux.SubRoute("/rh", rhSub)
	mux.SubRoute("/tea", teaSub)

	uuid := `[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}`
	//ip := regexp.MustCompile(`(([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:)|fe80:(:[0-9a-fA-F]{0,4}){0,4}%[0-9a-zA-Z]{1,}|::(ffff(:0{1,4}){0,1}:){0,1}((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])|([0-9a-fA-F]{1,4}:){1,4}:((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]))`)
	ip := regexp.MustCompile(`^(?:[0-9]{1,3}\.){3}[0-9]{1,3}$`)

	mux.AddRegex("@uuid", uuid)
	mux.AddRegex("@ip", ip)
	mux.Get("/uuid/@uuid", huuid) //"2E9C64A5-FF13-4DC5-A957-F39E39ABDC48"
	mux.Get("/ip/@ip", hip)

	http.ListenAndServe(":8080", mux)
}

func huuid(rw http.ResponseWriter, req *http.Request) {
	vuuid := tea.GetValue(req, "uuid")
	fmt.Println(vuuid)
	fmt.Println(req.URL)

}

func hip(rw http.ResponseWriter, req *http.Request) {
	vuuid := tea.GetValue(req, "ip")
	fmt.Println(vuuid)
	fmt.Println(req.URL)

}
