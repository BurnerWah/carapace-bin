package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_repo_importCmd = &cobra.Command{
	Use:   "import",
	Short: "Import repository list file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_repo_importCmd).Standalone()

	help_repoCmd.AddCommand(help_repo_importCmd)
}
