package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_repoCmd = &cobra.Command{
	Use:   "repo",
	Short: "Commands around repositories",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_repoCmd).Standalone()

	helpCmd.AddCommand(help_repoCmd)
}
