package http

import (
	nethttp "net/http"
	"time"
)

func TimeoutMiddleware(next nethttp.Handler) nethttp.Handler {
	return nethttp.TimeoutHandler(
		next,
		30*time.Second,
		`{"error":"request timeout"}`,
	)
}
