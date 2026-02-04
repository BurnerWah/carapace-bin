package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repo_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove repository with specified url, path or name",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repo_removeCmd).Standalone()

	repo_removeCmd.Flags().BoolP("help", "h", false, "Print help")
	repo_removeCmd.Flags().Bool("id", false, "Find repository to remove by id")
	repo_removeCmd.Flags().Bool("name", false, "Find repository to remove by name")
	repo_removeCmd.Flags().Bool("no-update", false, "do not update local repository cache")
	repo_removeCmd.Flags().Bool("offline", false, "do not connect to remote servers, use local caches only. implicitly --no-update")
	repo_removeCmd.Flags().Bool("path", false, "Find repository to remove by local path")
	repo_removeCmd.Flags().Bool("url", false, "Find repository to remove by url")
	repo_removeCmd.Flags().BoolP("version", "V", false, "Print version")
	repoCmd.AddCommand(repo_removeCmd)
}
