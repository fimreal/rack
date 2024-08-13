package main

import (
	"github.com/fimreal/rack/module"
	"github.com/fimreal/rack/module/healthcheck"
	"github.com/fimreal/rack/module/nosniffing"
	"github.com/fimreal/rack/pkg/cmd"
)

// Add the imported modules what you want
var MODS = []*module.Module{
	&healthcheck.Module,
	&nosniffing.Module,
}

func init() {
	module.Register(MODS)
	cmd.LoadModuleFlags()
	module.RunCron()
}
