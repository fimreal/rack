package config

import (
	"github.com/spf13/pflag"
)

var (
	// global 配置
	port    = pflag.StringP("port", "p", "5000", "指定启动端口")
	workdir = pflag.StringP("workdir", "w", "./", "设置工作目录，用于存放数据库文件")

	// ip2location 配置
	ip2location_enabled  = pflag.Bool("ip2location", false, "是否启用 ip2location 查询")
	ip2location_db_type  = pflag.String("ip2location.db_type", "IPv4", "IP 数据库类型，默认空即为IPv4，可选 IPv6")
	ip2location_db_level = pflag.String("ip2location.db_level", "DB11", "IP 数据库等级，默认为 DB11，可选 DB1 DB3 DB5 DB9 DB11， 数字越大数据库内容越丰富，相应数据库也就越大")
	ip2location_token    = pflag.String("ip2location.token", "", "ip2location lite token")
)
