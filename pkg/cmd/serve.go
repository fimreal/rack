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
	serveCmd.Flags().StringP("workdir", "w", ".", "set application working floder")
	serveCmd.Flags().BoolP("allservices", "a", false, "enable all service")

	// ngrok
	serveCmd.Flags().BoolP("ngrok", "n", false, "expose to ngrok")
	serveCmd.Flags().String("ngrok.token", "", "ngrok auth token, find in https://dashboard.ngrok.com/get-started/your-authtoken")
	serveCmd.Flags().String("ngrok.domain", "", "custom domain, eg. my-domain.ngrok.io")

	// // services
	// serveCmd.Flags().Bool("swagger", false, "swagger docs")
	// serveCmd.Flags().BoolP("fileserver", "f", false, "启用文件上传下载服务")

	// viper bind
	viper.BindPFlags(serveCmd.Flags())
}
