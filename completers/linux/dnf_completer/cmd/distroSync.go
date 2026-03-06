package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var distroSyncCmd = &cobra.Command{
	Use:     "distro-sync",
	Short:   "Upgrade or downgrade installed software to the latest available versions",
	GroupID: "core",
	Aliases: []string{"dsync"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(distroSyncCmd).Standalone()

	distroSyncCmd.Flags().Bool("allowerasing", false, "Allow erasing of installed packages to resolve problems")
	distroSyncCmd.Flags().Bool("downloadonly", false, "Only download packages for a transaction")
	distroSyncCmd.Flags().Bool("offline", false, "Store the transaction to be performed offline")
	distroSyncCmd.Flags().Bool("skip-broken", false, "Allow resolving of depsolve problems by skipping packages")
	distroSyncCmd.Flags().Bool("skip-unavailable", false, "Allow skipping unavailable packages")
	distroSyncCmd.Flags().String("store", "", "Store the current transaction in a directory instead of running it")
	rootCmd.AddCommand(distroSyncCmd)
}
