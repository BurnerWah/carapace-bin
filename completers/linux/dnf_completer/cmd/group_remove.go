package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var group_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove groups, including their packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(group_removeCmd).Standalone()

	group_removeCmd.Flags().Bool("no-packages", false, "Operate on groups only, no packages are changed")
	group_removeCmd.Flags().Bool("offline", false, "Store the transaction to be performed offline")
	group_removeCmd.Flags().String("store", "", "Store the current transaction in a directory instead of running it")
	groupCmd.AddCommand(group_removeCmd)
}
