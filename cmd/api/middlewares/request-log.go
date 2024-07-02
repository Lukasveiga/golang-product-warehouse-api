package middlewares

import (
	"log/slog"
	"net/http"
	"time"
)

func RequestLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func (res http.ResponseWriter, req *http.Request) {
		startTime := time.Now()

		next.ServeHTTP(res, req)

		elapsedTime := time.Since(startTime)

		slog.Info("http request information", "method", req.Method, "path", req.URL.Path, "duration_ms", elapsedTime.String())
	})
}