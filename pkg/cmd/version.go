/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/fimreal/rack/pkg/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "show rack version and module version",
	Long:  `show rack version and module version`,
	Run: func(cmd *cobra.Command, args []string) {
		config.PrintVersion()
		if viper.GetBool("show-mods") {
			config.PrintMods()
		}
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
	versionCmd.Flags().BoolP("show-mods", "M", false, "show embed mods")

	viper.BindPFlags(versionCmd.Flags())
}
