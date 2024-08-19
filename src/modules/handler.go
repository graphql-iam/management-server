package modules

import (
	"github.com/graphql-iam/management-server/src/handler"
	"go.uber.org/fx"
)

var Handler = fx.Module("handler",
	fx.Provide(handler.NewRolesHandler),
	fx.Provide(handler.NewPolicyHandler),
)
