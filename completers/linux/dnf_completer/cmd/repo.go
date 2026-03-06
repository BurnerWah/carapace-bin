package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repoCmd = &cobra.Command{
	Use:     "repo",
	Short:   "Manage repositories",
	GroupID: "subcommands",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repoCmd).Standalone()
	repoCmd.AddGroup(
		&cobra.Group{ID: "query", Title: ""},
	)

	rootCmd.AddCommand(repoCmd)
}
