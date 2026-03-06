package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var group_upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrade comp groups, including their packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(group_upgradeCmd).Standalone()

	group_upgradeCmd.Flags().Bool("allow-downgrade", false, "Allow downgrade of dependencies for resolve of requested operation")
	group_upgradeCmd.Flags().Bool("allowerasing", false, "Allow erasing of installed packages to resolve problems")
	group_upgradeCmd.Flags().Bool("downloadonly", false, "Only download packages for a transaction")
	group_upgradeCmd.Flags().Bool("no-allow-downgrade", false, "Disable downgrade of dependencies for resolve of requested operation")
	group_upgradeCmd.Flags().Bool("offline", false, "Store the transaction to be performed offline")
	group_upgradeCmd.Flags().Bool("skip-unavailable", false, "Allow skipping unavailable packages")
	group_upgradeCmd.Flags().String("store", "", "Store the current transaction in a directory instead of running it")
	groupCmd.AddCommand(group_upgradeCmd)
}
