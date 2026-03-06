package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var group_installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install groups, including their packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(group_installCmd).Standalone()

	group_installCmd.Flags().Bool("allow-downgrade", false, "Allow downgrade of dependencies for resolve of requested operation")
	group_installCmd.Flags().Bool("allowerasing", false, "Allow erasing of installed packages to resolve problems")
	group_installCmd.Flags().Bool("downloadonly", false, "Only download packages for a transaction")
	group_installCmd.Flags().Bool("no-allow-downgrade", false, "Disable downgrade of dependencies for resolve of requested operation")
	group_installCmd.Flags().Bool("no-packages", false, "Operate on groups only, no packages are changed")
	group_installCmd.Flags().Bool("offline", false, "Store the transaction to be performed offline")
	group_installCmd.Flags().Bool("skip-broken", false, "Allow resolving of depsolve problems by skipping packages")
	group_installCmd.Flags().Bool("skip-unavailable", false, "Allow skipping unavailable packages")
	group_installCmd.Flags().String("store", "", "Store the current transaction in a directory instead of running it")
	group_installCmd.Flags().Bool("with-optional", false, "Include optional packages from group")
	groupCmd.AddCommand(group_installCmd)
}
