package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_info_packageCmd = &cobra.Command{
	Use:   "package",
	Short: "Show project information",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_info_packageCmd).Standalone()

	help_infoCmd.AddCommand(help_info_packageCmd)
}
