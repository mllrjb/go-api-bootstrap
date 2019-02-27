package server

import (
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	"go.uber.org/zap"

	"github.com/mllrjb/go-api-bootstrap/auth"
	"github.com/mllrjb/go-api-bootstrap/generated/swagger/operations"
	"github.com/mllrjb/go-api-bootstrap/hello"
)

func ConfigureAPI(api *operations.API) http.Handler {
	logger, _ := zap.NewProduction()
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// Applies when the "Authorization" header is set
	api.BearerAuth = func(token string) (*auth.UserPrincipal, error) {
		principal := auth.UserPrincipal{
			Username: "anonymous",
			UserID:   -1,
		}
		return &principal, nil
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()
	helloHandler := hello.NewHelloController(logger)
	api.GetHelloWorldHandler = helloHandler.GetHelloHandler()
	api.DeleteHelloWorldHandler = helloHandler.DeleteHelloHandler()

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
