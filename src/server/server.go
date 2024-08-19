package server

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/graphql-iam/management-server/src/config"
	"github.com/graphql-iam/management-server/src/handler"
	"go.uber.org/fx"
	"net"
	"net/http"
)

func NewServer(lc fx.Lifecycle, cfg config.Config, rolesHandler handler.RolesHandler, policyHandler handler.PolicyHandler) *http.Server {
	r := gin.Default()
	r.GET("/role", rolesHandler.GetRole)
	r.GET("/roles", rolesHandler.GetRoles)
	r.POST("/role", rolesHandler.CreateRole)
	r.DELETE("/role", rolesHandler.DeleteRole)
	r.POST("/role/attach", rolesHandler.AttachPolicyToRole)
	r.POST("/role/detach", rolesHandler.DetachPolicyFromRole)
	r.GET("/policies", policyHandler.GetPolicies)
	r.GET("/policy", policyHandler.GetPolicy)
	r.POST("/policy", policyHandler.CreatePolicy)
	r.PUT("/policy", policyHandler.UpdatePolicy)
	r.DELETE("/policy", policyHandler.DeletePolicy)
	srv := &http.Server{
		Addr:    fmt.Sprintf("localhost:%d", cfg.Port),
		Handler: r.Handler(),
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				return err
			}
			fmt.Println("Starting HTTP server at", srv.Addr)
			go srv.Serve(ln)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})
	return srv
}
