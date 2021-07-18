package callback

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/sfpyhub/webhook-test/definitions"
)

func makeRespondEndpoint(service Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*definitions.Request)
		res, err := service.Respond(ctx, req)
		if err != nil {
			return nil, err
		}
		return res, nil
	}
}
