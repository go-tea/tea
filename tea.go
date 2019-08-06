package tea

import (
	"net/http"
	"strings"
)

// Router is the same as a http.Handler
type Router interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}

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
	//	middlewares   []func(http.Handler) http.Handler
	//	handler       http.Handler //// The computed mux handler made of the chained middleware stack and the tree router
}

var (
	static = "static"
	method = []string{"GET", "POST", "PUT", "DELETE", "HEAD", "PATCH", "OPTIONS"}
)

type adapter func(*Mux) *Mux

// New create a pointer to a Mux instance
func New(adapters ...adapter) *Mux {
	m := &Mux{Routes: make(map[string][]*Route), Serve: nil, CaseSensitive: true}
	for _, adap := range adapters {
		adap(m)
	}
	if m.Serve == nil {
		m.Serve = m.DefaultServe
	}
	m.compile = compileVars

	return m
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
	m.Serve(rw, req)
}

// AddRegex adds a ":named" regular expression
func (m *Mux) AddRegex(name string, regex interface{}) error {
	return m.compile.Set(name, regex)
}
