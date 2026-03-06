package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repo_infoCmd = &cobra.Command{
	Use:     "info",
	Short:   "Print details about repositories",
	GroupID: "query",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repo_infoCmd).Standalone()

	repo_infoCmd.Flags().Bool("all", false, "Show all repositories")
	repo_infoCmd.Flags().Bool("disabled", false, "Show disabled repositories")
	repo_infoCmd.Flags().Bool("enabled", false, "Show enabled repositories (default)")
	repo_infoCmd.Flags().Bool("json", false, "Request json output format")
	repoCmd.AddCommand(repo_infoCmd)
}
