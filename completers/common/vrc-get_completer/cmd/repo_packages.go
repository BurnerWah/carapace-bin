package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repo_packagesCmd = &cobra.Command{
	Use:   "packages",
	Short: "List packages in specified repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repo_packagesCmd).Standalone()

	repo_packagesCmd.Flags().BoolP("help", "h", false, "Print help")
	repo_packagesCmd.Flags().Bool("no-update", false, "do not update local repository cache")
	repo_packagesCmd.Flags().Bool("offline", false, "do not connect to remote servers, use local caches only. implicitly --no-update")
	repo_packagesCmd.Flags().BoolP("version", "V", false, "Print version")
	repoCmd.AddCommand(repo_packagesCmd)
}
