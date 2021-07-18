package callback

import (
	"context"

	"github.com/sfpyhub/webhook-test/definitions"
)

type Service interface {
	Respond(ctx context.Context, request *definitions.Request) (interface{}, error)
}
