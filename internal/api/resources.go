package api

import (
	"fmt"
	"github.com/eldius/prometheus-metrics/internal/middleware"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Start(port int) error {
	mux := http.NewServeMux()

	mux.HandleFunc("/", Home)
	mux.Handle("/metrics", promhttp.Handler())

	s := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: middleware.LoadMiddlewares(mux, middleware.WithLoggingHandler()),
	}
	return s.ListenAndServe()
}

func Home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func Metrics(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
