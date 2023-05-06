/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

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
	rootCmd.PersistentFlags().BoolP("debug", "d", false, "debug mode")
	viper.BindPFlag("debug", rootCmd.PersistentFlags().Lookup("debug"))
	// viper.BindPFlags(rootCmd.Flags())
}
