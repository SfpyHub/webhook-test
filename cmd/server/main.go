package main

import (
	"flag"
	"os"

	"github.com/sfpyhub/webhook-test/pkg/server"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/sfpyhub/notify/utils"
)

const (
	envAPIEnvironment     string = "ENVIRONMENT"
	defaultAPIEnvironment string = "LOCAL"

	envWebhookSecretKey     string = "WEBHOOK_SECRET"
	defaultWebhookSecretKey string = "780a1fa7d8a16d5479edd427320116ba2ea2f49a2dc39f1a924390ca4bf36c7b"

	envSfpyApiKey     string = "SFPY_API_KEY"
	defaultSfpyApiKey string = "f012313568a6ca280479d2b521814c75ff373e4d68ba2ac9c81cb085a62a4578"
)

// Variables set during compile via -X options
var (
	Version string
	Build   string
	GitHash string
)

func main() {
	var (
		logLevel    = flag.String("log.level", "info", "Logging Level: [debug | info | warn | error]")
		logFormat   = flag.String("log.format", "json", "Logging Format: [text | json]")
		httpCmdAddr = flag.String("http.cmd.addr", ":5678", "Command, Debug, and Metrics listen Address")
	)

	flag.Parse()

	// Set log level
	var levelFilter level.Option
	switch *logLevel {
	case "debug":
		levelFilter = level.AllowDebug()
	case "info":
		levelFilter = level.AllowInfo()
	case "warn":
		levelFilter = level.AllowWarn()
	case "error":
		levelFilter = level.AllowError()
	default:
		panic("no valid LOG_LEVEL set")
	}

	var logger log.Logger
	logger = level.NewFilter(logger, levelFilter)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC, "caller", log.DefaultCaller)
	switch *logFormat {
	case "text":
		logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	case "json":
		logger = log.NewJSONLogger(log.NewSyncWriter(os.Stderr))
	default:
		logger = log.NewJSONLogger(log.NewSyncWriter(os.Stderr))
	}

	// logging version
	logger.Log("Version", Version)
	logger.Log("Build", Build)
	logger.Log("Git Hash", GitHash)

	// logging flags
	logger.Log("log.level", *logLevel)
	logger.Log("log.format", *logFormat)
	logger.Log("http.cmd.addr", *httpCmdAddr)

	environment := utils.EnvString(envAPIEnvironment, defaultAPIEnvironment)
	secretkey := utils.EnvString(envWebhookSecretKey, defaultWebhookSecretKey)
	apikey := utils.EnvString(envSfpyApiKey, defaultSfpyApiKey)

	config := server.Config{
		Environment:      environment,
		HTTPCmdAddr:      *httpCmdAddr,
		WebhookSecretKey: secretkey,
		SfpyApiKey:       apikey,
	}

	serv, err := server.NewServer(config, logger)
	if err != nil {
		level.Error(logger).Log("err", "err")
		panic(err)
	}

	// run server
	err = serv.Run()
	logger.Log("err", err)
	logger.Log("exiting")
}
