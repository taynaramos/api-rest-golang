package middleware

import "net/http"

// adicionar o "Content-type": "application/json" no header de todas as requisições que usar o middleware
func ContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")
		next.ServeHTTP(w, r)
	})
}
