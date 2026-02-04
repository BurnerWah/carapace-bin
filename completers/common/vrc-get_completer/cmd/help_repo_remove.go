package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_repo_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove repository with specified url, path or name",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_repo_removeCmd).Standalone()

	help_repoCmd.AddCommand(help_repo_removeCmd)
}
