package auth

import (
	"github.com/fimreal/rack/module"
	"github.com/fimreal/rack/pkg/config"
	"github.com/spf13/cobra"
)

const (
	ID            = "auth"
	Comment       = "auth api"
	RoutePrefix   = "/api/"
	DefaultEnable = false
)

var Module = module.Module{
	ID:          ID,
	Comment:     Comment,
	RouteFunc:   AddRoute,
	RoutePrefix: RoutePrefix,
	FlagFunc:    ServeFlag,
}

// 根据需要配置数据库 dsn
func ServeFlag(serveCmd *cobra.Command) {
	serveCmd.Flags().Bool(ID, DefaultEnable, Comment)
	serveCmd.Flags().String("jwt_secret", "", "jwt secret")
	serveCmd.Flags().String("jwt_signingmethod", "", "jwt sign method, HS256, HS384, HS512, RS256, RS384, RS512, ES256, ES384, ES512, PS256, PS384, PS512")
	serveCmd.Flags().String("db_driver", "sqlite", "database driver, mysql sqlite postgresql")
	serveCmd.Flags().String("db_host", "", "database host")
	serveCmd.Flags().Int("db_port", 0, "database port")
	serveCmd.Flags().String("db_user", "", "database user")
	serveCmd.Flags().String("db_password", "", "database password")
	serveCmd.Flags().String("db_name", config.AppName, "database name")
}
