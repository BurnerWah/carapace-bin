package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repoCmd = &cobra.Command{
	Use:   "repo",
	Short: "Commands around repositories",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repoCmd).Standalone()

	repoCmd.Flags().BoolP("help", "h", false, "Print help")
	repoCmd.Flags().BoolP("version", "V", false, "Print version")
	rootCmd.AddCommand(repoCmd)
}
