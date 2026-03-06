package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var history_infoCmd = &cobra.Command{
	Use:     "info",
	Short:   "Print details about transactions",
	GroupID: "query",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(history_infoCmd).Standalone()

	history_infoCmd.Flags().Bool("reverse", false, "Reverse the order of transactions")
	historyCmd.AddCommand(history_infoCmd)
}
