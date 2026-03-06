package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var history_listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List transactions",
	GroupID: "query",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(history_listCmd).Standalone()

	history_listCmd.Flags().Bool("reverse", false, "Reverse the order of transactions")
	historyCmd.AddCommand(history_listCmd)
}
