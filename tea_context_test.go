// +build go1.7

package tea

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type key interface{}
type value interface{}

func TestRoutingVariableWithContext(t *testing.T) {
	var (
		expected = "variable"
		got      string
		mux      = New()
		w        = httptest.NewRecorder()
	)

	appFn := func(w http.ResponseWriter, r *http.Request) {
		got = GetValue(r, "vartest")
	}

	var k key
	var v value

	k = "key"
	v = "customValue"

	middlewareFn := func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), k, v)
		newReq := r.WithContext(ctx)
		appFn(w, newReq)
	}

	mux.Get("/:vartest", http.HandlerFunc(middlewareFn))
	r, err := http.NewRequest("GET", fmt.Sprintf("/%s", expected), nil)
	if err != nil {
		t.Fatal(err)
	}
	mux.ServeHTTP(w, r)

	if got != expected {
		t.Fatalf("expected %s, got %s", expected, got)
	}
}

func BenchmarkVariableWithContext(b *testing.B) {
	var (
		expected = "variable"
		got      string
		mux      = New()
		w        = httptest.NewRecorder()
	)

	appFn := func(w http.ResponseWriter, r *http.Request) {
		got = GetValue(r, "vartest")
	}

	var k key
	var v value

	k = "key"
	v = "customValue"

	middlewareFn := func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), k, v)
		newReq := r.WithContext(ctx)
		appFn(w, newReq)
	}

	mux.Get("/:vartest", http.HandlerFunc(middlewareFn))
	r, err := http.NewRequest("GET", fmt.Sprintf("/%s", expected), nil)
	if err != nil {
		b.Fatal(err)
	}
	for n := 0; n < b.N; n++ {
		mux.ServeHTTP(w, r)
	}

	if got != expected {
		b.Fatalf("expected %s, got %s", expected, got)
	}
}

/*
func BenchmarkVariableWithParams(b *testing.B) {
	var (
		expected = "variable"
		got      string
		mux      = New()
		w        = httptest.NewRecorder()
	)

	appFn := func(w http.ResponseWriter, r *http.Request) {
		params, _ := Params(r)
		got = params.Get("vartest")
	}

	middlewareFn := func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "key", "customValue")
		newReq := r.WithContext(ctx)
		appFn(w, newReq)
	}

	mux.Get("/:vartest", http.HandlerFunc(middlewareFn))
	r, err := http.NewRequest("GET", fmt.Sprintf("/%s", expected), nil)
	if err != nil {
		b.Fatal(err)
	}
	for n := 0; n < b.N; n++ {
		mux.ServeHTTP(w, r)
	}

	if got != expected {
		b.Fatalf("expected %s, got %s", expected, got)
	}
}
*/
