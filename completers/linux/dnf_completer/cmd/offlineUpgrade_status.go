package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var offlineUpgrade_statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show status of the current offline transaction",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(offlineUpgrade_statusCmd).Standalone()

	offlineUpgradeCmd.AddCommand(offlineUpgrade_statusCmd)
}
