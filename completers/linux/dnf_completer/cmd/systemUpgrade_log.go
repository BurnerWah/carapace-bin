package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var systemUpgrade_logCmd = &cobra.Command{
	Use:   "log",
	Short: "Show logs from past offline transactions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(systemUpgrade_logCmd).Standalone()

	systemUpgrade_logCmd.Flags().String("number", "", "Which log to show")
	systemUpgradeCmd.AddCommand(systemUpgrade_logCmd)
}
