package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/stripe/stripe-go/v74"
	"golang.org/x/exp/slog"
)

func Webhook(w http.ResponseWriter, req *http.Request) {
	const MaxBodyBytes = int64(65536)
	req.Body = http.MaxBytesReader(w, req.Body, MaxBodyBytes)
	payload, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading request body: %v\n", err)
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	event := stripe.Event{}

	if err := json.Unmarshal(payload, &event); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse webhook body json: %v\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Unmarshal the event data into an appropriate struct depending on its Type
	switch event.Type {
	case "payment_intent.created":
		fmt.Println("PaymentIntent created")

	case "payment_intent.succeeded":
		var paymentIntent stripe.PaymentIntent
		err := json.Unmarshal(event.Data.Raw, &paymentIntent)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing webhook JSON: %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		fmt.Println("PaymentIntent was successful!")
		prettyPrint(paymentIntent)
		//		fmt.Printf("%+v\n", paymentIntent)

		// turn on smart plug
		if err := requestMakerEvent(cfg.IftttTurnOn); err != nil {
			slog.Error("failed to trigger on ifttt", "err", err)
		} else {
			go func() {
				<-time.After(3 * time.Second)
				if err := requestMakerEvent(cfg.IftttTurnOff); err != nil {
					slog.Error("failed to trigger off on ifttt", "err", err)
				}
				turnOn = false
				slog.Info("completed trigger off")
			}()
		}
	case "charge.succeeded":
		fmt.Println("Charge was successfully!")

	case "payment_method.attached":
		var paymentMethod stripe.PaymentMethod
		err := json.Unmarshal(event.Data.Raw, &paymentMethod)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing webhook JSON: %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		fmt.Println("PaymentMethod was attached to a Customer!")
	// ... handle other event types
	default:
		fmt.Fprintf(os.Stderr, "Unhandled event type: %s\n", event.Type)
	}

	fmt.Println("=================================")
	w.WriteHeader(http.StatusOK)
}
