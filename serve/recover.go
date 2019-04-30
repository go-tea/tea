package serve

// The original work was derived from Goji's middleware, source:
// https://github.com/zenazn/goji/tree/master/web/middleware

import (
	"net/http"
	"runtime/debug"

	"github.com/go-tea/tea"
)

// Recoverer is a middleware that recovers from panics, logs the panic (and a
// backtrace), and returns a HTTP 500 (Internal Server Error) status if
// possible. Recoverer prints a request ID if one is provided.
//
// Alternatively, look at https://github.com/pressly/lg middleware pkgs.
func Recoverer(mux *tea.Mux) *tea.Mux {
	mux.Serve = func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rvr := recover(); rvr != nil {
				logEntry := GetLogEntry(r)
				if logEntry != nil {
					logEntry.Panic(rvr, debug.Stack())
				} else {
					debug.PrintStack()
				}

				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()

		mux.DefaultServe(w, r)
	}

	return mux
}
