package middleware

import (
	"net/http"
)

func AssetsCache(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Cache-Control", "private, max-age=86400")
		next.ServeHTTP(w, r)
	})
}
