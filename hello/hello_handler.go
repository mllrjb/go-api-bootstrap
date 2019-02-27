package hello

import (
	middleware "github.com/go-openapi/runtime/middleware"
	"github.com/mllrjb/go-api-bootstrap/auth"
	"github.com/mllrjb/go-api-bootstrap/generated/swagger/operations"
	"go.uber.org/zap"
)

func NewHelloController(logger *zap.Logger) HelloController {
	return HelloController{logger: logger}
}

type HelloController struct {
	logger *zap.Logger
}

func (c *HelloController) GetHelloHandler() operations.GetHelloWorldHandler {
	return operations.GetHelloWorldHandlerFunc(func(params operations.GetHelloWorldParams, principal *auth.UserPrincipal) middleware.Responder {
		return operations.NewHelloWorldOK()
	})
}

func (c *HelloController) DeleteHelloHandler() operations.DeleteHelloWorldHandler {
	return operations.DeleteHelloWorldHandlerFunc(func(params operations.DeleteHelloWorldParams, principal *auth.UserPrincipal) middleware.Responder {
		return operations.NewHelloWorldOK()
	})
}
