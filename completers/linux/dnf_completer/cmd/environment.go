package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var environmentCmd = &cobra.Command{
	Use:     "environment",
	Short:   "Manage environments",
	GroupID: "subcommands",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(environmentCmd).Standalone()
	environmentCmd.AddGroup(
		&cobra.Group{ID: "query", Title: ""},
	)

	rootCmd.AddCommand(environmentCmd)
}
