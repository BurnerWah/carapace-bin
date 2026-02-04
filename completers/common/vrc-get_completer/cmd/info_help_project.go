package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var info_help_projectCmd = &cobra.Command{
	Use:   "project",
	Short: "Show project information",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(info_help_projectCmd).Standalone()

	info_helpCmd.AddCommand(info_help_projectCmd)
}
