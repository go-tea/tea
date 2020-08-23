// +build go1.7

/********************************
*** Multiplexer for Go        ***
*** Bone is under MIT license ***
*** Code by CodingFerret      ***
*** github.com/go-zoo         ***
*********************************/

package tea

import (
	"context"
	"net/http"
)

// GetAllValues return the req PARAMs
func GetAllValues(req *http.Request) map[string]string {
	values, ok := req.Context().Value(defaultContextIdentifier).(map[string]string)
	if ok {
		return values
	}

	return map[string]string{}
}

// serveMatchedRequest is an extension point for Route which allows us to conditionally compile for
// go1.7 and <go1.7
func (r *Route) serveMatchedRequest(rw http.ResponseWriter, req *http.Request, vars map[string]string) {
	ctx := context.WithValue(req.Context(), defaultContextIdentifier, vars)
	newReq := req.WithContext(ctx)
	r.Handler.ServeHTTP(rw, newReq)
}
