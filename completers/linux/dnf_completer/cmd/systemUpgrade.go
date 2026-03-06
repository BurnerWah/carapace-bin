package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var systemUpgradeCmd = &cobra.Command{
	Use:     "system-upgrade",
	Short:   "Prepare system for upgrade to a new release",
	GroupID: "subcommands",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(systemUpgradeCmd).Standalone()

	rootCmd.AddCommand(systemUpgradeCmd)
}
