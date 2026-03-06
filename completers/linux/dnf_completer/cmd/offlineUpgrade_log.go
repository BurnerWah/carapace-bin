package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var offlineUpgrade_logCmd = &cobra.Command{
	Use:   "log",
	Short: "Show logs from past offline transactions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(offlineUpgrade_logCmd).Standalone()

	offlineUpgrade_logCmd.Flags().String("number", "", "Which log to show")
	offlineUpgradeCmd.AddCommand(offlineUpgrade_logCmd)
}
