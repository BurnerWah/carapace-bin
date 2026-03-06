package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var autoremoveCmd = &cobra.Command{
	Use:     "autoremove",
	Short:   "Remove all unneeded packages originally installed as dependencies",
	GroupID: "core",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autoremoveCmd).Standalone()

	autoremoveCmd.Flags().Bool("offline", false, "Store the transaction to be performed offline")
	autoremoveCmd.Flags().String("store", "", "Store the current transaction in a directory instead of running it")
	rootCmd.AddCommand(autoremoveCmd)
}
