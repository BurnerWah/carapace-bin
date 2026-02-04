package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repo_help_exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export user repository list file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repo_help_exportCmd).Standalone()

	repo_helpCmd.AddCommand(repo_help_exportCmd)
}
