/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/fimreal/goutils/ezap"
	"github.com/fimreal/rack/module"
	"github.com/fimreal/rack/pkg/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "rack",
	Short: "A application could do anything",
	Long:  `A application could do anything`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(fileConfig, config.LoadConfigs)

	rootCmd.PersistentFlags().BoolP("debug", "d", false, "debug mode")
	viper.BindPFlag("debug", rootCmd.PersistentFlags().Lookup("debug"))
	viper.BindPFlags(rootCmd.Flags())
}

var cfgFile string

// initConfig reads in config file and ENV variables if set.
func fileConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		ezap.Debugf("Could not found user home directory: %s", err)

		// Search config in home directory with name ".goproxy" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".rack")
	}

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		ezap.Info("Using config file:", viper.ConfigFileUsed())
	}
}

func LoadModuleFlags() {
	module.NewFlag(rootCmd)
	module.FlagParse(serveCmd)

	viper.BindPFlags(rootCmd.Flags())
	viper.BindPFlags(serveCmd.Flags())
}
