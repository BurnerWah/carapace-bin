package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var history_storeCmd = &cobra.Command{
	Use:     "store",
	Short:   "Store transaction to a file [experimental]",
	GroupID: "management",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(history_storeCmd).Standalone()

	history_storeCmd.Flags().StringP("output", "o", "", "Path to a directory for storing the transaction")
	historyCmd.AddCommand(history_storeCmd)
}
