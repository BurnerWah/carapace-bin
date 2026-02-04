package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repo_cleanupCmd = &cobra.Command{
	Use:   "cleanup",
	Short: "Cleanup repositories in Repos directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repo_cleanupCmd).Standalone()

	repo_cleanupCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	repo_cleanupCmd.Flags().Bool("no-update", false, "do not update local repository cache")
	repo_cleanupCmd.Flags().Bool("offline", false, "do not connect to remote servers, use local caches only. implicitly --no-update")
	repo_cleanupCmd.Flags().BoolP("version", "V", false, "Print version")
	repoCmd.AddCommand(repo_cleanupCmd)
}
