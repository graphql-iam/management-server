package modules

import (
	"github.com/graphql-iam/management-server/src/repository"
	"go.uber.org/fx"
)

var Repository = fx.Module("repository",
	fx.Provide(repository.NewRolesRepository),
	fx.Provide(repository.NewPolicyRepository),
)
