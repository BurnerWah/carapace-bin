package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repo_help_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add remote or local repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repo_help_addCmd).Standalone()

	repo_helpCmd.AddCommand(repo_help_addCmd)
}
