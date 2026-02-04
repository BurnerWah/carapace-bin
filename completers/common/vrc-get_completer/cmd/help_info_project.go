package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_info_projectCmd = &cobra.Command{
	Use:   "project",
	Short: "Show project information",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_info_projectCmd).Standalone()

	help_infoCmd.AddCommand(help_info_projectCmd)
}
