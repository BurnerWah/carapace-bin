package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_repo_packagesCmd = &cobra.Command{
	Use:   "packages",
	Short: "List packages in specified repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_repo_packagesCmd).Standalone()

	help_repoCmd.AddCommand(help_repo_packagesCmd)
}
