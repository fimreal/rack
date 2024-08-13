package nosniffing

import (
	"github.com/fimreal/rack/module"
	"github.com/spf13/cobra"
)

const (
	ID            = "nosniffing"
	Comment       = "Block malicious probes. Built-in rules to identify and intercept simple malicious visits."
	RoutePrefix   = "/api/nosniffing"
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
	serveCmd.Flags().Bool(ID+"_banip", false, "Enable IP blocking")
	serveCmd.Flags().Int(ID+"_duration", 10, "IP blocking duration, minute")
	serveCmd.Flags().String(ID+"_key", "", "Release IP restriction key (added to the route), empty to disable")
}
