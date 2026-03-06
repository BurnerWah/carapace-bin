package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var offlineDistrosync_cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Remove any stored offline transaction and delete cached package files.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(offlineDistrosync_cleanCmd).Standalone()

	offlineDistrosyncCmd.AddCommand(offlineDistrosync_cleanCmd)
}
