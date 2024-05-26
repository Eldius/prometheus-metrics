package main

import (
	"github.com/eldius/prometheus-metrics/internal/api"
	"github.com/eldius/prometheus-metrics/internal/config"
)

func main() {
	if err := config.Setup(""); err != nil {
		panic(err)
	}
	if err := api.Start(8080); err != nil {
		panic(err)
	}
}
