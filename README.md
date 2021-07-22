# webhook-test

This is a small service that showcases how to create an order and respond to webhooks after a payment is made.

## Installation

You can clone this repository by `go get`ing the package

```
go get -u "github.com/sfpyhub/webhook-test"
```

## Configuration

To use this project, you should change the Api Key and Shared Secret Key referenced in `main.go` with your own keys. If you don't have an SFPY account yet you can get one for free by [creating an account](https://app.sfpy.co/#/signup)

Once you have your API key and Shared Secret Key, open up `cmd/server/main.go` in your editor and replace the following lines with your keys:

```golang
envWebhookSecretKey     string = "WEBHOOK_SECRET"
defaultWebhookSecretKey string = "<YOUR_SHARED_SECRET>"

envSfpyApiKey     string = "SFPY_API_KEY"
defaultSfpyApiKey string = "<YOUR_API_KEY>"
```

## Run the service

To run the service open your terminal and navigate to the root of this repository and type

```
go run cmd/server/main.go
```

The App listens on port `:5678`

## Create Order

Using postman you can make a POST request to `http://localhost:5678/v1/webhook/create`. The body of this request should be an empty JSON object i.e `{}`. You can see the code for this over [here](https://github.com/SfpyHub/webhook-test/blob/main/pkg/callback/callback.go#L24)

### Response

The response will look something like this:

```
{
  "token": "C3SLANR2DNCAUU2H1GE0",
  "merchant": "C387LVMSARCE1G3S15M0",
  "address": "0x742Df1612A701a130c71C9Ce3971Db549917cE29",
  "reference": "",
  "chain_id": 1,
  "state": "STARTED",
  "cart": {
    "token": "C3SLANR2DNCAUU2H1GF0",
    "request_id": "C3SLANR2DNCAUU2H1GE0",
    "source": "mywebsite",
    "complete_url": "https://localhost:6666/order/complete",
    "cancel_url": "https://localhost:6666/cart"
  },
  "purchase_total": {
    "token": "C3SLANR2DNCAUU2H1GEG",
    "request_id": "C3SLANR2DNCAUU2H1GE0",
    "discount": {
      "amount": 2000,
      "currency": "USD"
    },
    "sub_total": {
      "amount": 100000,
      "currency": "USD"
    },
    "tax_total": {
      "amount": 1000,
      "currency": "USD"
    },
    "grand_total": {
      "amount": 99000,
      "currency": "USD"
    }
  },
  "created_at": "2021-07-22T11:24:15.317896Z",
  "updated_at": "2021-07-22T11:24:15.317896Z"
}
```

## SDK

This project uses our native golang SDK, [sfpy-go](https://github.com/sfpyhub/go-sfpy).


## Questions

If you have any questions, comments or concerns, please open up an issue or join us on our [Discord Server](https://discord.com/invite/PQffzU78Fx)
