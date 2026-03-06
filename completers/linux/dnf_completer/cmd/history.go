package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var historyCmd = &cobra.Command{
	Use:     "history",
	Short:   "Manage transaction history",
	GroupID: "subcommands",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(historyCmd).Standalone()
	historyCmd.AddGroup(
		&cobra.Group{ID: "query", Title: ""},
		&cobra.Group{ID: "management", Title: ""},
	)

	rootCmd.AddCommand(historyCmd)
}
