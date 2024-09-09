package main

import (
	"github.com/fimreal/rack/module"
	"github.com/fimreal/rack/module/auth"
	"github.com/fimreal/rack/module/healthcheck"
	"github.com/fimreal/rack/module/nosniffing"
	"github.com/fimreal/rack/module/website"
)

// Add the imported modules what you want
var MODS = []*module.Module{
	&healthcheck.Module,
	&nosniffing.Module,
	&website.Module,
	&auth.Module,
}

func init() {
	module.Register(MODS)
	module.RunCron()
}
