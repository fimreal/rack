/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/fimreal/rack/pkg/serve"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "show rack version and module version",
	Long:  `show rack version and module version`,
	Run: func(cmd *cobra.Command, args []string) {
		serve.PrintVersion()
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
