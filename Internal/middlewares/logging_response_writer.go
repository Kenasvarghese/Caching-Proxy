package middlewares

import "net/http"

type loggingResponseWriter struct {
	http.ResponseWriter
	status int
}

func newLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	// Default status is 200 OK (Go's default when WriteHeader is not explicitly called)
	return &loggingResponseWriter{ResponseWriter: w, status: http.StatusOK}
}

// WriteHeader stores the status code and forwards the call to the wrapped ResponseWriter
func (w *loggingResponseWriter) WriteHeader(code int) {
	w.status = code
	w.ResponseWriter.WriteHeader(code)
}
