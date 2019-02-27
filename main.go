package main

import (
	log "github.com/cihub/seelog"
	"github.com/go-openapi/loads"
	"github.com/joeshaw/envdecode"
	"github.com/logrhythm/go-api-bootstrap/generated/swagger"
	"github.com/logrhythm/go-api-bootstrap/generated/swagger/operations"
)

type Config struct {
	BindAddress  string `env:"BIND_ADDRESS,default=127.0.0.1"`
	HttpPort     int `env:"HTTP_PORT,default=8000"`
}

func main() {
	defer log.Flush()

	var config Config
	err := envdecode.StrictDecode(&config)
	if err != nil {
		log.Critical(err)
}

	spec, err := loads.Analyzed(swagger.SwaggerJSON, "")

	api := operations.NewAPI(spec)
	server := swagger.NewServer(api)
	server.Host = config.BindAddress
	server.Port = config.HttpPort
	defer server.Shutdown()

	server.Serve()
}
