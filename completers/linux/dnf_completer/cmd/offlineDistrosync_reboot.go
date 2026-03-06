package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var offlineDistrosync_rebootCmd = &cobra.Command{
	Use:   "reboot",
	Short: "Prepare the system to perform the offline transaction and reboot to start the transaction.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(offlineDistrosync_rebootCmd).Standalone()

	offlineDistrosync_rebootCmd.Flags().Bool("poweroff", false, "Power off the system after the operation is complete")
	offlineDistrosyncCmd.AddCommand(offlineDistrosync_rebootCmd)
}
