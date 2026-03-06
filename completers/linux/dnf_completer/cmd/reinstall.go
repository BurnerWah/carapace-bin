package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var reinstallCmd = &cobra.Command{
	Use:     "reinstall",
	Short:   "Reinstall software",
	GroupID: "core",
	Aliases: []string{"rei"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(reinstallCmd).Standalone()

	reinstallCmd.Flags().Bool("allow-downgrade", false, "Allow downgrade of dependencies for resolve of requested operation")
	reinstallCmd.Flags().Bool("allowerasing", false, "Allow erasing of installed packages to resolve problems")
	reinstallCmd.Flags().Bool("downloadonly", false, "Only download packages for a transaction")
	reinstallCmd.Flags().Bool("no-allow-downgrade", false, "Disable downgrade of dependencies for resolve of requested operation")
	reinstallCmd.Flags().Bool("offline", false, "Store the transaction to be performed offline")
	reinstallCmd.Flags().Bool("skip-broken", false, "Allow resolving of depsolve problems by skipping packages")
	reinstallCmd.Flags().Bool("skip-unavailable", false, "Allow skipping unavailable packages")
	reinstallCmd.Flags().String("store", "", "Store the current transaction in a directory instead of running it")
	rootCmd.AddCommand(reinstallCmd)
}
