package nosniffing

import (
	"github.com/fimreal/rack/module"
	"github.com/spf13/cobra"
)

var (
	NOROUTEMSG = "Hey, world!\n"
	ROBOTSTXT  = "User-agent: *\nDisallow: /"
)

const (
	ID            = "nosniffing"
	Comment       = "避免嗅探功能"
	RoutePrefix   = "/"
	DefaultEnable = true
)

var Module = module.Module{
	ID:      ID,
	Comment: Comment,
	// gin route
	RouteFunc:   AddRoute,
	RoutePrefix: RoutePrefix,
	// cobra flag
	FlagFunc: ServeFlag,
}

func ServeFlag(serveCmd *cobra.Command) {
	serveCmd.Flags().Bool(ID, DefaultEnable, Comment)
}
