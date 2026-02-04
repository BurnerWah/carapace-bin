package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_repo_cleanupCmd = &cobra.Command{
	Use:   "cleanup",
	Short: "Cleanup repositories in Repos directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_repo_cleanupCmd).Standalone()

	help_repoCmd.AddCommand(help_repo_cleanupCmd)
}
