/*
 * JSR Info
 *
 * This is an API server for a developer web portfolio. Used @jose-rodrigues.info
 *
 * API version: 1.0.0
 * Contact: jsr@jose-rodrigues.info
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package api

import (
	"log"
	"net/http"
	"time"
)

// Logger utility function to log all incoming requests
func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Printf(
			"%s %s %s %s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}
