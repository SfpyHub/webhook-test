package callback

import (
	"net/http"

	"github.com/go-kit/kit/log"
	"github.com/sfpyhub/webhook-test/pkg/encoder_decoder"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

// MakeHandler <>
func MakeHandler(es Service, logger log.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(encoder_decoder.EncodeError),
	}

	var respondEndpoint endpoint.Endpoint
	respondEndpoint = makeRespondEndpoint(es)

	respondHandler := kithttp.NewServer(
		respondEndpoint,
		encoder_decoder.DecodePostRequest,
		encoder_decoder.EncodeResponse,
		opts...,
	)

	r := mux.NewRouter()

	r.Handle("/v1/webhook/respond", respondHandler).Methods("POST")

	return r
}
