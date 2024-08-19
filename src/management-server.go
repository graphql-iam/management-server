package main

import (
	"github.com/graphql-iam/management-server/src/config"
	"github.com/graphql-iam/management-server/src/database"
	"github.com/graphql-iam/management-server/src/modules"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(config.NewConfig),
		fx.Provide(database.NewDatabase),
		modules.Repository,
		modules.Service,
		modules.Handler,
		modules.Middleware,
		modules.Server,
	).Run()
}
