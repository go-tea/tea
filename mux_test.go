package tea

import (
	"net/http"
	"reflect"
	"testing"
)

func TestMux_Register(t *testing.T) {
	type fields struct {
		Routes   map[string][]*Route
		prefix   string
		notFound http.Handler
		Serve    func(rw http.ResponseWriter, req *http.Request)
	}
	type args struct {
		method  string
		path    string
		handler http.Handler
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Route
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Mux{
				Routes:   tt.fields.Routes,
				prefix:   tt.fields.prefix,
				notFound: tt.fields.notFound,
				Serve:    tt.fields.Serve,
			}
			if got := m.Register(tt.args.method, tt.args.path, tt.args.handler); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Mux.Register() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMux_Get(t *testing.T) {
	type fields struct {
		Routes   map[string][]*Route
		prefix   string
		notFound http.Handler
		Serve    func(rw http.ResponseWriter, req *http.Request)
	}
	type args struct {
		path    string
		handler http.HandlerFunc
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Route
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Mux{
				Routes:   tt.fields.Routes,
				prefix:   tt.fields.prefix,
				notFound: tt.fields.notFound,
				Serve:    tt.fields.Serve,
			}
			if got := m.Get(tt.args.path, tt.args.handler); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Mux.GetFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMux_Post(t *testing.T) {
	type fields struct {
		Routes   map[string][]*Route
		prefix   string
		notFound http.Handler
		Serve    func(rw http.ResponseWriter, req *http.Request)
	}
	type args struct {
		path    string
		handler http.HandlerFunc
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Route
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Mux{
				Routes:   tt.fields.Routes,
				prefix:   tt.fields.prefix,
				notFound: tt.fields.notFound,
				Serve:    tt.fields.Serve,
			}
			if got := m.Post(tt.args.path, tt.args.handler); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Mux.PostFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMux_Put(t *testing.T) {
	type fields struct {
		Routes   map[string][]*Route
		prefix   string
		notFound http.Handler
		Serve    func(rw http.ResponseWriter, req *http.Request)
	}
	type args struct {
		path    string
		handler http.HandlerFunc
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Route
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Mux{
				Routes:   tt.fields.Routes,
				prefix:   tt.fields.prefix,
				notFound: tt.fields.notFound,
				Serve:    tt.fields.Serve,
			}
			if got := m.Put(tt.args.path, tt.args.handler); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Mux.PutFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMux_Delete(t *testing.T) {
	type fields struct {
		Routes   map[string][]*Route
		prefix   string
		notFound http.Handler
		Serve    func(rw http.ResponseWriter, req *http.Request)
	}
	type args struct {
		path    string
		handler http.HandlerFunc
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Route
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Mux{
				Routes:   tt.fields.Routes,
				prefix:   tt.fields.prefix,
				notFound: tt.fields.notFound,
				Serve:    tt.fields.Serve,
			}
			if got := m.Delete(tt.args.path, tt.args.handler); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Mux.DeleteFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMux_Head(t *testing.T) {
	type fields struct {
		Routes   map[string][]*Route
		prefix   string
		notFound http.Handler
		Serve    func(rw http.ResponseWriter, req *http.Request)
	}
	type args struct {
		path    string
		handler http.HandlerFunc
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Route
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Mux{
				Routes:   tt.fields.Routes,
				prefix:   tt.fields.prefix,
				notFound: tt.fields.notFound,
				Serve:    tt.fields.Serve,
			}
			if got := m.Head(tt.args.path, tt.args.handler); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Mux.HeadFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMux_Patch(t *testing.T) {
	type fields struct {
		Routes   map[string][]*Route
		prefix   string
		notFound http.Handler
		Serve    func(rw http.ResponseWriter, req *http.Request)
	}
	type args struct {
		path    string
		handler http.HandlerFunc
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Route
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Mux{
				Routes:   tt.fields.Routes,
				prefix:   tt.fields.prefix,
				notFound: tt.fields.notFound,
				Serve:    tt.fields.Serve,
			}
			if got := m.Patch(tt.args.path, tt.args.handler); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Mux.PatchFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMux_Options(t *testing.T) {
	type fields struct {
		Routes   map[string][]*Route
		prefix   string
		notFound http.Handler
		Serve    func(rw http.ResponseWriter, req *http.Request)
	}
	type args struct {
		path    string
		handler http.HandlerFunc
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Route
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Mux{
				Routes:   tt.fields.Routes,
				prefix:   tt.fields.prefix,
				notFound: tt.fields.notFound,
				Serve:    tt.fields.Serve,
			}
			if got := m.Options(tt.args.path, tt.args.handler); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Mux.OptionsFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMux_NotFound(t *testing.T) {
	type fields struct {
		Routes   map[string][]*Route
		prefix   string
		notFound http.Handler
		Serve    func(rw http.ResponseWriter, req *http.Request)
	}
	type args struct {
		handler http.HandlerFunc
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Mux{
				Routes:   tt.fields.Routes,
				prefix:   tt.fields.prefix,
				notFound: tt.fields.notFound,
				Serve:    tt.fields.Serve,
			}
			m.NotFound(tt.args.handler)
		})
	}
}
