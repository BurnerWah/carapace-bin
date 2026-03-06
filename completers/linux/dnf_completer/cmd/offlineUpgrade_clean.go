package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var offlineUpgrade_cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Remove any stored offline transaction and delete cached package files.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(offlineUpgrade_cleanCmd).Standalone()

	offlineUpgradeCmd.AddCommand(offlineUpgrade_cleanCmd)
}
