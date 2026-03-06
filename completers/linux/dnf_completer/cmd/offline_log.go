package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var offline_logCmd = &cobra.Command{
	Use:   "log",
	Short: "Show logs from past offline transactions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(offline_logCmd).Standalone()

	offline_logCmd.Flags().String("number", "", "Which log to show")
	offlineCmd.AddCommand(offline_logCmd)
}
