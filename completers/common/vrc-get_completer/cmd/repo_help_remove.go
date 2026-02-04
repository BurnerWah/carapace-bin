package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repo_help_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove repository with specified url, path or name",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repo_help_removeCmd).Standalone()

	repo_helpCmd.AddCommand(repo_help_removeCmd)
}
