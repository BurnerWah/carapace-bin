package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var configManagerCmd = &cobra.Command{
	Use:     "config-manager",
	Short:   "Manage configuration",
	GroupID: "subcommands",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(configManagerCmd).Standalone()

	rootCmd.AddCommand(configManagerCmd)
}
