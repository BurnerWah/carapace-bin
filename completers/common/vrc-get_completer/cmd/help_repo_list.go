package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_repo_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all repositories",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_repo_listCmd).Standalone()

	help_repoCmd.AddCommand(help_repo_listCmd)
}
