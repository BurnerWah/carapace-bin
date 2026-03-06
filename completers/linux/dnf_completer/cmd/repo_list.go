package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repo_listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List repositories",
	GroupID: "query",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repo_listCmd).Standalone()

	repo_listCmd.Flags().Bool("all", false, "Show all repositories")
	repo_listCmd.Flags().Bool("disabled", false, "Show disabled repositories")
	repo_listCmd.Flags().Bool("enabled", false, "Show enabled repositories (default)")
	repo_listCmd.Flags().Bool("json", false, "Request json output format")
	repoCmd.AddCommand(repo_listCmd)
}
