package middleware

import (
	"net/http"
	"strings"
)

// ContentTypeJson agrega una entrada en el header de cada respuesta
// fijando el Content-Type
func ContentTypeJson(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json;charset=utf8")
		next.ServeHTTP(w, r)
	})
}


// StripTrailingSlash remueve un / al final de URL.Path
func StripTrailingSlash(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.URL.Path[len(r.URL.Path)-1] == '/' {
			r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
		}

		next.ServeHTTP(w, r)
	})
}