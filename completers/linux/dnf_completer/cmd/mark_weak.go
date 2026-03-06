package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mark_weakCmd = &cobra.Command{
	Use:   "weak",
	Short: "Mark package as a weak dependency",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mark_weakCmd).Standalone()

	markCmd.AddCommand(mark_weakCmd)
}
