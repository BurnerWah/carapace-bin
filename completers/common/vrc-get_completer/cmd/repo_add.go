package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repo_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add remote or local repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repo_addCmd).Standalone()

	repo_addCmd.Flags().StringSliceP("header", "H", nil, "Headers")
	repo_addCmd.Flags().BoolP("help", "h", false, "Print help")
	repo_addCmd.Flags().Bool("no-update", false, "do not update local repository cache")
	repo_addCmd.Flags().Bool("offline", false, "do not connect to remote servers, use local caches only. implicitly --no-update")
	repo_addCmd.Flags().BoolP("version", "V", false, "Print version")
	repoCmd.AddCommand(repo_addCmd)
}
