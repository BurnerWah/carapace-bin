package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var swapCmd = &cobra.Command{
	Use:     "swap",
	Short:   "Remove software and install another in one transaction",
	GroupID: "core",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(swapCmd).Standalone()

	swapCmd.Flags().Bool("allowerasing", false, "Allow erasing of installed packages to resolve problems")
	swapCmd.Flags().Bool("offline", false, "Store the transaction to be performed offline")
	swapCmd.Flags().String("store", "", "Store the current transaction in a directory instead of running it")
	rootCmd.AddCommand(swapCmd)
}
