package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:     "remove",
	Short:   "Remove (uninstall) software",
	GroupID: "core",
	Aliases: []string{"rm"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(removeCmd).Standalone()

	removeCmd.Flags().Bool("no-autoremove", false, "Disable removal of dependencies that are no longer used")
	removeCmd.Flags().Bool("noautoremove", false, "Alias for '--no-autoremove'")
	removeCmd.Flags().Bool("offline", false, "Store the transaction to be performed offline")
	removeCmd.Flags().String("store", "", "Store the current transaction in a directory instead of running it")
	removeCmd.Flag("noautoremove").Hidden = true
	rootCmd.AddCommand(removeCmd)
}
