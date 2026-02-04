package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrade specified package or all packages to latest or specified version",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_upgradeCmd).Standalone()

	helpCmd.AddCommand(help_upgradeCmd)
}
