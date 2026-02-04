package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repo_exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export user repository list file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repo_exportCmd).Standalone()

	repo_exportCmd.Flags().BoolP("help", "h", false, "Print help")
	repo_exportCmd.Flags().Bool("no-update", false, "do not update local repository cache")
	repo_exportCmd.Flags().Bool("offline", false, "do not connect to remote servers, use local caches only. implicitly --no-update")
	repo_exportCmd.Flags().BoolP("version", "V", false, "Print version")
	repoCmd.AddCommand(repo_exportCmd)
}
