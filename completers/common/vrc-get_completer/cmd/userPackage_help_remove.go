package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var userPackage_help_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove user package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(userPackage_help_removeCmd).Standalone()

	userPackage_helpCmd.AddCommand(userPackage_help_removeCmd)
}
