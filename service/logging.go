package service

import (
	"net/http"
	"time"

	log "github.com/golang/glog"
)

// Use a logWriter struct to implement the ResponseWriter interface.
type logWriter struct {
	status int
	length int
	http.ResponseWriter
}

// Write implements ResponseWriter and writes the response
func (w *logWriter) Write(b []byte) (int, error) {
	w.length = w.length + len(b)
	return w.ResponseWriter.Write(b)
}

// WriteHeader implements ResponseWriter and accesses status code
func (w *logWriter) WriteHeader(code int) {
	w.status = code
	w.ResponseWriter.WriteHeader(code)
}

// Logger allows each request to be logged
func Logger(handler http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		localW := &logWriter{200, 0, w}
		start := time.Now()
		handler.ServeHTTP(localW, r)
		log.Infoln(r.Method, name+r.URL.String(), localW.length, localW.status, time.Since(start)/time.Nanosecond)
	})
}
