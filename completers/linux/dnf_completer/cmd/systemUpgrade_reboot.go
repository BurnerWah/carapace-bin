package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var systemUpgrade_rebootCmd = &cobra.Command{
	Use:   "reboot",
	Short: "Prepare the system to perform the offline transaction and reboot to start the transaction",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(systemUpgrade_rebootCmd).Standalone()

	systemUpgrade_rebootCmd.Flags().Bool("poweroff", false, "Power off the system after the operation is complete")
	systemUpgradeCmd.AddCommand(systemUpgrade_rebootCmd)
}
