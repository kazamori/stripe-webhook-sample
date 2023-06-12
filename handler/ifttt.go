package handler

import (
	"fmt"
	"io"
	"net/http"

	"golang.org/x/exp/slog"
)

var (
	turnOn bool // must lock via mutex
)

const (
	// https://ifttt.com/maker_webhooks/triggers/event
	host  = "maker.ifttt.com"
	tpath = "/trigger/%s/with/key/%s"
)

func requestMakerEvent(event string) error {
	if cfg.IftttKey == "" {
		slog.Info("no IFTTT key")
		return nil
	}
	if turnOn && event == cfg.IftttTurnOn {
		return nil
	} else if !turnOn && event == cfg.IftttTurnOff {
		return nil
	}

	path := fmt.Sprintf(tpath, event, cfg.IftttKey)
	uri := fmt.Sprintf("https://%s%s", host, path)
	r, err := http.Get(uri)
	if err != nil {
		return fmt.Errorf("failed to get: %w", err)
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return fmt.Errorf("failed to read body: %w", err)
	}
	slog.Info(string(body))
	turnOn = event == cfg.IftttTurnOn
	return nil
}
