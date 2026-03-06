package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var offlineUpgrade_rebootCmd = &cobra.Command{
	Use:   "reboot",
	Short: "Prepare the system to perform the offline transaction and reboot to start the transaction.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(offlineUpgrade_rebootCmd).Standalone()

	offlineUpgrade_rebootCmd.Flags().Bool("poweroff", false, "Power off the system after the operation is complete")
	offlineUpgradeCmd.AddCommand(offlineUpgrade_rebootCmd)
}
