package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var history_redoCmd = &cobra.Command{
	Use:     "redo",
	Short:   "Repeat all actions from the specified transaction",
	GroupID: "management",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(history_redoCmd).Standalone()

	history_redoCmd.Flags().Bool("skip-unavailable", false, "Allow skipping unavailable packages")
	historyCmd.AddCommand(history_redoCmd)
}
