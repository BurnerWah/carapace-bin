package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mark_userCmd = &cobra.Command{
	Use:   "user",
	Short: "Mark package as user-installed",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mark_userCmd).Standalone()

	markCmd.AddCommand(mark_userCmd)
}
