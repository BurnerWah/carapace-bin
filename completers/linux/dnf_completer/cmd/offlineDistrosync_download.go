package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var offlineDistrosync_downloadCmd = &cobra.Command{
	Use:    "download",
	Short:  "Alias for 'distro-sync --offline'",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(offlineDistrosync_downloadCmd).Standalone()

	offlineDistrosyncCmd.AddCommand(offlineDistrosync_downloadCmd)
}
