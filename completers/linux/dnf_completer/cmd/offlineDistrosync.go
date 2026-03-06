package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var offlineDistrosyncCmd = &cobra.Command{
	Use:     "offline-distrosync",
	Short:   "Store a distro-sync transaction to be performed offline",
	GroupID: "subcommands",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(offlineDistrosyncCmd).Standalone()

	rootCmd.AddCommand(offlineDistrosyncCmd)
}
