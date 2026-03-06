package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var copr_debugCmd = &cobra.Command{
	Use:   "debug",
	Short: "print useful info about the system, useful for debugging",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(copr_debugCmd).Standalone()

	coprCmd.AddCommand(copr_debugCmd)
}
