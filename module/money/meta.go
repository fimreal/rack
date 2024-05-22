package money

import (
	"github.com/fimreal/rack/module"
	"github.com/spf13/cobra"
)

const (
	ID            = "money"
	Comment       = "计算理财收益"
	RoutePrefix   = "/money"
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
