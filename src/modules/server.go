package modules

import (
	"github.com/graphql-iam/management-server/src/server"
	"go.uber.org/fx"
	"net/http"
)

var Server = fx.Module("server",
	fx.Provide(server.NewServer),
	fx.Invoke(func(server *http.Server) {}),
)
