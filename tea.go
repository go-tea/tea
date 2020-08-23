package tea

import (
	"net/http"
	"strings"
)

// Mux have routes and a notFound handler
// Route: all the registred route
// notFound: 404 handler, default http.NotFound if not provided
type Mux struct {
	compile       compileSet
	Routes        map[string][]*Route
	prefix        string
	notFound      http.Handler
	Serve         func(rw http.ResponseWriter, req *http.Request)
	CaseSensitive bool
	middlewares   []func(http.Handler) http.Handler
	ghandler      http.Handler
}

var (
	static = "static"
	method = []string{"GET", "POST", "PUT", "DELETE", "HEAD", "PATCH", "OPTIONS"}
)

//type adapter func(http.Handler) http.Handler

// New create a pointer to a Mux instance
func New(middlewares ...func(http.Handler) http.Handler) *Mux {
	m := &Mux{Routes: make(map[string][]*Route), Serve: nil, CaseSensitive: true}
	m.Use(middlewares...)

	if m.Serve == nil {
		m.Serve = m.DefaultServe
	}
	m.compile = compileVars

	return m
}

func (m *Mux) Use(middlewares ...func(http.Handler) http.Handler) {
	if m.ghandler != nil {
		panic("tea: all middlewares must be defined before routes on a mux")
	}
	m.middlewares = append(m.middlewares, middlewares...)
}

// Prefix set a default prefix for all routes registred on the router
func (m *Mux) Prefix(p string) *Mux {
	m.prefix = strings.TrimSuffix(p, "/")
	return m
}

// DefaultServe is the default http request handler
func (m *Mux) DefaultServe(rw http.ResponseWriter, req *http.Request) {
	// Check if a route match
	if !m.parse(rw, req) {
		// Check if it's a static ressource
		if !m.staticRoute(rw, req) {
			// Check if the request path doesn't end with /
			if !m.validate(rw, req) {
				// Check if same route exists for another HTTP method
				if !m.otherMethods(rw, req) {
					m.HandleNotFound(rw, req)
				}
			}
		}
	}
}

// ServeHTTP pass the request to the serve method of Mux
func (m *Mux) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if !m.CaseSensitive {
		req.URL.Path = strings.ToLower(req.URL.Path)
	}
	//	m.Serve(rw, req)
	m.ghandler = chain(m.middlewares, http.HandlerFunc(m.Serve))
	m.ghandler.ServeHTTP(rw, req)
}

// AddRegex adds a ":named" regular expression
func (m *Mux) AddRegex(name string, regex interface{}) error {
	return m.compile.Set(name, regex)
}
