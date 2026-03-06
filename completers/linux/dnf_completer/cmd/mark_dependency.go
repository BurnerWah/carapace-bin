package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mark_dependencyCmd = &cobra.Command{
	Use:   "dependency",
	Short: "Mark package as a dependency",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mark_dependencyCmd).Standalone()

	markCmd.AddCommand(mark_dependencyCmd)
}
