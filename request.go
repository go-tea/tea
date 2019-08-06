package tea

import (
	"context"
	"net/http"
)

var (
	defaultContextIdentifier = &struct {
		name string
	}{
		name: "tea",
	}
)

type urlParam struct {
	key   string
	value string
}

type urlParams []urlParam

// Get returns the URL parameter for the given key, or blank if not found
func (p urlParams) Get(key string) (param string) {

	for i := 0; i < len(p); i++ {
		if p[i].key == key {
			param = p[i].value
			return
		}
	}

	return
}

// ReqVars interface
type ReqVars interface {
	URLParam(pname string) string
}

type requestVars struct {
	ctx        context.Context // holds a copy of it's parent requestVars
	params     urlParams
	formParsed bool
}

// RequestVars returns the request scoped variables tracked by pure
func RequestVars(r *http.Request) ReqVars {

	rv := r.Context().Value(defaultContextIdentifier)
	if rv == nil {
		return new(requestVars)
	}

	return rv.(*requestVars)
}

// Params returns the current routes Params
func (r *requestVars) URLParam(pname string) string {
	return r.params.Get(pname)
}
