package rest

import "net/http"

func (s *RESTServer) loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.logger.Debug(r.URL)
		next.ServeHTTP(w, r)
	})
}
