package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var offlineDistrosync_logCmd = &cobra.Command{
	Use:   "log",
	Short: "Show logs from past offline transactions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(offlineDistrosync_logCmd).Standalone()

	offlineDistrosync_logCmd.Flags().String("number", "", "Which log to show")
	offlineDistrosyncCmd.AddCommand(offlineDistrosync_logCmd)
}
