package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repo_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all repositories",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repo_listCmd).Standalone()

	repo_listCmd.Flags().BoolP("help", "h", false, "Print help")
	repo_listCmd.Flags().Bool("no-update", false, "do not update local repository cache")
	repo_listCmd.Flags().Bool("offline", false, "do not connect to remote servers, use local caches only. implicitly --no-update")
	repo_listCmd.Flags().BoolP("version", "V", false, "Print version")
	repoCmd.AddCommand(repo_listCmd)
}
