package main

import (
	"github.com/go-openapi/loads"
	"github.com/joeshaw/envdecode"
	"github.com/mllrjb/go-api-bootstrap/generated/swagger"
	"github.com/mllrjb/go-api-bootstrap/generated/swagger/operations"
	"github.com/uber-go/zap"
)

type Config struct {
	BindAddress string `env:"BIND_ADDRESS,default=127.0.0.1"`
	HttpPort    int    `env:"HTTP_PORT,default=8000"`
}

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	var config Config
	err := envdecode.StrictDecode(&config)
	if err != nil {
		logger.Fatal(err)
	}

	spec, err := loads.Analyzed(swagger.SwaggerJSON, "")

	api := operations.NewAPI(spec)
	server := swagger.NewServer(api)
	server.Host = config.BindAddress
	server.Port = config.HttpPort
	defer server.Shutdown()

	server.Serve()
}
