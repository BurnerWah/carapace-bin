package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var history_rollbackCmd = &cobra.Command{
	Use:     "rollback",
	Short:   "Undo all transactions performed after the specified transaction",
	GroupID: "management",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(history_rollbackCmd).Standalone()

	history_rollbackCmd.Flags().Bool("ignore-extras", false, "Don't consider extra packages pulled into the transaction as errors")
	history_rollbackCmd.Flags().Bool("ignore-installed", false, "Don't consider mismatches between installed and stored transaction packages as errors")
	history_rollbackCmd.Flags().Bool("skip-unavailable", false, "Allow skipping unavailable packages")
	historyCmd.AddCommand(history_rollbackCmd)
}
