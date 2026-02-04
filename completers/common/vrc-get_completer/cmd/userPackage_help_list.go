package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var userPackage_help_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all user packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(userPackage_help_listCmd).Standalone()

	userPackage_helpCmd.AddCommand(userPackage_help_listCmd)
}
