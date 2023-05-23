package main

import (
	"github.com/fimreal/goutils/ezap"
	"github.com/fimreal/rack/module"
	"github.com/fimreal/rack/module/dnsquery"
	"github.com/fimreal/rack/module/healthcheck"
	"github.com/fimreal/rack/module/ipquery"
	"github.com/fimreal/rack/module/nosniffing"
	"github.com/fimreal/rack/module/randomgenerator"
	"github.com/fimreal/rack/pkg/cmd"
)

var MODS = []*module.Module{
	&healthcheck.Module,
	&nosniffing.Module,
	&ipquery.Module,
	&dnsquery.Module,
	&randomgenerator.Module}

func init() {
	module.Register(MODS)
	cmd.LoadModuleFlags()
	ezap.Info("Starting rack...\n")
}
