package utils

import "net/http"

// LoggingResponseWriter struct
type LoggingResponseWriter struct {
	http.ResponseWriter
	StatusCode int
}

// WriteHeader method
func (w *LoggingResponseWriter) WriteHeader(code int) {
	w.StatusCode = code
	w.ResponseWriter.WriteHeader(code)
}

// NewLoggingResponseWriter func
func NewLoggingResponseWriter(w http.ResponseWriter) *LoggingResponseWriter {
	return &LoggingResponseWriter{w, http.StatusOK}
}
