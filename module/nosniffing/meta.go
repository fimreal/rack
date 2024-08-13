package nosniffing

import (
	"github.com/fimreal/rack/module"
	"github.com/spf13/cobra"
)

const (
	NOROUTEMSG = "Hey, world!\n"
	ROBOTSTXT  = "User-agent: *\nDisallow: /"
)

const (
	ID            = "nosniffing"
	Comment       = "refusal of access rules"
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
