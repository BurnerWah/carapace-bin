package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var userPackage_help_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add user package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(userPackage_help_addCmd).Standalone()

	userPackage_helpCmd.AddCommand(userPackage_help_addCmd)
}
