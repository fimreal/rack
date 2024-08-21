// 无用
package trace

import (
	"github.com/fimreal/rack/module"
	"github.com/spf13/cobra"
)

const (
	ID            = "trace"
	Comment       = "trace "
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
