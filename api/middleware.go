package api

import (
	"log"
	"net/http"
)

// this allows us to extract the status code as part of the middleware after the handler has fully executed.
// this will be useful for logging overall handler durations along with the status codes
type statusRecorder struct {
	http.ResponseWriter
	statusCode int
}

func (rw *statusRecorder) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func ExampleMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// before the handler executes
		rec := statusRecorder{w, 200}

		log.Println("This is executing before the handler")

		next.ServeHTTP(&rec, r)
		// after the handler executes

		log.Println("This is executing after the handler")
	})
}
