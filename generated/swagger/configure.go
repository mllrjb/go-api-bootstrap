// This file is safe to edit. Once it exists it will not be overwritten

package swagger

import (
	"crypto/tls"
	"net/http"

	"github.com/mllrjb/go-api-bootstrap/generated/swagger/operations"
	"github.com/mllrjb/go-api-bootstrap/server"
)

//go:generate swagger generate server --target ../../go-api-bootstrap --name My --spec ../swagger.yaml --server-package swagger --principal auth.UserPrincipal

func configureFlags(api *operations.API) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.API) http.Handler {
	return server.ConfigureAPI(api)
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}
