package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_repo_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add remote or local repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_repo_addCmd).Standalone()

	help_repoCmd.AddCommand(help_repo_addCmd)
}
