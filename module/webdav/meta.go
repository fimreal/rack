package fileserver

import (
	"github.com/fimreal/rack/module"
	"github.com/spf13/cobra"
)

const (
	ID            = "fileserver"
	Comment       = "brwose fileserver, Does not support mixing of other mods"
	RoutePrefix   = "/"
	DefaultEnable = false
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
