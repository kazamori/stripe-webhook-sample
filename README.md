# stripe-webhook-sample

Sample application using [stripe-go](https://github.com/stripe/stripe-go) to provide webhook endpoint for stripe.

## How to run (with test mode)

Set your secret key. See your keys here: https://dashboard.stripe.com/apikeys

```bash
$ export SK_KEY="sk_test_xxx"
```

Start up the server provides your webhook endpoint.

```bash
$ go run main.go 
2023/06/11 12:56:38 INFO serve `/webhook` as an endpoint port=8080
```

Install [Stripe CLI](https://stripe.com/docs/stripe-cli), then listen any payment events.

```bash
$ stripe login
$ stripe listen --forward-to localhost:8080/webhook
```

Trigger an any payment event.

```bash
$ stripe trigger payment_intent.succeeded
```

Scan QR code of your product with your smartphone.

(This is an example of my test product for test mode)

![qr code](./qr_test_fZeeWxa7y326g48288-small.png)

Use [test cards](https://stripe.com/docs/testing#cards) to simulate a successful payment.

After you paid using above QR code, you can see some event logs in the stripe listener/webhook sample application.

### Integrate with IFTTT webhook

After a webhook endpoint was called, request another webhook endpoint on [IFTTT](https://ifttt.com).

Use simple webhook events described at [Receive a web request](https://ifttt.com/maker_webhooks/triggers/event). Define 3 environemnt variables to integrate with it.

```bash
$ export IFTTT_KEY="..."
$ export IFTTT_TURN_ON="turn-on-something"
$ export IFTTT_TURN_OFF="turn-off-something"
```

When the turn-on event is done, then trigger the turn-off event after 3 seconds.

```
=================================
2023/06/12 19:32:30 INFO Congratulations! You've fired the turn-on-something event
=================================
2023/06/12 19:32:33 INFO Congratulations! You've fired the turn-off-something event
2023/06/12 19:32:33 INFO completed trigger off
```

## Reference

* https://stripe.com
* https://stripe.com/docs/payments/handling-payment-events
* https://ifttt.com
