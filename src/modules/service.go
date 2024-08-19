package modules

import (
	"github.com/graphql-iam/management-server/src/service"
	"go.uber.org/fx"
)

var Service = fx.Module("service",
	fx.Provide(service.NewRolesService),
	fx.Provide(service.NewPolicyService),
)
