package config

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func setFlag() {
	// server
	pflag.StringP("port", "p", "5000", "指定启动端口")
	pflag.StringP("workdir", "w", "./", "设置工作目录，用于存放数据库文件")
	pflag.BoolP("debug", "d", false, "debug mode")
	// swagger
	pflag.Bool("swagger", false, "swagger docs")

	// common
	pflag.Bool("common", true, "通用工具包")

	// mail
	pflag.Bool("mail", false, "启用 http2mail 服务")
	pflag.String("mail.username", "", "smtp (发件人)用户名")
	pflag.String("mail.password", "", "smtp (发件人)密码")
	pflag.String("mail.smtpserver", "", "smtp 服务器地址")
	pflag.String("mail.smtpserverport", "", "smtp 服务器地址")
	pflag.Bool("mail.insecureskipverify", true, "是否跳过证书验证")

	// aliyun
	pflag.Bool("aliyun", false, "启用阿里云安全组控制")
	pflag.String("aliyun.akid", "", "ACCESS_KEY_ID")
	pflag.String("aliyun.aksecret", "", "ACCESS_KEY_SECRET")
	pflag.String("aliyun.regionid", "", "REGION_ID")
	pflag.Bool("aliyun.insecureskipverify", false, "是否跳过证书验证(容器内常遇到 https 连接证书验证失败)")

	// ip2location
	pflag.Bool("ip2location", false, "是否启用 ip2location 查询")
	pflag.String("ip2location.db", "DB11", "IP 数据库等级, 可选 DB1 DB3 DB5 DB9 DB11, 数字越大数据库内容越丰富, 相应数据库也就越大")
	pflag.String("ip2location.token", "", "ip2location lite token")

	// chatGPT
	pflag.Bool("chatgpt", false, "是否启用 chatgpt")
	pflag.String("chatgpt.token", "", "chatgpt token https://beta.openai.com/account/api-keys")

	// load pflag into viper
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
}
