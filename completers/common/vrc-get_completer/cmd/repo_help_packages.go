package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repo_help_packagesCmd = &cobra.Command{
	Use:   "packages",
	Short: "List packages in specified repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repo_help_packagesCmd).Standalone()

	repo_helpCmd.AddCommand(repo_help_packagesCmd)
}
