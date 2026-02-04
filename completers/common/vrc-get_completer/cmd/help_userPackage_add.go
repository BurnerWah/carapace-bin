package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_userPackage_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add user package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_userPackage_addCmd).Standalone()

	help_userPackageCmd.AddCommand(help_userPackage_addCmd)
}
