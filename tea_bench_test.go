package tea

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/daryl/zeus"
	fragmenta "github.com/fragmenta/mux"
	"github.com/go-chi/chi"
	"github.com/go-zoo/bone"
	"github.com/gorilla/mux"
	"github.com/gorilla/pat"
	"github.com/harshvladha/hasty"
	"github.com/julienschmidt/httprouter"
	"github.com/squiidz/boneX"
	"github.com/ursiform/bear"
)

// Test the ns/op
func BenchmarkTeaMux(b *testing.B) {
	request, _ := http.NewRequest("GET", "/sd", nil)
	response := httptest.NewRecorder()
	muxx := New()
	muxx.Get("/", http.HandlerFunc(Bench))
	muxx.Get("/a", http.HandlerFunc(Bench))
	muxx.Get("/aas", http.HandlerFunc(Bench))
	muxx.Get("/sd", http.HandlerFunc(Bench))
	for n := 0; n < b.N; n++ {
		muxx.ServeHTTP(response, request)
	}
}

// Test Bone ns/op
func BenchmarkBoneMux(b *testing.B) {
	request, _ := http.NewRequest("GET", "/sd", nil)
	response := httptest.NewRecorder()
	muxx := bone.New()
	muxx.Get("/", http.HandlerFunc(Bench))
	muxx.Get("/a", http.HandlerFunc(Bench))
	muxx.Get("/aas", http.HandlerFunc(Bench))
	muxx.Get("/sd", http.HandlerFunc(Bench))
	for n := 0; n < b.N; n++ {
		muxx.ServeHTTP(response, request)
	}
}

// Test Bone ns/op
func BenchmarkBoneXMux(b *testing.B) {
	request, _ := http.NewRequest("GET", "/sd", nil)
	response := httptest.NewRecorder()
	muxx := bonex.New()
	muxx.Get("/", BenchX)
	muxx.Get("/a", BenchX)
	muxx.Get("/aas", BenchX)
	muxx.Get("/sd", BenchX)
	for n := 0; n < b.N; n++ {
		muxx.ServeHTTP(response, request)
	}
}

func BenchmarkHastyMux(b *testing.B) {
	request, _ := http.NewRequest("GET", "/sd", nil)
	response := httptest.NewRecorder()
	muxx := hasty.New()
	muxx.Get("/", http.HandlerFunc(Bench))
	muxx.Get("/a", http.HandlerFunc(Bench))
	muxx.Get("/aas", http.HandlerFunc(Bench))
	muxx.Get("/sd", http.HandlerFunc(Bench))
	for n := 0; n < b.N; n++ {
		muxx.ServeHTTP(response, request)
	}
}

func BenchmarkFragmentaMux(b *testing.B) {
	request, _ := http.NewRequest("GET", "/sd", nil)
	response := httptest.NewRecorder()
	muxx := fragmenta.New()
	muxx.Get("/", BenchF)
	muxx.Get("/a", BenchF)
	muxx.Get("/aas", BenchF)
	muxx.Get("/sd", BenchF)
	for n := 0; n < b.N; n++ {
		muxx.ServeHTTP(response, request)
	}
}

func BenchmarkChiMux(b *testing.B) {
	request, _ := http.NewRequest("GET", "/sd", nil)
	response := httptest.NewRecorder()
	muxx := chi.NewRouter()
	muxx.Get("/", http.HandlerFunc(Bench))
	muxx.Get("/a", http.HandlerFunc(Bench))
	muxx.Get("/aas", http.HandlerFunc(Bench))
	muxx.Get("/sd", http.HandlerFunc(Bench))
	for n := 0; n < b.N; n++ {
		muxx.ServeHTTP(response, request)
	}
}

// Test httprouter ns/op
func BenchmarkHttpRouterMux(b *testing.B) {
	request, _ := http.NewRequest("GET", "/sd", nil)
	response := httptest.NewRecorder()
	muxx := httprouter.New()
	muxx.Handler("GET", "/", http.HandlerFunc(Bench))
	muxx.Handler("GET", "/a", http.HandlerFunc(Bench))
	muxx.Handler("GET", "/aas", http.HandlerFunc(Bench))
	muxx.Handler("GET", "/sd", http.HandlerFunc(Bench))
	for n := 0; n < b.N; n++ {
		muxx.ServeHTTP(response, request)
	}
}

// Test daryl/zeus ns/op
func BenchmarkZeusMux(b *testing.B) {
	request, _ := http.NewRequest("GET", "/sd/test", nil)
	response := httptest.NewRecorder()
	muxx := zeus.New()
	muxx.GET("/", Bench)
	muxx.GET("/a", Bench)
	muxx.GET("/aas", Bench)
	muxx.GET("/sd/:id", Bench)
	for n := 0; n < b.N; n++ {
		muxx.ServeHTTP(response, request)
	}
}

// Test net/http ns/op
func BenchmarkNetHttpMux(b *testing.B) {
	request, _ := http.NewRequest("GET", "/sd", nil)
	response := httptest.NewRecorder()
	muxx := http.NewServeMux()
	muxx.HandleFunc("/", Bench)
	muxx.HandleFunc("/a", Bench)
	muxx.HandleFunc("/aas", Bench)
	muxx.HandleFunc("/sd", Bench)
	for n := 0; n < b.N; n++ {
		muxx.ServeHTTP(response, request)
	}
}

// Test ursiform/bear ns/op
func BenchmarkBearMux(b *testing.B) {
	request, _ := http.NewRequest("GET", "/sd", nil)
	response := httptest.NewRecorder()
	muxx := bear.New()
	muxx.On("GET", "/", Bench)
	muxx.On("GET", "/a", Bench)
	muxx.On("GET", "/aas", Bench)
	muxx.On("GET", "/sd", Bench)
	for n := 0; n < b.N; n++ {
		muxx.ServeHTTP(response, request)
	}
}

// Test gorilla/mux ns/op
func BenchmarkGorillaMux(b *testing.B) {
	request, _ := http.NewRequest("GET", "/sd", nil)
	response := httptest.NewRecorder()
	muxx := mux.NewRouter()
	muxx.Handle("/", http.HandlerFunc(Bench))
	muxx.Handle("/a", http.HandlerFunc(Bench))
	muxx.Handle("/aas", http.HandlerFunc(Bench))
	muxx.Handle("/sd", http.HandlerFunc(Bench))
	for n := 0; n < b.N; n++ {
		muxx.ServeHTTP(response, request)
	}
}

// Test gorilla/pat ns/op
func BenchmarkGorillaPatMux(b *testing.B) {
	request, _ := http.NewRequest("GET", "/sd", nil)
	response := httptest.NewRecorder()
	muxx := pat.New()
	muxx.Get("/", Bench)
	muxx.Get("/a", Bench)
	muxx.Get("/aas", Bench)
	muxx.Get("/sd", Bench)
	for n := 0; n < b.N; n++ {
		muxx.ServeHTTP(response, request)
	}
}

func Bench(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("b"))
}

func BenchF(rw http.ResponseWriter, req *http.Request) error {
	rw.Write([]byte("b"))
	return nil
}

func BenchX(rw http.ResponseWriter, req *http.Request, args bonex.Args) {
	rw.Write([]byte("b"))
}
