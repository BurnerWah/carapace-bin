package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var offlineCmd = &cobra.Command{
	Use:     "offline",
	Short:   "Manage offline transactions",
	GroupID: "subcommands",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(offlineCmd).Standalone()

	rootCmd.AddCommand(offlineCmd)
}
