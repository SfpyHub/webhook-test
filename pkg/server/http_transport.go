package server

import (
	"net/http"

	"github.com/go-kit/kit/log"
	"github.com/sfpyhub/go-sfpy/sfpy"
	"github.com/sfpyhub/webhook-test/pkg/callback"
)

type httpTransport struct {
	logger log.Logger
	addr   string
	secret string
	apikey string
}

// NewHTTPTransport <>
func NewHTTPTransport(config Config, logger log.Logger) (RunCloser, error) {
	logger = log.With(logger, "module", "httpTransport")

	return &httpTransport{
		logger: logger,
		addr:   config.HTTPCmdAddr,
		secret: config.WebhookSecretKey,
		apikey: config.SfpyApiKey,
	}, nil
}

func (t *httpTransport) Run() error {
	logger := log.With(t.logger, "function", "Run")

	client := sfpy.NewClient(t.apikey, t.secret)

	var cs callback.Service
	cs = callback.NewCallbackService(t.logger, client)

	mux := http.NewServeMux()

	mux.Handle("/v1/webhook/", callback.MakeHandler(cs, logger))

	http.Handle("/", accessControl(mux))

	logger.Log("Listening", t.addr)

	return http.ListenAndServe(t.addr, nil)
}

func (t *httpTransport) Close() error {
	return nil
}

func accessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization, X-SFPY-API-KEY")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}
