package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var offlineUpgrade_downloadCmd = &cobra.Command{
	Use:    "download",
	Short:  "Alias for 'upgrade --offline'",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(offlineUpgrade_downloadCmd).Standalone()

	offlineUpgradeCmd.AddCommand(offlineUpgrade_downloadCmd)
}
