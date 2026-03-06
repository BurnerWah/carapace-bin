package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var systemUpgrade_cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Remove any stored offline transaction and delete cached package files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(systemUpgrade_cleanCmd).Standalone()

	systemUpgradeCmd.AddCommand(systemUpgrade_cleanCmd)
}
