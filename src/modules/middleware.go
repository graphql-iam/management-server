package modules

import (
	"github.com/graphql-iam/management-server/src/middleware"
	"go.uber.org/fx"
)

var Middleware = fx.Module("middleware", fx.Provide(middleware.NewAuthMiddleware))
