package main

import (
	"github.com/fimreal/rack/module"
	"github.com/fimreal/rack/module/crtag"
	"github.com/fimreal/rack/module/dnsquery"
	"github.com/fimreal/rack/module/emailcheck"
	"github.com/fimreal/rack/module/healthcheck"
	"github.com/fimreal/rack/module/hostinfo"
	"github.com/fimreal/rack/module/ipquery"
	"github.com/fimreal/rack/module/nosniffing"
	"github.com/fimreal/rack/module/randomgenerator"
	"github.com/fimreal/rack/module/scripts"
	"github.com/fimreal/rack/module/servertime"
	"github.com/fimreal/rack/pkg/cmd"
	"github.com/rack-plugins/aliyun"
	"github.com/rack-plugins/chatgpt"
	"github.com/rack-plugins/coord"
	"github.com/rack-plugins/coze"
	"github.com/rack-plugins/email"
	"github.com/rack-plugins/qcloud"
	"github.com/rack-plugins/wechatmp"
)

var MODS = []*module.Module{
	&healthcheck.Module,
	&nosniffing.Module,
	//
	&hostinfo.Module,
	&servertime.Module,
	&ipquery.Module,
	&dnsquery.Module,
	&randomgenerator.Module,
	//
	&emailcheck.Module,
	&scripts.Module,
	&crtag.Module,

	// opt
	&chatgpt.Module,
	&email.Module,
	// &phone.Module, // 待改
	&aliyun.Module,
	// &shorturl.Module,
	&coord.Module,
	&qcloud.Module,
	&wechatmp.Module,
	&coze.Module,
}

func init() {
	module.Register(MODS)
	cmd.LoadModuleFlags()
	module.RunCron()
}
