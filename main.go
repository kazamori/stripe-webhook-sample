package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/caarlos0/env"
	"golang.org/x/exp/slog"

	"github.com/kazamori/stripe/config"
	"github.com/kazamori/stripe/handler"
)

func main() {
	cfg := config.Stripe{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("failed to set config: %s", err)
	}

	http.HandleFunc("/webhook", handler.Webhook)

	slog.Info("serve `/webhook` as an endpoint", "port", cfg.Port)
	addr := fmt.Sprintf(":%d", cfg.Port)
	log.Fatal(http.ListenAndServe(addr, nil))
}
