package metrics

import (
	"github.com/fimreal/rack/module"
	"github.com/spf13/cobra"
)

const (
	ID            = "metrics"
	Comment       = "metrics api"
	RoutePrefix   = "/"
	DefaultEnable = false
)

var Module = module.Module{
	ID:          ID,
	Comment:     Comment,
	RouteFunc:   AddRoute,
	RoutePrefix: RoutePrefix,
	FlagFunc:    ServeFlag,
}

func ServeFlag(serveCmd *cobra.Command) {
	serveCmd.Flags().Bool(ID, DefaultEnable, Comment)
}
