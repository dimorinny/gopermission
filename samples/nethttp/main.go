package main

import (
	"net/http"

	"github.com/dimorinny/gopermission"
)

type HasQwertyHeader struct{}

func (ch HasQwertyHeader) HasPermission(request *http.Request) bool {
	return request.Header.Get("Qwerty") != ""
}

func main() {
	permission := gopermission.New(HasQwertyHeader{})
	qwertyMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !permission.IsPermitted(r) {
				w.Write([]byte("Error"))
				return
			}
			next.ServeHTTP(w, r)
		})
	}

	http.Handle("/", qwertyMiddleware(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Granted"))
		}),
	))
	http.ListenAndServe(":9090", nil)
}
