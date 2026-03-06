package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var moduleCmd = &cobra.Command{
	Use:     "module",
	Short:   "Manage modules",
	GroupID: "subcommands",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(moduleCmd).Standalone()
	moduleCmd.AddGroup(
		&cobra.Group{ID: "query", Title: ""},
		&cobra.Group{ID: "management", Title: ""},
	)

	rootCmd.AddCommand(moduleCmd)
}
