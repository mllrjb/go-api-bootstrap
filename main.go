package main

import (
	"github.com/go-openapi/loads"
	"github.com/joeshaw/envdecode"
	"github.com/mllrjb/go-api-bootstrap/generated/swagger"
	"github.com/mllrjb/go-api-bootstrap/generated/swagger/operations"
	"go.uber.org/zap"
)

type Config struct {
	BindAddress string `env:"BIND_ADDRESS,default=127.0.0.1"`
	HttpPort    int    `env:"HTTP_PORT,default=8000"`
}

func main() {
	rootLogger, _ := zap.NewProduction()
	logger := rootLogger.Sugar()
	defer logger.Sync()

	var config Config
	err := envdecode.StrictDecode(&config)
	if err != nil {
		logger.Fatalf("cannot decode configuration: %v", err)
	}

	spec, err := loads.Analyzed(swagger.SwaggerJSON, "")

	api := operations.NewAPI(spec)
	server := swagger.NewServer(api)
	server.ConfigureAPI()
	server.Host = config.BindAddress
	server.Port = config.HttpPort
	defer server.Shutdown()

	server.Serve()
}
