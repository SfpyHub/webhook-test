package callback

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/sfpyhub/go-sfpy/requests"
	"github.com/sfpyhub/go-sfpy/responses"
	"github.com/sfpyhub/go-sfpy/sfpy"
	"github.com/sfpyhub/go-sfpy/types"
	"github.com/sfpyhub/webhook-test/definitions"
)

type callbackService struct {
	logger log.Logger

	client *sfpy.Client
}

func (as *callbackService) Create(ctx context.Context, request *definitions.Request) (interface{}, error) {
	response, err := as.client.Endpoints.AddOrder(ctx, &requests.Request{
		OrderService: &requests.OrderService{
			Order: &requests.Order{
				Address:   "0x742Df1612A701a130c71C9Ce3971Db549917cE29",
				Reference: "anyInternalID",
				ChainID:   uint(types.MAINNET),
				Cart: &requests.Cart{
					Source:      "MyWebsite",
					CancelURL:   "https://localhost:6666/cart",
					CompleteURL: "https://localhost:6666/order/complete",
				},
				PurchaseTotals: &requests.PurchaseTotal{
					SubTotal: &requests.PriceMoney{
						Currency: "USD",
						Amount:   100000, // $1,000 in cents
					},
					Discount: &requests.PriceMoney{
						Currency: "USD",
						Amount:   2000, // $20 in cents
					},
					TaxTotal: &requests.PriceMoney{
						Currency: "USD",
						Amount:   1000, // $10 in cents
					},
				},
			},
		},
	})

	if err != nil {
		level.Error(as.logger).Log(
			"message", "unable to create order",
			"error", err.Error(),
		)
		return nil, err
	}

	order := responses.Order{}
	if err := json.Unmarshal(response.Data, &order); err != nil {
		return nil, err
	}

	as.client.ConstructLink(order.Token)

	return &order, nil
}

func (as *callbackService) Respond(ctx context.Context, request *definitions.Request) (interface{}, error) {
	signature := request.Head.Signature
	if len(signature) == 0 {
		return nil, errors.New("signature missing from webhook")
	}

	if err := as.client.ValidateSignature(signature, request.Data); err != nil {
		level.Error(as.logger).Log(
			"message", "invalid signature",
			"signature", signature,
			"error", err.Error(),
		)

		return nil, errors.New("invalid signature")
	}

	var payload interface{}
	switch request.Data.Type {
	case "payment:created":
		payload = &responses.PaymentNotification{}
	case "refund:created":
		payload = &responses.RefundNotification{}
	default:
		return nil, errors.New("unknown event")
	}

	message := request.Data.Data.Notification

	if err := json.Unmarshal(message, &payload); err != nil {
		return nil, err
	}

	fmt.Println(fmt.Sprintf("%+v", payload))

	return "success", nil
}

// NewCallbackService gives us a new one
func NewCallbackService(l log.Logger, c *sfpy.Client) Service {
	return &callbackService{
		logger: l,
		client: c,
	}
}
