package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repo_importCmd = &cobra.Command{
	Use:   "import",
	Short: "Import repository list file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repo_importCmd).Standalone()

	repo_importCmd.Flags().BoolP("help", "h", false, "Print help")
	repo_importCmd.Flags().Bool("no-update", false, "do not update local repository cache")
	repo_importCmd.Flags().Bool("offline", false, "do not connect to remote servers, use local caches only. implicitly --no-update")
	repo_importCmd.Flags().BoolP("version", "V", false, "Print version")
	repo_importCmd.Flags().BoolP("yes", "y", false, "skip confirm")
	repoCmd.AddCommand(repo_importCmd)
}
