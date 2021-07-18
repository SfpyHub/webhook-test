package callback

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/sfpyhub/webhook-test/definitions"
)

type callbackService struct {
	logger log.Logger

	secretkey string
}

func (as *callbackService) Respond(ctx context.Context, request *definitions.Request) (interface{}, error) {
	signature := request.Head.Signature
	if len(signature) == 0 {
		return nil, errors.New("signature missing from webhook")
	}

	if len(as.secretkey) > 0 {
		if err := ValidateSignature(signature, request.Data, []byte(as.secretkey)); err != nil {
			level.Error(as.logger).Log(
				"message", "invalid signature",
				"signature", signature,
				"error", err.Error(),
			)

			return nil, errors.New("invalid signature")
		}
	}

	var payload interface{}
	switch request.Data.Type {
	case "payment:created":
		payload = &definitions.PaymentNotification{}
	case "refund:created":
		payload = &definitions.RefundNotification{}
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

// NewApiKeyService gives us a new one
func NewCallbackService(l log.Logger, s string) Service {
	return &callbackService{
		logger:    l,
		secretkey: s,
	}
}
