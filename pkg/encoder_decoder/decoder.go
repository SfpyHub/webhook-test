package encoder_decoder

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/sfpyhub/webhook-test/definitions"
)

func DecodePostRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req definitions.Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}

	req.Head.Signature = r.Header.Get("X-SFPY-SIGNATURE")

	return &req, nil
}
