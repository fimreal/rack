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
	serveCmd.Flags().String("jwt_seed", "", "jwt seed")
	serveCmd.Flags().String("jwt_signingmethod", "", "jwt sign method, hmac sha256 or rsa")
	serveCmd.Flags().String("db_driver", "sqlite", "database driver, mysql sqlite postgresql")
	serveCmd.Flags().String("db_host", "", "database host")
	serveCmd.Flags().Int("db_port", 0, "database port")
	serveCmd.Flags().String("db_user", "", "database user")
	serveCmd.Flags().String("db_password", "", "database password")
	serveCmd.Flags().String("db_name", config.AppName, "database name")
}
