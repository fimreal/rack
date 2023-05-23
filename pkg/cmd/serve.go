/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/fimreal/rack/pkg/serve"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "flags for Starting (gin) web service",
	Long:  `flags for Starting (gin) web service`,
	Run: func(cmd *cobra.Command, args []string) {
		serve.Run()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().BoolP("daemon", "D", false, "daemon mode")
	serveCmd.Flags().StringP("port", "p", "3333", "set listen port")
	serveCmd.Flags().StringP("workdir", "w", "./", "set work floder")
	serveCmd.Flags().BoolP("allservices", "a", false, "enable all service")

	// ngrok
	serveCmd.Flags().BoolP("ngrok", "n", false, "expose to ngrok")
	serveCmd.Flags().String("ngrok.token", "", "ngrok auth token, find in https://dashboard.ngrok.com/get-started/your-authtoken")
	serveCmd.Flags().String("ngrok.domain", "", "custom domain, eg. my-domain.ngrok.io")

	// services
	serveCmd.Flags().BoolP("common", "", false, "common service")
	serveCmd.Flags().Bool("scripts", false, "scripts service")
	serveCmd.Flags().Bool("docker", false, "dockerhub 镜像查询")
	serveCmd.Flags().Bool("swagger", false, "swagger docs")
	serveCmd.Flags().BoolP("fileserver", "f", false, "启用文件上传下载服务")

	// aliyun
	serveCmd.Flags().Bool("aliyun", false, "启用阿里云安全组控制")
	serveCmd.Flags().String("aliyun.akid", "", "ACCESS_KEY_ID")
	serveCmd.Flags().String("aliyun.aksecret", "", "ACCESS_KEY_SECRET")
	serveCmd.Flags().String("aliyun.regionid", "", "REGION_ID")
	serveCmd.Flags().Bool("aliyun.insecureskipverify", false, "是否跳过证书验证(小容器没有证书会遇到 https 连接证书验证失败)")

	// chatGPT
	serveCmd.Flags().Bool("chatgpt", false, "是否启用 chatgpt")
	serveCmd.Flags().String("chatgpt.api", "https://api.openai.com", "chatgpt API 地址，方便添加个人 api 代理")
	serveCmd.Flags().String("chatgpt.proxyurl", "", "http proxy 地址，方便添加代理")
	serveCmd.Flags().String("chatgpt.token", "", "chatgpt token https://beta.openai.com/account/api-keys")

	// ip2location
	serveCmd.Flags().Bool("ip2location", false, "是否启用 ip2location 查询")
	serveCmd.Flags().String("ip2location.db", "DB11", "IP 数据库等级, 可选 DB1 DB3 DB5 DB9 DB11, 数字越大数据库内容越丰富, 相应数据库也就越大")
	serveCmd.Flags().String("ip2location.token", "", "ip2location lite token")

	// viper bind
	viper.BindPFlags(serveCmd.Flags())
}
