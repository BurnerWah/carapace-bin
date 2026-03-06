package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var history_undoCmd = &cobra.Command{
	Use:     "undo",
	Short:   "Revert all actions from the specified transaction",
	GroupID: "management",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(history_undoCmd).Standalone()

	history_undoCmd.Flags().Bool("ignore-extras", false, "Don't consider extra packages pulled into the transaction as errors")
	history_undoCmd.Flags().Bool("ignore-installed", false, "Don't consider mismatches between installed and stored transaction packages as errors")
	history_undoCmd.Flags().Bool("skip-unavailable", false, "Allow skipping unavailable packages")
	historyCmd.AddCommand(history_undoCmd)
}
