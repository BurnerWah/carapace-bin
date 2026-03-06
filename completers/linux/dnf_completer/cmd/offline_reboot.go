package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var offline_rebootCmd = &cobra.Command{
	Use:   "reboot",
	Short: "Prepare the system to perform the offline transaction and reboot to start the transaction.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(offline_rebootCmd).Standalone()

	offline_rebootCmd.Flags().Bool("poweroff", false, "Power off the system after the operation is complete")
	offlineCmd.AddCommand(offline_rebootCmd)
}
