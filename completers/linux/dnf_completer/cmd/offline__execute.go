package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var offline__executeCmd = &cobra.Command{
	Use:    "_execute",
	Short:  "Internal use only",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(offline__executeCmd).Standalone()

	offlineCmd.AddCommand(offline__executeCmd)
}
