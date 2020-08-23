package tea

import (
	"net/http"
	"strings"
)

// Register the route in the router
func (m *Mux) Register(method string, path string, handler http.Handler) *Route {
	return m.register(method, path, handler)
}

// Get add a new route to the Mux with the Get method
func (m *Mux) Get(path string, handler http.HandlerFunc) *Route {
	return m.register("GET", path, handler)
}

// Post add a new route to the Mux with the Post method
func (m *Mux) Post(path string, handler http.HandlerFunc) *Route {
	return m.register("POST", path, handler)
}

// Put add a new route to the Mux with the Put method
func (m *Mux) Put(path string, handler http.HandlerFunc) *Route {
	return m.register("PUT", path, handler)
}

// Delete add a new route to the Mux with the Delete method
func (m *Mux) Delete(path string, handler http.HandlerFunc) *Route {
	return m.register("DELETE", path, handler)
}

// Head add a new route to the Mux with the Head method
func (m *Mux) Head(path string, handler http.HandlerFunc) *Route {
	return m.register("HEAD", path, handler)
}

// Patch add a new route to the Mux with the Patch method
func (m *Mux) Patch(path string, handler http.HandlerFunc) *Route {
	return m.register("PATCH", path, handler)
}

// Options add a new route to the Mux with the Options method
func (m *Mux) Options(path string, handler http.HandlerFunc) *Route {
	return m.register("OPTIONS", path, handler)
}

// GetSH add a new route to the Mux with the Get method
func (m *Mux) GetSH(path string, handler http.Handler) *Route {
	return m.register("GET", path, handler)
}

// PostSH add a new route to the Mux with the Post method
func (m *Mux) PostSH(path string, handler http.Handler) *Route {
	return m.register("POST", path, handler)
}

// PutSH add a new route to the Mux with the Put method
func (m *Mux) PutSH(path string, handler http.Handler) *Route {
	return m.register("PUT", path, handler)
}

// DeleteSH add a new route to the Mux with the Delete method
func (m *Mux) DeleteSH(path string, handler http.Handler) *Route {
	return m.register("DELETE", path, handler)
}

// HeadSH add a new route to the Mux with the Head method
func (m *Mux) HeadSH(path string, handler http.Handler) *Route {
	return m.register("HEAD", path, handler)
}

// PatchSH add a new route to the Mux with the Patch method
func (m *Mux) PatchSH(path string, handler http.Handler) *Route {
	return m.register("PATCH", path, handler)
}

// OptionsSH add a new route to the Mux with the Options method
func (m *Mux) OptionsSH(path string, handler http.Handler) *Route {
	return m.register("OPTIONS", path, handler)
}

// NotFound the mux custom 404 handler
func (m *Mux) NotFound(handlerFunc http.HandlerFunc) {
	m.notFound = handlerFunc
}

// Handle is use to pass a func(http.ResponseWriter, *Http.Request) instead of http.Handler
func (m *Mux) Handle(path string, handler http.HandlerFunc) {
	m.Handler(path, handler)
}

// Handler registers  route with all the methods
// Handler add a new route to the Mux without a HTTP method
func (m *Mux) Handler(path string, handler http.Handler) {
	for _, mt := range method {
		m.register(mt, path, handler)
	}
}

// SubRoute register a router as a SubRouter of bone
func (m *Mux) SubRoute(path string, router http.Handler) *Route {
	r := NewRoute(m.prefix+path, router)
	if valid(path) {
		r.Atts += SUB
		for _, mt := range method {
			m.Routes[mt] = append(m.Routes[mt], r)
		}
		return r
	}
	return nil
}

// Temporary way for serving static files
func (m *Mux) AddStatic(path string, dir string) {
	fileServer := http.FileServer(http.Dir(dir))
	newpath := strings.TrimSuffix(path, "*")
	fileHandler := http.StripPrefix(newpath, fileServer)
	m.GetSH(path, fileHandler)
}

// Register the new route in the router with the provided method and handler
func (m *Mux) register(method string, path string, handler http.Handler) *Route {
	r := NewRoute(m.prefix+path, handler)
	r.Method = method
	if valid(path) {
		m.Routes[method] = append(m.Routes[method], r)
		return r
	}
	m.Routes[static] = append(m.Routes[static], r)
	return r
}
