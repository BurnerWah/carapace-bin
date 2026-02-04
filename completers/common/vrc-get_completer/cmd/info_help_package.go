package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var info_help_packageCmd = &cobra.Command{
	Use:   "package",
	Short: "Show project information",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(info_help_packageCmd).Standalone()

	info_helpCmd.AddCommand(info_help_packageCmd)
}
