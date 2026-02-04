package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repo_help_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all repositories",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repo_help_listCmd).Standalone()

	repo_helpCmd.AddCommand(repo_help_listCmd)
}
