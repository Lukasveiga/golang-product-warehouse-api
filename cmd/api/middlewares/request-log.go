package middlewares

import (
	"fmt"
	"net/http"
	"time"
)

func RequestLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func (res http.ResponseWriter, req *http.Request) {
		startTime := time.Now()

		next.ServeHTTP(res, req)

		elapsedTime := time.Since(startTime)

		fmt.Printf("Http request -> Method: %s - Path: %s - Duration: %s - Datetime: %s\n",
		req.Method, req.URL.Path, elapsedTime.String(),time.Now())
	})
}