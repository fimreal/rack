package website

import (
	"github.com/fimreal/rack/module"
	"github.com/spf13/cobra"
)

const (
	ID            = "website"
	Comment       = "embed website"
	RoutePrefix   = "/"
	DefaultEnable = true
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
